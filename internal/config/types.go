package config

import (
	"github.com/usvc/dev/internal/link"
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
	Links     []link.Link         `json:"links" yaml:"links"`
}

// MergeWith merges the current File configuration instance with
// a second provided :other File configuration instance, intention for
// this is to merge a local configuration with a global configuration
func (f *File) MergeWith(other *File) {
	f.Dev = other.Dev

	seenMap := map[string]bool{}

	// copy over github configuration
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

	// copy over gitlab configuration
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

	// copy over pivotal tracker configuration
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

	// copy over trello configuration
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

	// copy over networks
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

// GetSanitized returns a sanitised deep copy of the current Platforms instance
func (p Platforms) GetSanitized() Platforms {
	return Platforms{
		PivotalTracker: p.PivotalTracker.GetSanitized(),
		Github:         p.Github.GetSanitized(),
		Gitlab:         p.Gitlab.GetSanitized(),
		Trello:         p.Trello.GetSanitized(),
	}
}

// Dev specifies the configurations available for the CLI tool itself
type Dev struct {
	Client   DevClient   `json:"client" yaml:"client"`
	Defaults DevDefaults `json:"defaults" yaml:"defaults"`
}

// DevClient holds configurations related to the CLI tool
type DevClient struct {
	Database      DevClientDatabase      `json:"database" yaml:"database"`
	Notifications DevClientNotifications `json:"notifications" yaml:"notifications"`
}

// DevClientDatabase holds configurations related to the data persistence
// mechanism of the CLI tool
type DevClientDatabase struct {
	Path string `json:"path" yaml:"path"`
}

// DevClientNotifications holds configurations related to the notifications
// mechanisms of the CLI tool
type DevClientNotifications struct {
	Telegram DevClientNotificationsTelegram `json:"telegram" yaml:"telegram"`
}

// DevClientNotificationsTelegram holds configurations related to the
// telegram integration for sending notifications
type DevClientNotificationsTelegram struct {
	Token string `json:"token" yaml:"token"`
	ID    string `json:"id" yaml:"id"`
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
