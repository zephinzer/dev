package github

// Config defines the structure of the configuration
// for the github platform
type Config struct {
	Accounts AccountConfigs `json:"accounts" yaml:"accounts"`
}

// GetSanitized returns a clone of the current Config instance
// without sensitive credentials
func (c Config) GetSanitized() Config {
	return Config{
		Accounts: c.Accounts.GetSanitized(),
	}
}

// MergeWith merges the current Config instance with a provided
// Config instance. The merge strategy is add-only
func (c *Config) MergeWith(o Config) {
	seen := map[string]bool{}
	for _, a := range c.Accounts {
		seen[a.AccessToken] = true
	}
	for _, a := range o.Accounts {
		if value, ok := seen[a.AccessToken]; value && ok {
			continue
		}
		c.Accounts = append(c.Accounts, a)
		seen[a.AccessToken] = true
	}
}

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
	Name string `json:"name" yaml:"name"`
	// Description is a user-defined description of what this account is for
	Description string `json:"description" yaml:"description"`
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
	Public bool `json:"public" yaml:"public"`
}

// GetSanitized returns a copy of the current Account instance with
// sensitive credentials removed
func (a AccountConfig) GetSanitized() AccountConfig {
	return AccountConfig{
		Name:        a.Name,
		Description: a.Description,
		AccessToken: "[REDACTED]",
		Public:      a.Public,
	}
}
