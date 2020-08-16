package github

// Config defines the structure of the configuration
// for the github platform
type Config struct {
	Accounts AccountConfigs `json:"accounts" yaml:"accounts,omitempty"`
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
