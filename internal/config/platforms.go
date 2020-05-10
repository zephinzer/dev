package config

import (
	"github.com/usvc/dev/pkg/github"
	"github.com/usvc/dev/pkg/gitlab"
	"github.com/usvc/dev/pkg/pivotaltracker"
	"github.com/usvc/dev/pkg/trello"
)

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
