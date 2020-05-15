package gitlab

import (
	"github.com/zephinzer/dev/internal/types"
	gl "github.com/zephinzer/dev/pkg/gitlab"
)

func GetAccount(hostname, accessToken string) (types.Account, error) {
	accountInfo, getAccountError := gl.GetAccount(hostname, accessToken)
	if getAccountError != nil {
		return nil, getAccountError
	}
	return AccountSerializer{accountInfo}, nil
}

type AccountSerializer struct {
	instance *gl.APIv4UserResponse
}

func (as AccountSerializer) GetName() *string       { return &as.instance.Name }
func (as AccountSerializer) GetUsername() *string   { return &as.instance.Username }
func (as AccountSerializer) GetEmail() *string      { return &as.instance.Email }
func (as AccountSerializer) GetLink() *string       { return &as.instance.WebURL }
func (as AccountSerializer) GetCreatedAt() *string  { return &as.instance.CreatedAt }
func (as AccountSerializer) GetLastSeen() *string   { return &as.instance.LastActivityOn }
func (as AccountSerializer) GetFollowerCount() *int { return nil }
func (as AccountSerializer) GetProjectCount() *int  { return nil }
func (as AccountSerializer) GetIsAdmin() *bool      { return &as.instance.IsAdmin }
