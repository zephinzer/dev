package config

import (
	"github.com/zephinzer/dev/internal/github"
	"github.com/zephinzer/dev/internal/gitlab"
	"github.com/zephinzer/dev/internal/pivotaltracker"
	"github.com/zephinzer/dev/internal/trello"
)

// Platforms specifies the configurations available
type Platforms struct {
	PivotalTracker pivotaltracker.Config `json:"pivotaltracker" yaml:"pivotaltracker,omitempty"`
	Github         github.Config         `json:"github" yaml:"github,omitempty"`
	Gitlab         gitlab.Config         `json:"gitlab" yaml:"gitlab,omitempty"`
	Trello         trello.Config         `json:"trello" yaml:"trello,omitempty"`
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
