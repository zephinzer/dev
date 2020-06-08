package repositories

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Aliases: constants.RepositoryAliases,
		Short:   "verifies that the repositories specified in the configuration exists",
		Example: strings.Trim(`
dev check repos
		`, "\n "),
		Run: run,
	}
	return &cmd
}

func run(command *cobra.Command, args []string) {
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
	workspaceIndex := map[string]bool{}
	for _, repository := range config.Global.Repositories {
		repositoryName := "unnamed"
		if len(repository.Name) > 0 {
			repositoryName = repository.Name
		}
		localPath, getPathError := repository.GetPath(homeDir)
		if getPathError != nil {
			log.Warnf("failed to process local path '%s': %s", repositoryName, getPathError)
			errorCount++
			continue
		}
		repositoryExistsLocally := false

		log.Debugf("repository '%s': %s", repository.Name, repository.Description)
		log.Debugf("  url : %s", repository.URL)
		log.Debugf("  path: %s", localPath)
		fileInfo, lstatError := os.Stat(path.Join(localPath, "/.git"))
		if lstatError != nil {
			if !os.IsNotExist(lstatError) {
				log.Warnf("  failed to check existence: %s", lstatError)
				errorCount++
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
		if len(repository.Workspaces) > 0 {
			log.Printf("\n  workspaces: %s", strings.Join(repository.Workspaces, ","))
			for _, workspace := range repository.Workspaces {
				workspaceIndex[workspace] = true
			}
		}
		log.Printf("\n  path      : %s", localPath)
		log.Printf("\n  url       : %s", repository.URL)
		log.Printf("\n")
	}
	workspaces := []string{}
	for workspace, _ := range workspaceIndex {
		workspaces = append(workspaces, workspace)
	}

	log.Infof("available workspaces   : %s", strings.Join(workspaces, ", "))
	log.Infof("total repositories     : %v", len(config.Global.Repositories))

	os.Exit(errorCount)
}
