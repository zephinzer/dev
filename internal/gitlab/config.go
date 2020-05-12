package gitlab

// Config defines the structure of the configuration for the
// Gitlab integration
type Config struct {
	Accounts Accounts `json:"accounts" yaml:"accounts"`
}

// GetSanitized returns a copy of this Config instance without
// sensitive credentials
func (c Config) GetSanitized() Config {
	return Config{
		Accounts: c.Accounts.GetSanitized(),
	}
}

// Accounts defines a list of Account instances which can be
// operated on internally
type Accounts []Account

// GetSanitized returns a copy of this Accounts instance without
// sensitive credentials
func (a Accounts) GetSanitized() []Account {
	accounts := []Account{}
	for _, account := range a {
		if account.Public {
			accounts = append(accounts, account.GetSanitized())
		}
	}
	return accounts
}

// Account defines the configuration structure for a logical Gitlab
// account to be consumed
type Account struct {
	// Name is the user-defined label for this account
	Name string `json:"name" yaml:"name"`
	// Description is a user-defined description of what this account is for
	Description string `json:"description" yaml:"description"`
	// Hostname is the hostname of the Gitlab installation. If this is not provided
	// this should default to "gitlab.com"
	Hostname string `json:"hostname" yaml:"hostname"`
	// AccessToken is your Gitlab Personal Access Token which can be generated
	// from this page at https://gitlab.com/profile/personal_access_tokens
	//
	// The required Scopes are as follows:
	// - read_user
	// - read_api
	// - read_repository
	// - read_registry
	AccessToken string `json:"accessToken" yaml:"accessToken"`
	// Public indicates whether this account should be public, if so, the /platforms
	// endpoint on the dev server will expose this account; this is done to accomodate
	// using both personal and work accounts, note that even if this is set to true,
	// the AccessToken should not be exposed
	Public bool `json:"public" yaml:"public"`
}

// GetSanitized returns a copy of this Account instance without
// sensitive credentials
func (a Account) GetSanitized() Account {
	return Account{
		Name:        a.Name,
		Description: a.Description,
		AccessToken: "[REDACTED]",
		Hostname:    a.Hostname,
		Public:      a.Public,
	}
}
