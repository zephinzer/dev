package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zephinzer/dev/cmd/dev/_/cmdutils"
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/git"
	"github.com/zephinzer/dev/internal/log"
	. "github.com/zephinzer/dev/internal/repository"
	"github.com/zephinzer/dev/pkg/utils/str"
	"github.com/zephinzer/dev/pkg/validator"
	"gopkg.in/yaml.v2"
)

const (
	Example string = `
1. Adding a single repository using URL:
    dev add repository https://github.com/zephinzer/dev

2. Adding a single repository using HTTPS clone URL:
    dev add repository https://github.com/zephinzer/dev.git

3. Adding a single repository using SSH clone URL:
    dev add repository git@github.com/zephinzer/dev.git

4. Adding multiple repositories:
    dev add repository https://github.com/zephinzer/dev https://github.com/zephinzer/godev
`
)

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     fmt.Sprintf("%s [flags] <repository_urls...>", constants.RepositoryCanonicalNoun),
		Aliases: constants.RepositoryAliases,
		Example: strings.Trim(Example, "\n\r\t"),
		Short:   "add a repository to your configuration",
		Long:    "add a repository to your configuration (use `dev add this repo` to add the current repository you're on)",
		PreRun:  preRun,
		Run:     run,
	}
	return &cmd
}

func preRun(command *cobra.Command, args []string) {
	configCount := 0
	for w := range config.Loaded {
		log.Info(w)
		configCount++
	}
	if configCount == 0 {
		cmdutils.ExitWithProblem("this command requires an existing configuration file. create a default one by running `touch ~/.dev.yaml` before trying again", constants.ExitErrorUser)
	}
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
	loadedReposMap := cmdutils.GetLoadedRepositories()
	loadedReposMapCount := 1
	log.Debug("repositories already present:")
	for repoPath, configPath := range loadedReposMap {
		log.Debugf("%v. %s from '%s'", loadedReposMapCount, repoPath, configPath)
		loadedReposMapCount++
	}

	// filter out repositories that are already in the configuration
	filteredRepoURLs := cmdutils.FilterOutLoadedRepositories(targetRepoURLs, loadedReposMap)
	if len(filteredRepoURLs) == 0 {
		log.Info("no repositories to work on, exiting with status 0 now")
		os.Exit(constants.ExitOK)
	}
	log.Infof("will be adding following repositories: %v", strings.Join(filteredRepoURLs, ", "))

	// let's do dis
	for _, repoURL := range filteredRepoURLs {
		parsedURL, parseURLError := validator.ParseURL(repoURL)
		if parseURLError != nil {
			log.Warnf("skipping '%s', failed to get ssh git clone url: %s", repoURL, parseURLError)
			continue
		}

		// set configuration file
		configurationPath, err := config.PromptSelectLoadedConfiguration(
			fmt.Sprintf("which configuration file should we add '%s' to?", repoURL),
		)
		if err != nil {
			log.Errorf("failed to get a valid configuration file: %s", err)
			continue
		}
		if str.IsEmpty(configurationPath) {
			log.Infof("skipping adding of repo '%s'", repoURL)
			continue
		}
		log.Infof("adding repo '%s' to configuration at '%s'...", repoURL, configurationPath)

		targetConfiguration := config.Loaded[configurationPath]
		targetRepository := Repository{}
		targetRepository.SetURL(parsedURL.GetSSHString())

		// set name
		namePromptError := targetRepository.PromptForName()
		if namePromptError != nil {
			log.Warnf("failed to set a name for repo '%s': %s", repoURL, namePromptError)
			continue
		}
		log.Infof("using '%s' as the repository name", targetRepository.Name)

		// set description
		descriptionPromptError := targetRepository.PromptForDescription()
		if descriptionPromptError != nil {
			log.Warnf("failed to set a description for repo '%s': %s", repoURL, descriptionPromptError)
			continue
		}
		log.Infof("using '%s' as the repository description", targetRepository.Description)

		// set workspaces
		availableWorkspaces := targetConfiguration.Repositories.GetWorkspaces()
		log.Infof("existing workspaces: %s", strings.Join(availableWorkspaces, ", "))
		workspacePromptError := targetRepository.PromptForWorkspaces()
		if workspacePromptError != nil {
			log.Warnf("failed to set workspaces for repo '%s': %s", repoURL, workspacePromptError)
			continue
		}
		log.Infof("using [%s] as the repository workspaces", strings.Join(targetRepository.Workspaces, ", "))

		repo := targetRepository.ToRepository()
		targetConfiguration.Repositories = append(targetConfiguration.Repositories, repo)
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

		homeDir := cmdutils.GetHomeDirectory()
		repoPath, getPathError := repo.GetPath(homeDir)
		if getPathError != nil {
			log.Errorf("failed to get path of repository: %s", getPathError)
			os.Exit(constants.ExitErrorApplication | constants.ExitErrorInput)
		}
		if gitCloneError := git.Clone(parsedURL.GetSSHString(), repoPath); gitCloneError != nil {
			log.Errorf("failed to clone new repository: %s", gitCloneError)
			os.Exit(constants.ExitErrorApplication | constants.ExitErrorInput)
		}
	}
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
	repositoryURLs = str.Dedupe(repositoryURLs)
	return repositoryURLs
}
