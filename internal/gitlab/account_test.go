package gitlab

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/internal/types"
	"github.com/zephinzer/dev/pkg/utils/request"
	"github.com/zephinzer/dev/tests"
)

type AccountTests struct {
	suite.Suite
}

func TestAccount(t *testing.T) {
	suite.Run(t, &AccountTests{})
}

func (s *AccountTests) Test_Getters() {
	account := Account{
		Name:           "__name",
		Username:       "__username",
		Email:          "__email",
		WebURL:         "__web_url",
		CreatedAt:      "__created_at",
		LastActivityOn: "__last_activity_on",
		IsAdmin:        true,
	}
	s.Equal("__name", *account.GetName())
	s.Equal("__username", *account.GetUsername())
	s.Equal("__email", *account.GetEmail())
	s.Equal("__web_url", *account.GetLink())
	s.Equal("__created_at", *account.GetCreatedAt())
	s.Equal("__last_activity_on", *account.GetLastSeen())
	s.True(*account.GetIsAdmin())
	s.Nil(account.GetFollowerCount())
	s.Nil(account.GetProjectCount())
}

func (s *AccountTests) Test_Account_fitsInterface() {
	var acc types.Account
	s.NotPanics(func() {
		acc = Account{}
	})
	s.NotNil(acc)
}

func (s AccountTests) Test_Account_GetAccount() {
	s.Nil(tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetAccount(client, "__hostname", "__access_token")
			return err
		},
		func(req *http.Request) error {
			s.Equal("__hostname", req.Host)
			s.Equal([]string{"__access_token"}, req.Header["Private-Token"])
			return nil
		},
		[]byte("{}"),
	))
}
