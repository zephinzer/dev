package trello

import (
	"github.com/zephinzer/dev/internal/types"
	pkg "github.com/zephinzer/dev/pkg/trello"
)

func GetAccount(accessKey, accessToken string) (types.Account, error) {
	account, getAccountError := pkg.GetAccount(accessKey, accessToken)
	if getAccountError != nil {
		return nil, getAccountError
	}
	return Account(*account), nil
}

// Account implements types.Account
type Account pkg.APIv1MemberResponse

func (a Account) GetName() *string     { return &a.FullName }
func (a Account) GetUsername() *string { return &a.Username }
func (a Account) GetEmail() *string    { return &a.Email }
func (a Account) GetLink() *string {
	link := "https://trello.com/" + a.Username
	return &link
}
func (a Account) GetCreatedAt() *string  { return nil }
func (a Account) GetLastSeen() *string   { return nil }
func (a Account) GetFollowerCount() *int { return nil }
func (a Account) GetProjectCount() *int {
	numberOfBoards := len(a.BoardIDs)
	return &numberOfBoards
}
func (a Account) GetIsAdmin() *bool { return nil }
