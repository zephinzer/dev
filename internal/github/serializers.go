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
	return AccountSerializer{accountInfo}, nil
}

type AccountSerializer struct {
	instance *gh.APIv3UserResponse
}

func (as AccountSerializer) GetName() *string     { return &as.instance.Name }
func (as AccountSerializer) GetUsername() *string { return &as.instance.Login }
func (as AccountSerializer) GetEmail() *string {
	if email, ok := as.instance.Email.(string); ok {
		return &email
	}
	return nil
}
func (as AccountSerializer) GetLink() *string       { return &as.instance.HTMLURL }
func (as AccountSerializer) GetCreatedAt() *string  { return &as.instance.CreatedAt }
func (as AccountSerializer) GetLastSeen() *string   { return &as.instance.UpdatedAt }
func (as AccountSerializer) GetFollowerCount() *int { return &as.instance.Followers }
func (as AccountSerializer) GetProjectCount() *int {
	totalRepos := as.instance.PublicRepos + as.instance.TotalPrivateRepos
	return &totalRepos
}
func (as AccountSerializer) GetIsAdmin() *bool { return nil }
