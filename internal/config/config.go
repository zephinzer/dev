package config

import (
	"fmt"
	"io/ioutil"

	"github.com/zephinzer/dev/pkg/utils"
	"gopkg.in/yaml.v2"
)

// Loaded contains a dictionary of configurations where the key equals
// the file's absolute path and the value containing the loaded configuration
var Loaded = map[string]Config{}

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

type Config struct {
	Dev          Dev          `json:"dev" yaml:"dev,omitempty"`
	Links        Links        `json:"links" yaml:"links,omitempty"`
	Networks     Networks     `json:"networks" yaml:"networks,omitempty"`
	Platforms    Platforms    `json:"platforms" yaml:"platforms,omitempty"`
	Softwares    Softwares    `json:"softwares" yaml:"softwares,omitempty"`
	Repositories Repositories `json:"repositories" yaml:"repositories,omitempty"`
}

// MergeWith merges the current Config configuration instance with
// a second provided :other Config configuration instance, intention for
// this is to merge a local configuration with a global configuration
func (f *Config) MergeWith(other *Config) []error {
	mergeWarnings := []error{}
	// TODO: copy over dev configurations
	if warnings := f.Dev.MergeWith(other.Dev); warnings != nil && len(warnings) > 0 {
		mergeWarnings = append(mergeWarnings, warnings...)
	}

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
	f.Platforms.Trello.MergeWith(other.Platforms.Trello)

	// copy over softwares
	f.Softwares.MergeWith(other.Softwares)

	// copy over repositories
	f.Repositories.MergeWith(other.Repositories)

	return mergeWarnings
}
