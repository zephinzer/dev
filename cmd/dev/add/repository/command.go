package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	. "github.com/zephinzer/dev/internal/repository"
	"github.com/zephinzer/dev/pkg/repository"
	"github.com/zephinzer/dev/pkg/utils"
	"github.com/zephinzer/dev/pkg/validator"
	"gopkg.in/yaml.v2"
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
	targetRepoURLs := getRepoURLsFromArguments(args)
	log.Infof("requested to add the repositories: %v", strings.Join(targetRepoURLs, ", "))

	// parse loaded configurations into a map of set repositories
	loadedReposMap := getLoadedRepositories()
	loadedReposMapCount := 1
	log.Debug("repositories already present:")
	for repoPath, configPath := range loadedReposMap {
		log.Debugf("%v. %s from '%s'", loadedReposMapCount, repoPath, configPath)
		loadedReposMapCount++
	}

	// filter out repositories that are already in the configuration
	filteredRepoURLs := filterOutLoadedRepositories(targetRepoURLs, loadedReposMap)
	if len(filteredRepoURLs) == 0 {
		log.Info("no repositories to work on, exiting with status 0 now")
		os.Exit(constants.ExitOK)
	}
	log.Infof("will be adding following repositories: %v", strings.Join(filteredRepoURLs, ", "))

	// let's do dis
	for _, repoURL := range filteredRepoURLs {
		configurationPath := config.PromptSelectLoadedConfiguration(
			fmt.Sprintf("which configuration file should we add '%s' to?", repoURL),
		)
		if utils.IsEmptyString(configurationPath) {
			log.Infof("skipping adding of repo '%s'", repoURL)
			continue
		}

		targetConfiguration := config.Loaded[configurationPath]
		availableWorkspaces := targetConfiguration.Repositories.GetWorkspaces()
		targetRepository := Repository{}
		targetRepository.SetURL(repoURL)
		parsedURL, parseURLError := validator.ParseURL(repoURL)
		if parseURLError != nil {
			log.Warnf("failed to get ssh git clone url for '%s': %s", repoURL, parseURLError)
		} else {
			targetRepository.SetURL(parsedURL.GetSSHString())
		}
		log.Infof("adding repo '%s' to configuration at '%s'...", targetRepository.URL, configurationPath)

		// set name
		fmt.Println("")
		namePromptError := targetRepository.PromptForName()
		if namePromptError != nil {
			log.Warnf("failed to set a name for repo '%s': %s", repoURL, namePromptError)
			continue
		}
		log.Infof("using '%s' as the repository name", targetRepository.Name)

		// set description
		fmt.Println("")
		descriptionPromptError := targetRepository.PromptForDescription()
		if descriptionPromptError != nil {
			log.Warnf("failed to set a description for repo '%s': %s", repoURL, descriptionPromptError)
			continue
		}
		log.Infof("using '%s' as the repository description", targetRepository.Description)

		// set workspaces
		fmt.Println("")
		log.Infof("existing workspaces: %s", strings.Join(availableWorkspaces, ", "))
		workspacePromptError := targetRepository.PromptForWorkspaces()
		if workspacePromptError != nil {
			log.Warnf("failed to set workspaces for repo '%s': %s", repoURL, workspacePromptError)
			continue
		}
		log.Infof("using [%s] as the repository workspaces", strings.Join(targetRepository.Workspaces, ", "))

		targetConfiguration.Repositories = append(targetConfiguration.Repositories, targetRepository.ToRepository())
		targetConfiguration.Repositories.Sort()

		configuration, marshalError := yaml.Marshal(targetConfiguration)
		if marshalError != nil {
			log.Errorf("failed to marshal the configuration to yaml: %s", marshalError)
			os.Exit(constants.ExitErrorApplication)
		}

		writeFileError := ioutil.WriteFile(configurationPath, configuration, os.ModePerm)
		if writeFileError != nil {
			log.Errorf("failed to write the new configuration to file at '%s': %s", configurationPath, writeFileError)
			os.Exit(constants.ExitErrorApplication | constants.ExitErrorSystem)
		}
	}
}

func filterOutLoadedRepositories(toAdd []string, alreadyPresent map[string]string) []string {
	finalRepositoryURLsToAdd := []string{}
	for _, repoURL := range toAdd {
		repo := repository.Repository{URL: repoURL}
		repoPath, getPathError := repo.GetPath()
		if getPathError != nil {
			log.Warnf("failed to get repository path for '%s': %s", repo.URL, getPathError)
			continue
		}
		if value, ok := alreadyPresent[repoPath]; ok && len(value) > 0 {
			log.Warnf("repository '%s' already configured from '%s'", repoURL, value)
			continue
		}
		finalRepositoryURLsToAdd = append(finalRepositoryURLsToAdd, repoURL)
	}
	return finalRepositoryURLsToAdd
}

// getLoadedRepositories retrieves repositories that
// have already been loaded and returns a `map[string]string` where the
// key value equals the local path of the repository and the value equals
// the full absolute path of the configuration file
func getLoadedRepositories() map[string]string {
	alreadyPresentRepositories := map[string]string{}
	loadedConfigurationPaths := getLoadedConfigFilePaths(config.Loaded)
	for _, loadedConfigurationPath := range loadedConfigurationPaths {
		repos := config.Loaded[loadedConfigurationPath].Repositories
		for _, repo := range repos {
			repoPath, getPathError := repo.GetPath()
			if getPathError != nil {
				log.Warnf("failed to get repository path for '%s': %s", repo.URL, getPathError)
				continue
			}
			alreadyPresentRepositories[repoPath] = loadedConfigurationPath
		}
	}
	return alreadyPresentRepositories
}

func getLoadedConfigFilePaths(theMap map[string]config.Config) []string {
	output := []string{}
	for key := range theMap {
		output = append(output, key)
	}
	return output
}

func getRepoURLsFromArguments(args []string) []string {
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
