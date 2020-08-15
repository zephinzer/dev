package gitlab

import (
	"net/http"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/internal/types"
	pkg "github.com/zephinzer/dev/pkg/gitlab"
)

func GetAccount(hostname, accessToken string) (types.Account, error) {
	client := &http.Client{Timeout: constants.DefaultAPICallTimeout}
	account, getAccountError := pkg.GetAccount(client, hostname, accessToken)
	if getAccountError != nil {
		return nil, getAccountError
	}
	return Account(*account), nil
}

// Account implements types.Account
type Account pkg.APIv4UserResponse

func (a Account) GetName() *string       { return &a.Name }
func (a Account) GetUsername() *string   { return &a.Username }
func (a Account) GetEmail() *string      { return &a.Email }
func (a Account) GetLink() *string       { return &a.WebURL }
func (a Account) GetCreatedAt() *string  { return &a.CreatedAt }
func (a Account) GetLastSeen() *string   { return &a.LastActivityOn }
func (a Account) GetFollowerCount() *int { return nil }
func (a Account) GetProjectCount() *int  { return nil }
func (a Account) GetIsAdmin() *bool      { return &a.IsAdmin }
