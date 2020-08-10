package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/github"
	"github.com/zephinzer/dev/internal/gitlab"
	"github.com/zephinzer/dev/internal/pivotaltracker"
	"github.com/zephinzer/dev/internal/trello"
)

type PlatformsTests struct {
	suite.Suite
}

func TestPlatform(t *testing.T) {
	suite.Run(t, &PlatformsTests{})
}

func (s *PlatformsTests) Test_GetSanitized() {
	platforms := Platforms{
		PivotalTracker: pivotaltracker.Config{
			AccessToken: "__pivotal_access_token",
			Projects: pivotaltracker.Projects{
				{
					AccessToken: "__pivotal_project_access_token",
					Public:      true,
				},
			},
		},
		Github: github.Config{
			Accounts: github.AccountConfigs{
				{
					AccessToken: "__github_access_token",
					Public:      true,
				},
			},
		},
		Gitlab: gitlab.Config{
			Accounts: gitlab.AccountConfigs{
				{
					AccessToken: "__gitlab_access_token",
					Public:      true,
				},
			},
		},
		Trello: trello.Config{
			AccessKey:   "__trello_access_key",
			AccessToken: "__trello_access_token",
			Boards: trello.Boards{
				{
					AccessKey:   "__trello_board_access_key",
					AccessToken: "__trello_board_access_token",
					Public:      true,
				},
			},
		},
	}
	sanitized := platforms.GetSanitized()
	s.Equal(constants.DefaultRedactedString, sanitized.PivotalTracker.AccessToken)
	s.Equal(constants.DefaultRedactedString, sanitized.PivotalTracker.Projects[0].AccessToken)
	s.Equal(constants.DefaultRedactedString, sanitized.Github.Accounts[0].AccessToken)
	s.Equal(constants.DefaultRedactedString, sanitized.Gitlab.Accounts[0].AccessToken)
	s.Equal(constants.DefaultRedactedString, sanitized.Trello.AccessKey)
	s.Equal(constants.DefaultRedactedString, sanitized.Trello.AccessToken)
	s.Equal(constants.DefaultRedactedString, sanitized.Trello.Boards[0].AccessKey)
	s.Equal(constants.DefaultRedactedString, sanitized.Trello.Boards[0].AccessToken)
}
