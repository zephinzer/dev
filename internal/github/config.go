package github

// Config defines the structure of the configuration
// for the github platform
type Config struct {
	Accounts Accounts `json:"accounts" yaml:"accounts"`
}

// GetSanitized returns a clone of the current Config instance
// without sensitive credentials
func (c Config) GetSanitized() Config {
	return Config{
		Accounts: c.Accounts.GetSanitized(),
	}
}

// Accounts defines a list of accounts that can be operated on
type Accounts []Account

// GetSanitized returns a clone of the current Accounts instance
// without sensitive credentials
func (a Accounts) GetSanitized() Accounts {
	accounts := []Account{}
	for _, account := range a {
		if account.Public {
			accounts = append(accounts, account.GetSanitized())
		}
	}
	return accounts
}

// Account defines the structure of the configuration for a single
// logical Github account in the configuration
type Account struct {
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
func (a Account) GetSanitized() Account {
	return Account{
		Name:        a.Name,
		Description: a.Description,
		AccessToken: "[REDACTED]",
		Public:      a.Public,
	}
}
