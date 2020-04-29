package github

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/usvc/dev/internal/constants"
)

type Config struct {
	Accounts Accounts `json:"accounts" yaml:"accounts"`
}

func (c Config) GetSanitized() Config {
	return Config{
		Accounts: c.Accounts.GetSanitized(),
	}
}

type Accounts []Account

func (a Accounts) GetSanitized() Accounts {
	accounts := []Account{}
	for _, account := range a {
		if account.Public {
			accounts = append(accounts, account.GetSanitized())
		}
	}
	return accounts
}

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

func (a Account) GetSanitized() Account {
	return Account{
		Name:        a.Name,
		Description: a.Description,
		AccessToken: "[REDACTED]",
		Public:      a.Public,
	}
}

// APIv3EmailsResponse defines the response structure from the
// https://api.github.com/user/public_emails endpoint
type APIv3EmailsResponse []APIv3Email

// APIv3UserResponse defines the response structure from the
// https://api.github.com/user endpoint
type APIv3UserResponse struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeID                  string      `json:"node_id"`
	AvatarURL               string      `json:"avatar_url"`
	GravatarID              string      `json:"gravatar_id"`
	URL                     string      `json:"url"`
	HTMLURL                 string      `json:"html_url"`
	FollowersURL            string      `json:"followers_url"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	OrganizationsURL        string      `json:"organizations_url"`
	ReposURL                string      `json:"repos_url"`
	EventsURL               string      `json:"events_url"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 string      `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   interface{} `json:"email"`
	Hireable                bool        `json:"hireable"`
	Bio                     string      `json:"bio"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               string      `json:"created_at"`
	UpdatedAt               string      `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    APIv3Plan   `json:"plan"`
}

func (u APIv3UserResponse) String() string {
	var output strings.Builder
	output.WriteString("github account information\n")
	output.WriteString(fmt.Sprintf("username      : %s\n", u.Login))
	output.WriteString(fmt.Sprintf("real name     : %s\n", u.Name))
	if u.Email != nil {
		output.WriteString(fmt.Sprintf("account email : %s\n", u.Email))
	} else {
		output.WriteString("account email : (hidden)\n")
	}
	output.WriteString(fmt.Sprintf("2fa-enabled   : %v\n", u.TwoFactorAuthentication))
	output.WriteString(fmt.Sprintf("link          : %s\n", u.HTMLURL))
	if createdAt, err := time.Parse(constants.GithubAPITimeFormat, u.CreatedAt); err != nil {
		output.WriteString(fmt.Sprintf("created at    : %s\n", u.CreatedAt))
	} else {
		output.WriteString(fmt.Sprintf("created at    : %s (about %s)\n", u.CreatedAt, humanize.Time(createdAt)))
	}
	if lastUpdated, err := time.Parse(constants.GithubAPITimeFormat, u.UpdatedAt); err != nil {
		output.WriteString(fmt.Sprintf("last updated  : %s\n", u.UpdatedAt))
	} else {
		output.WriteString(fmt.Sprintf("last updated  : %s (about %s)\n", u.UpdatedAt, humanize.Time(lastUpdated)))
	}
	output.WriteString("vanities      :\n")
	output.WriteString(fmt.Sprintf("  followers   : %v\n", u.Followers))
	output.WriteString(fmt.Sprintf("  public repos: %v\n", u.PublicRepos))
	output.WriteString(fmt.Sprintf("  total repos : %v\n", u.PublicRepos+u.TotalPrivateRepos))
	output.WriteString(fmt.Sprintf("  public gists: %v\n", u.PublicGists))
	output.WriteString(fmt.Sprintf("  total gists : %v\n", u.PublicGists+u.PrivateGists))
	return output.String()
}

type APIv3Email struct {
	Email      string `json:"email"`
	Verified   bool   `json:"verified"`
	Primary    bool   `json:"primary"`
	Visibility string `json:"visibility"`
}

type APIv3Plan struct {
	Name          string `json:"name"`
	Space         int    `json:"space"`
	Collaborators int    `json:"collaborators"`
	PrivateRepos  int    `json:"private_repos"`
}
