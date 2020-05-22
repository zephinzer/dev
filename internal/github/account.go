package github

import (
	"github.com/zephinzer/dev/internal/types"
	gh "github.com/zephinzer/dev/pkg/github"
)

func GetAccount(accessToken string) (types.Account, error) {
	accountInfo, getAccountError := gh.GetAccount(accessToken)
	if getAccountError != nil {
		return nil, getAccountError
	}
	return Account{accountInfo}, nil
}

type Account struct {
	instance *gh.APIv3UserResponse
}

func (as Account) GetName() *string     { return &as.instance.Name }
func (as Account) GetUsername() *string { return &as.instance.Login }
func (as Account) GetEmail() *string {
	if email, ok := as.instance.Email.(string); ok {
		return &email
	}
	return nil
}
func (as Account) GetLink() *string       { return &as.instance.HTMLURL }
func (as Account) GetCreatedAt() *string  { return &as.instance.CreatedAt }
func (as Account) GetLastSeen() *string   { return &as.instance.UpdatedAt }
func (as Account) GetFollowerCount() *int { return &as.instance.Followers }
func (as Account) GetProjectCount() *int {
	totalRepos := as.instance.PublicRepos + as.instance.TotalPrivateRepos
	return &totalRepos
}
func (as Account) GetIsAdmin() *bool { return nil }
