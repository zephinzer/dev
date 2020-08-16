package github

import "github.com/zephinzer/dev/internal/constants"

// AccountConfigs defines a list of accounts that can be operated on
type AccountConfigs []AccountConfig

// GetSanitized returns a clone of the current Accounts instance
// without sensitive credentials
func (a AccountConfigs) GetSanitized() AccountConfigs {
	accounts := []AccountConfig{}
	for _, account := range a {
		if account.Public {
			accounts = append(accounts, account.GetSanitized())
		}
	}
	return accounts
}

// AccountConfig defines the structure of the configuration for a single
// logical Github account in the configuration
type AccountConfig struct {
	// Name is the user-defined label for this github account
	Name string `json:"name" yaml:"name,omitempty"`
	// Description is a user-defined description of what this account is for
	Description string `json:"description" yaml:"description,omitempty"`
	// AccessToken is the token that can be generated for use as a Personal Access Token,
	// this can be created at https://github.com/settings/tokens
	//
	// You'll need the following list of permissions when generating this:
	// - repo:status
	// - repo_deployment
	// - public_repo
	// - repo:invite
	// - read:packages
	// - read:org
	// - read:public_key
	// - read:repo_hook
	// - notifications
	// - read:user
	// - read:discussion
	// - read:enterprise
	// - read:gpg_key
	AccessToken string `json:"accessToken" yaml:"accessToken"`
	// Public indicates whether this account should be public, if so, the /platforms
	// endpoint on the dev server will expose this account; this is done to accomodate
	// using both personal and work accounts
	Public bool `json:"public" yaml:"public,omitempty"`
}

// GetSanitized returns a copy of the current Account instance with
// sensitive credentials removed
func (a AccountConfig) GetSanitized() AccountConfig {
	return AccountConfig{
		Name:        a.Name,
		Description: a.Description,
		AccessToken: constants.DefaultRedactedString,
		Public:      a.Public,
	}
}
