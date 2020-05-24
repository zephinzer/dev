package pivotaltracker

import (
	"github.com/zephinzer/dev/internal/types"
	pkg "github.com/zephinzer/dev/pkg/pivotaltracker"
)

func GetAccount(accessToken string) (types.Account, error) {
	account, getAccountError := pkg.GetAccount(accessToken)
	if getAccountError != nil {
		return nil, getAccountError
	}
	return Account(*account), nil
}

// Account implements types.Account
type Account pkg.APIv5AccountResponse

func (a Account) GetName() *string       { return &a.Name }
func (a Account) GetUsername() *string   { return &a.Username }
func (a Account) GetEmail() *string      { return &a.Email }
func (a Account) GetLink() *string       { return nil }
func (a Account) GetCreatedAt() *string  { return &a.CreatedAt }
func (a Account) GetLastSeen() *string   { return &a.UpdatedAt }
func (a Account) GetFollowerCount() *int { return nil }
func (a Account) GetProjectCount() *int {
	projectsCount := len(a.Projects)
	return &projectsCount
}
func (a Account) GetIsAdmin() *bool { return nil }
