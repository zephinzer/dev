package client

import (
	"errors"

	"github.com/zephinzer/dev/pkg/utils/str"
)

// Platforms holds configurations related to the platform integrations
// for individual users
type Platforms struct {
	Github PlatformsGithub `json:"github" yaml:"github,omitempty"`
}

func (dcp *Platforms) MergeWith(o Platforms) []error {
	warnings := []error{}
	warnings = append(warnings, dcp.Github.MergeWith(o.Github)...)
	return warnings
}

type PlatformsGithub struct {
	ClientID     string `json:"clientId" yaml:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret" yaml:"clientSecret,omitempty"`
	RedirectURI  string `json:"redirectUri" yaml:"redirectUri,omitempty"`
}

func (dcpgh *PlatformsGithub) MergeWith(o PlatformsGithub) []error {
	var warnings []error
	if !str.IsEmpty(dcpgh.ClientID) && dcpgh.ClientID != o.ClientID {
		warnings = append(warnings, errors.New("dev.client.platforms.github.clientID already set"))
	}
	if !str.IsEmpty(dcpgh.ClientSecret) && dcpgh.ClientSecret != o.ClientSecret {
		warnings = append(warnings, errors.New("dev.client.platforms.github.clientSecret already set"))
	}
	if !str.IsEmpty(dcpgh.RedirectURI) && dcpgh.RedirectURI != o.RedirectURI {
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
