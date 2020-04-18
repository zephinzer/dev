package github

type Config struct {
	Accounts Accounts `yaml:"accounts"`
}

type Accounts []Account

type Account struct {
	// Name is the user-defined label for this github account
	Name string `yaml:"name"`
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
	AccessToken string `yaml:"accessToken"`
}
