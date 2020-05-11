package repository

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/usvc/dev/internal/config"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/utils"
	"github.com/usvc/dev/pkg/validator"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Aliases: constants.RepositoryAliases,
		Short:   "verifies that the repositories specified in the configuration exists",
		Run:     Run,
	}
	return &cmd
}

func Run(command *cobra.Command, args []string) {
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

	errorCount := 0
	for _, repository := range config.Global.Repositories {
		repositoryName := "unnamed"
		if len(repository.Name) > 0 {
			repositoryName = repository.Name
		}
		log.Infof("processing repository '%s': %s", repositoryName, repository.Description)
		var localPath string
		if validator.IsGitHTTPUrl(repository.CloneURL) || validator.IsGitSSHUrl(repository.CloneURL) {
			parsedURL, parseError := validator.ParseURL(repository.CloneURL)
			if parseError != nil {
				log.Warnf("failed to parse clone url '%s'", repository.CloneURL)
				continue
			}
			localPath = path.Join(homeDir, parsedURL.Hostname, parsedURL.User, parsedURL.Path)
		} else if _, parseError := validator.ParseURL(repository.CloneURL); parseError != nil {
			log.Warnf("failed to parse url '%s'", repository.CloneURL)
			continue
		} else {
			cloneURL, getCloneURLError := utils.GetSshCloneUrlFromHttpLinkUrl(repository.CloneURL)
			if getCloneURLError != nil {
				log.Warnf("failed to convert '%s' to a git SSH clone URL", repository.CloneURL)
				continue
			}
			parsedURL, parseError := validator.ParseURL(cloneURL)
			if parseError != nil {
				log.Warnf("failed to parse clone url '%s'", cloneURL)
				continue
			}
			localPath = path.Join(homeDir, parsedURL.Hostname, parsedURL.User, parsedURL.Path)
		}
		log.Debugf("  url : %s", repository.CloneURL)
		log.Debugf("  path: %s", localPath)
		repositoryExistsLocally := false
		fileInfo, lstatError := os.Stat(path.Join(localPath, "/.git"))
		if lstatError != nil {
			if !os.IsNotExist(lstatError) {
				log.Warnf("  failed to check existence: %s", lstatError)
				continue
			}
			repositoryExistsLocally = false
		} else {
			repositoryExistsLocally = fileInfo.IsDir()
		}
		if repositoryExistsLocally {
			log.Debugf("repository '%s' exists", repositoryName)
		} else {
			log.Debugf("repository '%s' does not exist, attempting to clone to '%s'", repositoryName, localPath)
			_, cloneError := git.PlainClone(localPath, false, &git.CloneOptions{
				URL: repository.CloneURL,
			})
			if cloneError != nil {
				log.Warnf("failed to clone repository from url '%s': %s", repository.CloneURL, cloneError)
				errorCount++
				continue
			}
			log.Infof("repository '%s' successfully set up at '%s'", repository.CloneURL, localPath)
		}
	}
	if errorCount > 0 {
		os.Exit(errorCount)
	}
}
