package config

// Loaded contains a dictionary of configurations where the key equals
// the file's absolute path and the value containing the loaded configuration
var Loaded = map[string]Config{}

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
