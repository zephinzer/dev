package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DevConfigTests struct {
	suite.Suite
}

func TestDevConfig(t *testing.T) {
	suite.Run(t, &DevConfigTests{})
}

func (s *DevConfigTests) TestDev() {
	dev := []Dev{
		{ // golden config
			Client: DevClient{
				Database: DevClientDatabase{
					Path: "__database_path",
				},
				Notifications: DevClientNotifications{
					Telegram: DevClientNotificationsTelegram{
						Token: "__telegram_token",
						ID:    "__telegram_id",
					},
				},
				Platforms: DevClientPlatforms{
					Github: DevClientPlatformsGithub{
						ClientID:     "__github_client_id",
						ClientSecret: "__github_client_secret",
						RedirectURI:  "__github_redirect_uri",
					},
				},
			},
			Repository: DevRepository{
				Templates: []DevRepositoryTemplate{
					{
						Name: "__name_1",
						URL:  "__url_1",
						Path: "__path/to/1",
					},
					{
						Name: "__name_2",
						URL:  "__url_2",
						Path: "__path/to/2",
					},
					{
						Name: "__name_3",
						URL:  "__url_3",
						Path: "__path/to/3",
					},
				},
			},
		},
		{ // empty string set
			Client: DevClient{
				Database: DevClientDatabase{
					Path: "",
				},
				Notifications: DevClientNotifications{
					Telegram: DevClientNotificationsTelegram{
						Token: "",
						ID:    "",
					},
				},
				Platforms: DevClientPlatforms{
					Github: DevClientPlatformsGithub{
						ClientID:     "",
						ClientSecret: "",
						RedirectURI:  "",
					},
				},
			},
			Repository: DevRepository{
				Templates: []DevRepositoryTemplate{
					{
						Name: "__name_0",
						URL:  "__url_0",
						Path: "__path/to/0",
					}},
			},
		},
		{ // zero-value set
			Client:     DevClient{},
			Repository: DevRepository{},
		},
		{}, // no fucks given
		{ // half a fuck provided
			Client: DevClient{
				Database: DevClientDatabase{
					Path: "/",
				},
				Notifications: DevClientNotifications{
					Telegram: DevClientNotificationsTelegram{
						ID: "__not_empty",
					},
				},
				Platforms: DevClientPlatforms{
					Github: DevClientPlatformsGithub{
						ClientID: "__not_empty",
					},
				},
			},
		},
	}
	errs := dev[1].MergeWith(dev[0])
	s.Len(errs, 0)
	s.Len(dev[1].Repository.Templates, 4, "repository templates are always merged")
	s.True(len(dev[1].Client.Platforms.Github.ClientID) > 0, "github tokens are copied over if an empty string is detected")
	s.True(len(dev[1].Client.Notifications.Telegram.ID) > 0, "telegram tokens are copied over if an empty string is detected")
	s.True(len(dev[1].Client.Database.Path) > 0, "database paths are copied over if an empty string is detected")

	errs = dev[2].MergeWith(dev[0])
	s.Len(errs, 0)
	s.Len(dev[2].Repository.Templates, 3, "repository templates are copied over if .Repository.Templates is not defined")
	s.True(len(dev[2].Client.Platforms.Github.ClientID) > 0, "github tokens are copied over when the .Client property is not defined explicitly")
	s.True(len(dev[2].Client.Notifications.Telegram.ID) > 0, "telegram tokens are copied over when the .Client property is not defined explicitly")
	s.True(len(dev[2].Client.Database.Path) > 0, "database paths are copied over when the .Client property is not defined explicitly")

	errs = dev[3].MergeWith(dev[0])
	s.Len(errs, 0)
	s.Len(dev[3].Repository.Templates, 3, "repository templates are copied over")
	s.True(len(dev[3].Client.Platforms.Github.ClientID) > 0, "github tokens are copied over")
	s.True(len(dev[3].Client.Notifications.Telegram.ID) > 0, "telegram tokens are copied over")
	s.True(len(dev[3].Client.Database.Path) > 0, "database paths are copied over")

	errs = dev[4].MergeWith(dev[0])
	s.Len(errs, 3)
	s.Len(dev[4].Repository.Templates, 3, "repository templates are copied over if .Repository.Templates is not defined")
	s.Len(dev[4].Client.Platforms.Github.ClientSecret, 0, "github secret is NOT copied over if client id is already defined")
	s.Len(dev[4].Client.Notifications.Telegram.Token, 0, "telegram token is NOT copied over if id is already defined")
	s.Equal("/", dev[4].Client.Database.Path, "database path is not overwritten")

}
