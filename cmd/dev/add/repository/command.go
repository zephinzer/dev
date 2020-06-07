package repository

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/repository"
	"github.com/zephinzer/dev/pkg/utils"
	"github.com/zephinzer/dev/pkg/validator"
)

const (
	Example string = `
  Adding a single repository:
    dev add repository https://github.com/zephinzer/dev

  Adding multiple repositories:
    dev add repository https://github.com/zephinzer/dev https://github.com/zephinzer/godev
`
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     fmt.Sprintf("%s [flags] <repository_urls...>", constants.RepositoryCanonicalNoun),
		Aliases: constants.RepositoryAliases,
		Example: strings.Trim(Example, "\n\r\t"),
		Short:   "add a repository to your configuration",
		Run:     run,
	}
	return &cmd
}

func run(command *cobra.Command, args []string) {
	if len(args) == 0 {
		command.Help()
		log.Errorf("no repository url(s) specified, see usage above for how to use this command")
		os.Exit(1)
		return
	}

	// parse arguments into a list of repository urls
	repositoryURLsToAdd := parseArgumentsIntoURLs(args)
	log.Info("repositories to add:\n")
	for index, repositoryURL := range repositoryURLsToAdd {
		log.Infof(fmt.Sprintf("%v. %s\n", index+1, repositoryURL))
	}

	// parse loaded configurations into a map of set repositories
	alreadyPresentRepositories := map[string]string{}
	loadedConfigurationPaths := getLoadedConfigFilePaths(config.Loaded)
	log.Info("loaded configuration file paths:\n")
	for index, loadedConfigurationPath := range loadedConfigurationPaths {
		repos := config.Loaded[loadedConfigurationPath].Repositories
		log.Infof("%v. %s (repo count: %v)\n", index+1, loadedConfigurationPath, len(repos))
		for _, repo := range repos {
			repoPath, getPathError := repo.GetPath()
			if getPathError != nil {
				log.Warnf("failed to get repository path for '%s': %s", repo.URL, getPathError)
				continue
			}
			alreadyPresentRepositories[repoPath] = loadedConfigurationPath
		}
	}

	// filter out repositories that are already in the configuration
	finalRepositoryURLsToAdd := []string{}
	for _, repoURL := range repositoryURLsToAdd {
		repo := repository.Repository{URL: repoURL}
		repoPath, getPathError := repo.GetPath()
		if getPathError != nil {
			log.Warnf("failed to get repository path for '%s': %s", repo.URL, getPathError)
			continue
		}
		if value, ok := alreadyPresentRepositories[repoPath]; ok && len(value) > 0 {
			log.Warnf("repository '%s' already configured from '%s'", repoURL, value)
			continue
		}
		finalRepositoryURLsToAdd = append(finalRepositoryURLsToAdd, repoURL)
	}

	for _, repoURL := range finalRepositoryURLsToAdd {
		log.Infof("processing repository url '%s'...", repoURL)
	}
}

func getLoadedConfigFilePaths(theMap map[string]config.Config) []string {
	output := []string{}
	for key := range theMap {
		output = append(output, key)
	}
	return output
}

func parseArgumentsIntoURLs(args []string) []string {
	repositoryURLs := []string{}
	for _, arg := range args {
		log.Tracef("parsing url '%s'...", arg)
		repositoryURL, parseURLError := validator.ParseURL(arg)
		if parseURLError != nil {
			log.Warnf("failed to parse url '%s': %s", arg, parseURLError)
			continue
		}
		repositoryURLs = append(repositoryURLs, repositoryURL.String())
	}
	repositoryURLs = utils.DedupeStrings(repositoryURLs)
	return repositoryURLs
}
