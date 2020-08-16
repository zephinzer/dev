package github

import (
	"github.com/zephinzer/dev/internal/types"
	pkg "github.com/zephinzer/dev/pkg/github"
	"github.com/zephinzer/dev/pkg/utils/request"
)

// GetAccount is a utility function that calls the main Github API and wraps the response
// so that it can be used with an Account interface
func GetAccount(client request.Doer, accessToken string) (types.Account, error) {
	account, getAccountError := pkg.GetAccount(client, accessToken)
	if getAccountError != nil {
		return nil, getAccountError
	}
	return Account(*account), nil
}

// Account implements types.Account
type Account pkg.APIv3UserResponse

// GetName implements the internal/types.Account interface
func (a Account) GetName() *string { return &a.Name }

// GetUsername implements the internal/types.Account interface
func (a Account) GetUsername() *string { return &a.Login }

// GetEmail implements the internal/types.Account interface
func (a Account) GetEmail() *string {
	if email, ok := a.Email.(string); ok {
		return &email
	}
	return nil
}

// GetLink implements the internal/types.Account interface
func (a Account) GetLink() *string { return &a.HTMLURL }

// GetCreatedAt implements the internal/types.Account interface
func (a Account) GetCreatedAt() *string { return &a.CreatedAt }

// GetLastSeen implements the internal/types.Account interface
func (a Account) GetLastSeen() *string { return &a.UpdatedAt }

// GetFollowerCount implements the internal/types.Account interface
func (a Account) GetFollowerCount() *int { return &a.Followers }

// GetProjectCount implements the internal/types.Account interface
func (a Account) GetProjectCount() *int {
	totalRepos := a.PublicRepos + a.TotalPrivateRepos
	return &totalRepos
}

// GetIsAdmin implements the internal/types.Account interface
func (a Account) GetIsAdmin() *bool { return nil }
