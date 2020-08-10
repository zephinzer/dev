package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/zephinzer/dev/pkg/utils"
	"gopkg.in/yaml.v2"
)

const (
	// RegexpFileName defines a regular expression pattern which is used
	// to decide whether a file is a `dev` configuration file based on
	// the file name
	RegexpFileName = `^\.dev(?P<label>\.[a-zA-Z0-9\-\_]+)*\.y(a)?ml$`
)

// FilterConfigurations accepts a list of `os.FileInfo` and returns
// a list of `os.FileInfo`s whose file names comply to the configuration
// file name pattern as defined by RegexpFileName
func FilterConfigurations(fileInfos []os.FileInfo) []os.FileInfo {
	configurations := []os.FileInfo{}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}
		fileName := fileInfo.Name()
		regex := regexp.MustCompile(RegexpFileName)
		if regex.Match([]byte(fileName)) {
			configurations = append(configurations, fileInfo)
		}
	}
	return configurations
}

// GetFilesOptions enables parameterisation of the GetFiles function by differentiating
// between the contextual configuration and the global configuration
type GetFilesOptions struct {
	// ContextualWorkspacePath defines where to draw context specific configuration from
	ContextualWorkspacePath string
	// GlobalWorkspacePath defines where to draw globval configuration from
	GlobalWorkspacePath string
}

// GetFiles returns a list of absolute file paths corresponding to configuration files
// found in 1. the current user's home directory and 2. the current working directory
func GetFiles(opt ...GetFilesOptions) ([]string, error) {
	configurationFiles := []string{}

	// local configurations
	contextualConfigPath, getWdError := os.Getwd()
	if getWdError != nil {
		return nil, fmt.Errorf("failed to get the working directory: %s", getWdError)
	}
	if len(opt) > 0 && len(opt[0].ContextualWorkspacePath) > 0 {
		contextualConfigPath = opt[0].ContextualWorkspacePath
	}
	contextualConfigs, readDirError := ioutil.ReadDir(contextualConfigPath)
	if readDirError != nil {
		return nil, fmt.Errorf("failed to get the listing of directory '%s': %s", contextualConfigPath, readDirError)
	}
	contextualConfigs = FilterConfigurations(contextualConfigs)
	for _, contextualConfig := range contextualConfigs {
		configurationFiles = append(configurationFiles, path.Join(contextualConfigPath, contextualConfig.Name()))
	}

	// global configurations
	globalConfigPath, userHomeDirError := os.UserHomeDir()
	if userHomeDirError != nil {
		return nil, fmt.Errorf("failed to get the user's home directory: %s", userHomeDirError)
	}
	if len(opt) > 0 && len(opt[0].GlobalWorkspacePath) > 0 {
		globalConfigPath = opt[0].GlobalWorkspacePath
	}
	globalConfigs, readDirError := ioutil.ReadDir(globalConfigPath)
	if readDirError != nil {
		return nil, fmt.Errorf("failed to get the listing of directory '%s': %s", contextualConfigPath, readDirError)
	}
	globalConfigs = FilterConfigurations(globalConfigs)
	for _, globalConfig := range globalConfigs {
		configurationFiles = append(configurationFiles, path.Join(globalConfigPath, globalConfig.Name()))
	}
	return configurationFiles, nil
}

// NewFromFile is a convenience function that reads the configuration
// in from a file at the provided :filePath
func NewFromFile(filePath string) (*Config, error) {
	var configuration Config

	absoluteFilePath, resolvePathError := utils.ResolvePath(filePath)
	if resolvePathError != nil {
		return nil, fmt.Errorf("failed to resolve path '%s': %s", filePath, resolvePathError)
	}

	if _, ok := Loaded[absoluteFilePath]; ok {
		return nil, fmt.Errorf("skipped loading configuration at '%s' because it's already been loaded", absoluteFilePath)
	}

	configFile, readFileError := ioutil.ReadFile(absoluteFilePath)
	if readFileError != nil {
		return nil, readFileError
	}

	if unmarshalError := yaml.Unmarshal(configFile, &configuration); unmarshalError != nil {
		return nil, unmarshalError
	}

	Loaded[absoluteFilePath] = configuration
	return &configuration, nil
}
