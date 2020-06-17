package config

import (
	"errors"

	"github.com/zephinzer/dev/pkg/utils"
)

// Dev specifies the configurations available for the CLI tool itself
type Dev struct {
	Client DevClient `json:"client" yaml:"client,omitempty"`
}

func (d *Dev) MergeWith(o Dev) []error {
	return d.Client.MergeWith(o.Client)
}

// DevClient holds configurations related to the CLI tool
type DevClient struct {
	Database      DevClientDatabase      `json:"database" yaml:"database,omitempty"`
	Notifications DevClientNotifications `json:"notifications" yaml:"notifications,omitempty"`
	Platforms     DevClientPlatforms     `json:"platforms" yaml:"platforms,omitempty"`
}

func (dc *DevClient) MergeWith(o DevClient) []error {
	warnings := []error{}
	warnings = append(warnings, dc.Database.MergeWith(o.Database)...)
	warnings = append(warnings, dc.Notifications.MergeWith(o.Notifications)...)
	warnings = append(warnings, dc.Platforms.MergeWith(o.Platforms)...)
	return warnings
}

// DevClientDatabase holds configurations related to the data persistence
// mechanism of the CLI tool
type DevClientDatabase struct {
	Path string `json:"path" yaml:"path,omitempty"`
}

func (dcdb *DevClientDatabase) MergeWith(o DevClientDatabase) []error {
	if len(dcdb.Path) > 0 {
		return []error{errors.New("dev.client.database.path already set")}
	}
	dcdb.Path = o.Path
	return nil
}

// DevClientNotifications holds configurations related to the notifications
// mechanisms of the CLI tool
type DevClientNotifications struct {
	Telegram DevClientNotificationsTelegram `json:"telegram" yaml:"telegram,omitempty"`
}

func (dcn *DevClientNotifications) MergeWith(o DevClientNotifications) []error {
	return dcn.Telegram.MergeWith(o.Telegram)
}

// DevClientNotificationsTelegram holds configurations related to the
// telegram integration for sending notifications
type DevClientNotificationsTelegram struct {
	Token string `json:"token" yaml:"token,omitempty"`
	ID    string `json:"id" yaml:"id,omitempty"`
}

func (dcntg *DevClientNotificationsTelegram) MergeWith(o DevClientNotificationsTelegram) []error {
	if len(dcntg.Token) > 0 {
		return []error{errors.New("dev.client.notifications.telegram.token already set")}
	}
	dcntg.Token = o.Token
	dcntg.ID = o.ID
	return nil
}

// DevClientPlatforms holds configurations related to the platform integrations
// for individual users
type DevClientPlatforms struct {
	Github DevClientPlatformsGithub `json:"github" yaml:"github,omitempty"`
}

func (dcp *DevClientPlatforms) MergeWith(o DevClientPlatforms) []error {
	warnings := []error{}
	warnings = append(warnings, dcp.Github.MergeWith(o.Github)...)
	return warnings
}

type DevClientPlatformsGithub struct {
	ClientID     string `json:"clientId" yaml:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret" yaml:"clientSecret,omitempty"`
	RedirectURI  string `json:"redirectUri" yaml:"redirectUri,omitempty"`
}

func (dcpgh *DevClientPlatformsGithub) MergeWith(o DevClientPlatformsGithub) []error {
	var warnings []error
	if !utils.IsEmptyString(dcpgh.ClientID) && dcpgh.ClientID != o.ClientID {
		warnings = append(warnings, errors.New("dev.client.platforms.github.clientID already set"))
	}
	if !utils.IsEmptyString(dcpgh.ClientSecret) && dcpgh.ClientSecret != o.ClientSecret {
		warnings = append(warnings, errors.New("dev.client.platforms.github.clientSecret already set"))
	}
	if !utils.IsEmptyString(dcpgh.RedirectURI) && dcpgh.RedirectURI != o.RedirectURI {
		warnings = append(warnings, errors.New("dev.client.platforms.github.redirectUri already set"))
	}
	if len(warnings) > 0 {
		return warnings
	}
	dcpgh.ClientID = o.ClientID
	dcpgh.ClientSecret = o.ClientSecret
	dcpgh.RedirectURI = o.RedirectURI
	return nil
}
