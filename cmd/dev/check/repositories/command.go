package repositories

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/validator"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Aliases: constants.RepositoryAliases,
		Short:   "verifies that the repositories specified in the configuration exists",
		Run: func(command *cobra.Command, args []string) {
			if config.Global.Repositories == nil {
				log.Error("no repositories have been defined")
				os.Exit(1)
				return
			} else if len(config.Global.Repositories) == 0 {
				log.Error("no repositories found")
				os.Exit(1)
				return
			}

			homeDir, getHomeDirError := os.UserHomeDir()
			if getHomeDirError != nil {
				log.Errorf("unable to retrieve user's home directory: %s", getHomeDirError)
				os.Exit(1)
			}

			for _, repository := range config.Global.Repositories {
				repositoryName := "unnamed"
				if len(repository.Name) > 0 {
					repositoryName = repository.Name
				}
				log.Debugf("repository '%s': %s", repository.Name, repository.Description)
				var localPath string
				if validator.IsGitHTTPUrl(repository.CloneURL) || validator.IsGitSSHUrl(repository.CloneURL) {
					parsedURL, parseError := validator.ParseURL(repository.CloneURL)
					if parseError != nil {
						log.Warnf("failed to parse clone url '%s'", repository.CloneURL)
						continue
					}
					localPath = path.Join(homeDir, parsedURL.Hostname, parsedURL.User, parsedURL.Path)
				}
				log.Debugf("  url : %s", repository.CloneURL)
				log.Debugf("  path: %s", localPath)
				repositoryExistsLocally := false
				fileInfo, lstatError := os.Lstat(path.Join(localPath, "/.git"))
				if lstatError != nil {
					if lstatError != os.ErrNotExist {
						log.Warnf("  failed to check existence: %s", lstatError)
						continue
					}
					repositoryExistsLocally = false
				} else {
					repositoryExistsLocally = fileInfo.IsDir()
				}
				log.Debugf("  exists: %v", repositoryExistsLocally)
				if repositoryExistsLocally {
					log.Printf(constants.CheckSuccessFormat, repositoryName)
				} else {
					log.Printf(constants.CheckFailureFormat, repositoryName)
				}
				log.Printf("(path: %s", localPath)
				if len(repository.Workspaces) > 0 {
					log.Printf(" | workspaces: %s", strings.Join(repository.Workspaces, ","))
				}
				log.Printf(" | src: %s", repository.CloneURL)
				log.Printf(")\n")
			}
		},
	}
	return &cmd
}
