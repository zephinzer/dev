package config

import (
	"github.com/usvc/dev/pkg/github"
	"github.com/usvc/dev/pkg/gitlab"
	"github.com/usvc/dev/pkg/pivotaltracker"
	"github.com/usvc/dev/pkg/software"
)

type File struct {
	Dev       Dev                 `yaml:"dev"`
	Platforms Platforms           `yaml:"platforms"`
	Softwares []software.Software `yaml:"software"`
}

// Platforms specifies the configurations available
type Platforms struct {
	PivotalTracker pivotaltracker.Config `yaml:"pivotaltracker"`
	Github         github.Config         `yaml:"github"`
	Gitlab         gitlab.Config         `yaml:"gitlab"`
}

// Dev specifies the configurations available for the CLI tool itself
type Dev struct {
	Client struct {
		Database struct {
			Path string `yaml:"path"`
		} `yaml:"database"`
	} `yaml:"client"`
	Defaults DevDefaults `yaml:"defaults"`
}

// DevDefaults specifies default selections when running `dev <verb> <noum> <this>` commands
type DevDefaults struct {
	// Account defines the default platform to use for account
	Account string `yaml:"account"`
	// Notifications defines the default platform to use for notifications
	Notifications string `yaml:"notifications"`
	// Work defines the default platform to use for work
	Work string `yaml:"work"`
}
