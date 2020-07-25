package repo

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils"
	"github.com/zephinzer/dev/pkg/validator"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Short:   "opens the browser to this repository's url",
		Aliases: constants.RepositoryAliases,
		Run: func(command *cobra.Command, args []string) {
			gitRepoRoot := cmdutils.GetGitRepoRootFromWorkingDirectory()
			repo, getRepoErr := git.PlainOpen(gitRepoRoot)
			if getRepoErr != nil {
				cmdutils.ExitWithError(fmt.Sprintf("current directory may not be a git repository: %s", getRepoErr), constants.ExitErrorUser)
			}
			remotes, getRemotesErr := repo.Remotes()
			if getRemotesErr != nil {
				cwd, _ := os.Getwd()
				cmdutils.ExitWithError(fmt.Sprintf("unable to retrieve remotes from %s: %s", cwd, getRemotesErr), constants.ExitErrorUser)
			}
			var remoteURL string
			for _, remote := range remotes {
				if remote.Config().Name == "origin" {
					remoteURL = remote.Config().URLs[0]
					break
				}
			}

			url, parseURLErr := validator.ParseURL(remoteURL)
			if parseURLErr != nil {
				cmdutils.ExitWithError(fmt.Sprintf("failed to parse the url '%s': %s", remoteURL, parseURLErr), constants.ExitErrorApplication|constants.ExitErrorUser)
			}

			log.Infof("opening url '%s' in the default browser application...", url.String())
			utils.OpenURIWithDefaultApplication(url.String())
			os.Exit(constants.ExitOK)
		},
	}
	return &cmd
}
