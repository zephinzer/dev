package repo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/git"
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
			remote, err := git.GetRemote(gitRepoRoot)
			if err != nil {
				cmdutils.ExitWithError(fmt.Sprintf("failed to get remotes from repository at '%s'", gitRepoRoot), constants.ExitErrorInput|constants.ExitErrorSystem)
			}

			url, parseURLErr := validator.ParseURL(remote.URL)
			if parseURLErr != nil {
				cmdutils.ExitWithError(fmt.Sprintf("failed to parse the url '%s': %s", remote.URL, parseURLErr), constants.ExitErrorApplication|constants.ExitErrorUser)
			}

			log.Infof("opening url '%s' in the default browser application...", url.String())
			utils.OpenURIWithDefaultApplication(url.String())
			os.Exit(constants.ExitOK)
		},
	}
	return &cmd
}
