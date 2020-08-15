package github

import (
	"net/http"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/types"
	pkg "github.com/zephinzer/dev/pkg/github"
)

func GetAccount(accessToken string) (types.Account, error) {
	client := &http.Client{Timeout: constants.DefaultAPICallTimeout}
	account, getAccountError := pkg.GetAccount(client, accessToken)
	if getAccountError != nil {
		return nil, getAccountError
	}
	return Account(*account), nil
}

// Account implements types.Account
type Account pkg.APIv3UserResponse

func (a Account) GetName() *string     { return &a.Name }
func (a Account) GetUsername() *string { return &a.Login }
func (a Account) GetEmail() *string {
	if email, ok := a.Email.(string); ok {
		return &email
	}
	return nil
}
func (a Account) GetLink() *string       { return &a.HTMLURL }
func (a Account) GetCreatedAt() *string  { return &a.CreatedAt }
func (a Account) GetLastSeen() *string   { return &a.UpdatedAt }
func (a Account) GetFollowerCount() *int { return &a.Followers }
func (a Account) GetProjectCount() *int {
	totalRepos := a.PublicRepos + a.TotalPrivateRepos
	return &totalRepos
}
func (a Account) GetIsAdmin() *bool { return nil }
