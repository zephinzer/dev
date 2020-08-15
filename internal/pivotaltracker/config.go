package pivotaltracker

// Config defines the structure of the configuration that can be consumed for tracking
// notifications and events from the Pivotal Tracker platform
type Config struct {
	// Projects is the list of projects to monitor
	Projects Projects `json:"projects" yaml:"projects"`
	// AccessToken to be used for all listed projects if their own access token is not defined.
	// This value can be retrieved from the section named API Token at https://www.pivotaltracker.com/profile
	AccessToken string `json:"accessToken" yaml:"accessToken"`
}

// GetSanitized returns a copy of this Config instance without sensitive
// details like access tokens
func (c Config) GetSanitized() Config {
	config := Config{
		Projects: c.Projects.GetSanitized(),
	}
	// if AccessToken was never defined, don't need to include it
	if len(c.AccessToken) > 0 {
		config.AccessToken = "[REDACTED]"
	}
	return config
}

// MergeWith merges the current Config instance with a provided
// Config instance. The merge strategy is add-only
func (c *Config) MergeWith(o Config) {
	if len(c.AccessToken) == 0 && len(o.AccessToken) > 0 {
		c.AccessToken = o.AccessToken
	}
	seen := map[string]bool{}
	for _, p := range c.Projects {
		seen[p.AccessToken] = true
	}
	for _, p := range o.Projects {
		if value, ok := seen[p.AccessToken]; value && ok {
			continue
		}
		c.Projects = append(c.Projects, p)
		seen[p.AccessToken] = true
	}
}
