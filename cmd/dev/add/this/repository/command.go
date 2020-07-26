package repository

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/git"
	"github.com/zephinzer/dev/internal/log"
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     constants.RepositoryCanonicalNoun,
		Aliases: constants.RepositoryAliases,
		Short:   "adds the current repository you're in to dev's configuration",
		Run:     run,
	}
	return &cmd
}

func run(cmd *cobra.Command, args []string) {
	// get current repository
	gitRepoRoot := cmdutils.GetGitRepoRootFromWorkingDirectory()
	targetRepo, err := git.GetRemote(gitRepoRoot)
	if err != nil {
		cmdutils.ExitWithError(
			fmt.Sprintf("failed to get remote url for repository at '%s'", gitRepoRoot),
			constants.ExitErrorInput|constants.ExitErrorSystem,
		)
	}

	// parse loaded configurations into a map of set repositories
	loadedReposMap := cmdutils.GetLoadedRepositories()
	loadedReposMapCount := 1
	log.Debug("repositories already present:")
	for repoPath, configPath := range loadedReposMap {
		log.Debugf("%v. %s from '%s'", loadedReposMapCount, repoPath, configPath)
		loadedReposMapCount++
	}

	// filter out repositories that are already in the configuration
	repoURLs := cmdutils.FilterOutLoadedRepositories([]string{targetRepo.URL}, loadedReposMap)
	if len(repoURLs) == 0 {
		log.Info("current repository is already in your configuration, exiting with status 0 now")
		os.Exit(0)
	}

	// add it
	log.Infof("adding repository with url '%v'", targetRepo.URL)

	cmdutils.AddRepositoryToConfig(targetRepo.URL)
}
