package config

import (
	"github.com/usvc/dev/pkg/github"
	"github.com/usvc/dev/pkg/gitlab"
	"github.com/usvc/dev/pkg/network"
	"github.com/usvc/dev/pkg/pivotaltracker"
	"github.com/usvc/dev/pkg/software"
	"github.com/usvc/dev/pkg/trello"
)

type File struct {
	Dev       Dev                 `json:"dev" yaml:"dev"`
	Networks  []network.Network   `json:"network" yaml:"network"`
	Platforms Platforms           `json:"platforms" yaml:"platforms"`
	Softwares []software.Software `json:"software" yaml:"software"`
}

func (f *File) MergeWith(other *File) {
	f.Dev = other.Dev

	seenMap := map[string]bool{}
	for _, githubAccount := range f.Platforms.Github.Accounts {
		seenMap["github"+githubAccount.AccessToken] = true
	}
	for _, githubAccount := range other.Platforms.Github.Accounts {
		if seenMap["github"+githubAccount.AccessToken] == true {
			continue
		}
		f.Platforms.Github.Accounts = append(f.Platforms.Github.Accounts, githubAccount)
		seenMap["github"+githubAccount.AccessToken] = true
	}

	for _, gitlabAccount := range f.Platforms.Gitlab.Accounts {
		seenMap["gitlab"+gitlabAccount.AccessToken] = true
	}
	for _, gitlabAccount := range other.Platforms.Gitlab.Accounts {
		if seenMap["gitlab"+gitlabAccount.AccessToken] == true {
			continue
		}
		f.Platforms.Gitlab.Accounts = append(f.Platforms.Gitlab.Accounts, gitlabAccount)
		seenMap["gitlab"+gitlabAccount.AccessToken] = true
	}

	if len(f.Platforms.PivotalTracker.AccessToken) == 0 && len(other.Platforms.PivotalTracker.AccessToken) > 0 {
		f.Platforms.PivotalTracker.AccessToken = other.Platforms.PivotalTracker.AccessToken
	}
	for _, pivotalProject := range f.Platforms.PivotalTracker.Projects {
		seenMap["pivotal"+pivotalProject.ProjectID] = true
	}
	for _, pivotalProject := range other.Platforms.PivotalTracker.Projects {
		if seenMap["pivotal"+pivotalProject.ProjectID] == true {
			continue
		}
		f.Platforms.PivotalTracker.Projects = append(f.Platforms.PivotalTracker.Projects, pivotalProject)
		seenMap["pivotal"+pivotalProject.ProjectID] = true
	}

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

	for _, network := range f.Networks {
		seenMap["network"+network.Check.URL] = true
	}
	for _, network := range other.Networks {
		if seenMap["network"+network.Check.URL] == true {
			continue
		}
		f.Networks = append(f.Networks, network)
		seenMap["network"+network.Check.URL] = true
	}
}

// Platforms specifies the configurations available
type Platforms struct {
	PivotalTracker pivotaltracker.Config `json:"pivotaltracker" yaml:"pivotaltracker"`
	Github         github.Config         `json:"github" yaml:"github"`
	Gitlab         gitlab.Config         `json:"gitlab" yaml:"gitlab"`
	Trello         trello.Config         `json:"trello" yaml:"trello"`
}

func (p Platforms) GetSanitized() Platforms {
	return Platforms{
		PivotalTracker: p.PivotalTracker.GetSanitized(),
		Github:         p.Github.GetSanitized(),
		Gitlab:         p.Gitlab.GetSanitized(),
	}
}

// Dev specifies the configurations available for the CLI tool itself
type Dev struct {
	Client   DevClient   `json:"client" yaml:"client"`
	Defaults DevDefaults `json:"defaults" yaml:"defaults"`
}

type DevClient struct {
	Database DevClientDatabase `json:"database" yaml:"database"`
}

type DevClientDatabase struct {
	Path string `json:"path" yaml:"path"`
}

// DevDefaults specifies default selections when running `dev <verb> <noum> <this>` commands
type DevDefaults struct {
	// Account defines the default platform to use for account
	Account string `json:"account" yaml:"account"`
	// Notifications defines the default platform to use for notifications
	Notifications string `json:"notifications" yaml:"notifications"`
	// Work defines the default platform to use for work
	Work string `json:"work" yaml:"work"`
}
