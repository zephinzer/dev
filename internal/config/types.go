package config

import (
	"github.com/usvc/dev/pkg/github"
	"github.com/usvc/dev/pkg/gitlab"
	"github.com/usvc/dev/pkg/pivotaltracker"
	"github.com/usvc/dev/pkg/software"
)

type File struct {
	Platforms Platforms           `yaml:"platforms"`
	Softwares []software.Software `yaml:"software"`
}

type Platforms struct {
	PivotalTracker pivotaltracker.Config `yaml:"pivotaltracker"`
	Github         github.Config         `yaml:"github"`
	Gitlab         gitlab.Config         `yaml:"gitlab"`
}
