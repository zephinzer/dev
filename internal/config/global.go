package config

// Global contains the global configuration usable by all CLI commands,
// all root properties are guaranteed to exist
var Global = &Config{
	Dev:          Dev{},
	Links:        Links{},
	Networks:     Networks{},
	Platforms:    Platforms{},
	Softwares:    Softwares{},
	Repositories: Repositories{},
}
