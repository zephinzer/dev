package repository

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/utils"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Aliases: constants.RepositoryAliases,
		Short:   "verifies that the repositories specified in the configuration exists",
		Run:     run,
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

	homeDir := cmdutils.GetHomeDirectory()

	errorCount := 0
	successCount := 0
	newCount := 0
	workspaceIndex := map[string]bool{}
	for _, repository := range config.Global.Repositories {
		for _, workspace := range repository.Workspaces {
			workspaceIndex[workspace] = true
		}
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

		log.Tracef("processing repository '%s': %s", repositoryName, repository.Description)
		log.Tracef("  url : %s", repository.URL)
		log.Tracef("  path: %s", localPath)
		fileInfo, lstatError := os.Stat(path.Join(localPath, "/.git"))
		if lstatError != nil {
			if !os.IsNotExist(lstatError) {
				log.Warnf("  failed to check existence: %s", lstatError)
				errorCount++
				continue
			}
		} else {
			repositoryExistsLocally = fileInfo.IsDir()
		}
		if repositoryExistsLocally {
			log.Debugf("repository '%s' already exists, skipping...", repositoryName)
		} else {
			log.Debugf("repository '%s' does not exist, attempting to clone to '%s'", repositoryName, localPath)
			cloneError := utils.GitClone(repository.URL, localPath)
			newCount++
			if cloneError != nil {
				log.Warnf("failed to clone repository from url '%s': %s", repository.URL, cloneError)
				errorCount++
				continue
			}
			log.Infof("repository '%s' successfully set up at '%s'", repository.URL, localPath)
		}
		successCount++
	}

	workspaces := []string{}
	for workspaceName, _ := range workspaceIndex {
		workspaces = append(workspaces, workspaceName)
	}

	log.Infof("available workspaces   : %s", strings.Join(workspaces, ", "))
	log.Infof("total repositories     : %v", len(config.Global.Repositories))
	log.Infof("new repositories       : %v", newCount)
	log.Infof("successfully processed : %v", len(config.Global.Repositories))
	log.Infof("failed to process      : %v", errorCount)

	os.Exit(errorCount)
}
