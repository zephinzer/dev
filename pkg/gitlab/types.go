package gitlab

type Config struct {
	Accounts Accounts `yaml:"accounts"`
}

type Accounts []Account

type Account struct {
	// Name is the user-defined label for this account
	Name string `yaml:"name"`
	// Hostname is the hostname of the Gitlab installation. If this is not provided
	// this should default to "gitlab.com"
	Hostname string `yaml:"hostname"`
	// AccessToken is your Gitlab Personal Access Token which can be generated
	// from this page at https://gitlab.com/profile/personal_access_tokens
	//
	// The required Scopes are as follows:
	// - read_user
	// - read_api
	// - read_repository
	// - read_registry
	AccessToken string `yaml:"accessToken"`
}
