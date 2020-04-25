package repo

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "repo",
		Short:   "opens the browser to this repository's url",
		Aliases: []string{"rp", "r"},
		Run: func(command *cobra.Command, args []string) {
			cwd, getCwdErr := os.Getwd()
			if getCwdErr != nil {
				log.Errorf("unable to retrieve current working directory: %s", getCwdErr)
				os.Exit(1)
			}
			repo, getRepoErr := git.PlainOpen(cwd)
			if getRepoErr != nil {
				log.Errorf("current directory may not be a git repository: %s", getRepoErr)
			}
			remotes, getRemotesErr := repo.Remotes()
			if getRemotesErr != nil {
				log.Errorf("unable to retrieve remotes from %s: %s", cwd, getRemotesErr)
			}
			var remoteURL string
			for _, remote := range remotes {
				if remote.Config().Name == "origin" {
					remoteURL = remote.Config().URLs[0]
					break
				}
			}
			log.Error(remoteURL)
		},
	}
	return &cmd
}
