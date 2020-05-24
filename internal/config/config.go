package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// Loaded contains a dictionary of configurations where the key equals
// the file's absolute path and the value containing the loaded configuration
var Loaded = map[string]Config{}

// NewFromFile is a convenience function that reads the configuration
// in from a file at the provided :filePath
func NewFromFile(filePath string) (*Config, error) {
	var configuration Config

	absoluteFilePath := filePath
	if !path.IsAbs(absoluteFilePath) {
		cwd, getWdError := os.Getwd()
		if getWdError != nil {
			return nil, fmt.Errorf("failed to get working directory: %s", getWdError)
		}
		absoluteFilePath = path.Join(cwd, absoluteFilePath)
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

type Config struct {
	Dev          Dev          `json:"dev" yaml:"dev"`
	Includes     Includes     `json:"includes" yaml:"includes"`
	Links        Links        `json:"links" yaml:"links"`
	Networks     Networks     `json:"networks" yaml:"networks"`
	Platforms    Platforms    `json:"platforms" yaml:"platforms"`
	Softwares    Softwares    `json:"softwares" yaml:"softwares"`
	Repositories Repositories `json:"repositories" yaml:"repositories"`
}

// MergeWith merges the current Config configuration instance with
// a second provided :other Config configuration instance, intention for
// this is to merge a local configuration with a global configuration
func (f *Config) MergeWith(other *Config) {
	seenMap := map[string]bool{}

	// TODO: copy over dev configurations

	// copy over links
	f.Links.MergeWith(other.Links)

	// copy over networks
	f.Networks.MergeWith(other.Networks)

	// copy over platform configurations
	/// copy over github configuration
	f.Platforms.Github.MergeWith(other.Platforms.Github)

	/// copy over gitlab configuration
	f.Platforms.Gitlab.MergeWith(other.Platforms.Gitlab)

	/// copy over pivotal tracker configuration
	f.Platforms.PivotalTracker.MergeWith(other.Platforms.PivotalTracker)

	/// copy over trello configuration
	if len(f.Platforms.Trello.AccessToken) == 0 && len(other.Platforms.Trello.AccessToken) > 0 {
		f.Platforms.Trello.AccessToken = other.Platforms.Trello.AccessToken
	}
	if len(f.Platforms.Trello.AccessKey) == 0 && len(other.Platforms.Trello.AccessKey) > 0 {
		f.Platforms.Trello.AccessKey = other.Platforms.Trello.AccessKey
	}
	for _, trelloBoard := range f.Platforms.Trello.Boards {
		seenMap["trello"+trelloBoard.ID] = true
	}
	for _, trelloBoard := range other.Platforms.Trello.Boards {
		if seenMap["trello"+trelloBoard.ID] == true {
			continue
		}
		f.Platforms.Trello.Boards = append(f.Platforms.Trello.Boards, trelloBoard)
		seenMap["trello"+trelloBoard.ID] = true
	}

	// copy over softwares
	for _, software := range f.Softwares {
		seenMap["software"+software.Check.Command[0]] = true
	}
	for _, software := range other.Softwares {
		if seenMap["software"+software.Check.Command[0]] == true {
			continue
		}
		f.Softwares = append(f.Softwares, software)
		seenMap["software"+software.Check.Command[0]] = true
	}

	// copy over repositories
	for _, repository := range f.Repositories {
		repositoryPath, _ := repository.GetPath()
		seenMap["repo"+repositoryPath] = true
	}
	for _, repository := range other.Repositories {
		repositoryPath, _ := repository.GetPath()
		if seenMap["repo"+repositoryPath] == true {
			continue
		}
		f.Repositories = append(f.Repositories, repository)
		seenMap["repo"+repositoryPath] = true
	}
}
