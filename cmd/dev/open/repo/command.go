package repo

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
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
			cwd, getCwdErr := os.Getwd()
			if getCwdErr != nil {
				log.Errorf("failed to retrieve current working directory: %s", getCwdErr)
				os.Exit(constants.ExitErrorSystem)
			}
			gitRepoRoot, findGitRepoRootError := utils.FindParentContainingChildDirectory(".git", cwd)
			if findGitRepoRootError != nil {
				log.Errorf("failed to detect if current directory resides in a git repository: %s", findGitRepoRootError)
				os.Exit(constants.ExitErrorSystem)
			} else if len(gitRepoRoot) == 0 {
				log.Errorf("current directory does not seem to reside in a git repository: %s", findGitRepoRootError)
				os.Exit(constants.ExitErrorUser)
			}
			repo, getRepoErr := git.PlainOpen(gitRepoRoot)
			if getRepoErr != nil {
				log.Errorf("current directory may not be a git repository: %s", getRepoErr)
				os.Exit(constants.ExitErrorUser)
			}
			remotes, getRemotesErr := repo.Remotes()
			if getRemotesErr != nil {
				log.Errorf("unable to retrieve remotes from %s: %s", cwd, getRemotesErr)
				os.Exit(constants.ExitErrorUser)
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
				log.Errorf("unable to parse the url '%s': %s", remoteURL, parseURLErr)
				os.Exit(constants.ExitErrorApplication | constants.ExitErrorUser)
			}

			log.Infof("opening url '%s' in the default browser application...", url.GetHTTPSURL())
			utils.OpenURIWithDefaultApplication(url.GetHTTPSURL())
			os.Exit(constants.ExitOK)
		},
	}
	return &cmd
}
