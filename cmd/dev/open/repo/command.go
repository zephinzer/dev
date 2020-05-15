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

const (
	ExitOK               = 0
	ExitErrorSystem      = 1
	ExitErrorUser        = 2
	ExitErrorApplicaiton = 4
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Short:   "opens the browser to this repository's url",
		Aliases: constants.RepositoryAliases,
		Run: func(command *cobra.Command, args []string) {
			cwd, getCwdErr := os.Getwd()
			if getCwdErr != nil {
				log.Errorf("unable to retrieve current working directory: %s", getCwdErr)
				os.Exit(ExitErrorSystem)
			}
			repo, getRepoErr := git.PlainOpen(cwd)
			if getRepoErr != nil {
				log.Errorf("current directory may not be a git repository: %s", getRepoErr)
				os.Exit(ExitErrorUser)
			}
			remotes, getRemotesErr := repo.Remotes()
			if getRemotesErr != nil {
				log.Errorf("unable to retrieve remotes from %s: %s", cwd, getRemotesErr)
				os.Exit(ExitErrorUser)
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
				os.Exit(ExitErrorApplicaiton | ExitErrorUser)
			}

			log.Tracef("opening url '%s' in the default browser application...", url.String())
			utils.OpenURIWithDefaultApplication(url.String())
			os.Exit(ExitOK)
		},
	}
	return &cmd
}
