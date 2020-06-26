package repository

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/usvc/go-config"
	c "github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/internal/prompt"
	"github.com/zephinzer/dev/pkg/utils"
	"github.com/zephinzer/dev/pkg/validator"
)

const FlagKeepGitHistory = "keep-git-history"

var conf = config.Map{
	FlagKeepGitHistory: &config.Bool{
		Shorthand: "k",
		Usage:     "when specified, retains the .git directory after cloning the template",
	},
}

func GetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     fmt.Sprintf("%s [path]", constants.RepositoryCanonicalNoun),
		Aliases: constants.RepositoryAliases,
		Short:   "Initialises a repository using a selected template",
		Run:     run,
	}
	conf.ApplyToFlagSet(cmd.Flags())
	return &cmd
}

func run(cmd *cobra.Command, args []string) {
	pathToInitialiseAt := path.Join(args...)
	log.Debug("resolving provided path %s...", pathToInitialiseAt)
	pathToInitialiseAt, resolvePathError := utils.ResolvePath(pathToInitialiseAt)
	if resolvePathError != nil {
		log.Errorf("failed to resolve path '%s': %s", pathToInitialiseAt, resolvePathError)
		os.Exit(constants.ExitErrorInput)
	}

	log.Debugf("ensuring path leads to a directory...")
	if ensureDirectoryExistsError := utils.EnsureDirectoryExists(pathToInitialiseAt); ensureDirectoryExistsError != nil {
		log.Errorf("failed to ensure directory exists at '%s': %s", pathToInitialiseAt, ensureDirectoryExistsError)
		os.Exit(constants.ExitErrorSystem | constants.ExitErrorInput)
	}

	log.Debugf("ensuring selected directory is free of objects...")
	isDirectoryEmpty, isDirectoryEmptyError := utils.IsDirectoryEmpty(pathToInitialiseAt)
	if isDirectoryEmptyError != nil {
		log.Errorf("failed to use selected directory at '%s': %s", pathToInitialiseAt, isDirectoryEmptyError)
		os.Exit(constants.ExitErrorInput | constants.ExitErrorSystem)
	} else if !isDirectoryEmpty {
		log.Errorf("directory at '%s' already contains files", pathToInitialiseAt)
		os.Exit(constants.ExitErrorValidation)
	}

	log.Debugf("processing %v templates...", len(c.Global.Dev.Repository.Templates))
	availableTemplates := []string{}
	for _, template := range c.Global.Dev.Repository.Templates {
		availableTemplates = append(availableTemplates, template.String())
	}

	log.Debug("prompting user to select template to use for repository at '%s'...", pathToInitialiseAt)
	selectedIndex, selectError := prompt.ToSelect(prompt.InputOptions{
		BeforeMessage:     "choose a repository template to use",
		SerializedOptions: availableTemplates,
	})
	if selectError != nil {
		log.Errorf("failed to get input: %s", selectError)
		os.Exit(constants.ExitErrorInput)
	}
	selectedTemplate := c.Global.Dev.Repository.Templates[selectedIndex]
	selectedURL, parseURLError := validator.ParseURL(selectedTemplate.URL)
	if parseURLError != nil {
		log.Errorf("failed to parse url '%s': %s", selectedTemplate.URL, parseURLError)
		os.Exit(constants.ExitErrorConfiguration | constants.ExitErrorValidation)
	}

	log.Debugf("cloning template from %s...\n", selectedURL.GetHTTPSString())
	gitCloneError := utils.GitClone(selectedURL.GetHTTPSString(), pathToInitialiseAt)
	if gitCloneError != nil {
		log.Errorf("failed to clone template locally: %s", gitCloneError)
		os.Exit(constants.ExitErrorInput | constants.ExitErrorSystem)
	}
	if conf.GetBool(FlagKeepGitHistory) {
		return
	}
	if removeAllError := os.RemoveAll(path.Join(pathToInitialiseAt, "/.git")); removeAllError != nil {
		log.Warnf("failed to remove .git directory: %s", removeAllError)
		os.Exit(constants.ExitErrorSystem)
	}
}
