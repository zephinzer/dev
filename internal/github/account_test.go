package github

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

func (s *AccountTests) Test_Account_fitsInterface() {
	var acc types.Account
	s.NotPanics(func() {
		acc = Account{}
	})
	s.NotNil(acc)
}

func (s AccountTests) Test_Account_GetAccount() {
	systemError := tests.CaptureRequestWithTLS(
		func(client request.Doer) error {
			_, err := GetAccount(client, "__access_token")
			return err
		},
		func(req *http.Request) error {
			s.Equal("api.github.com", req.Host)
			s.EqualValues("application/vnd.github.v3+json", req.Header["Accept"][0])
			s.EqualValues("token __access_token", req.Header["Authorization"][0])
			return nil
		},
		[]byte("{}"),
	)
	s.Nil(systemError)
}

func (s AccountTests) Test_Account_Getters() {
	acc := Account{
		Name:              "__name",
		Login:             "__username",
		Email:             "__email",
		HTMLURL:           "__link",
		CreatedAt:         "__created_at",
		UpdatedAt:         "__updated_at",
		Followers:         12345,
		PublicRepos:       1,
		TotalPrivateRepos: 1,
	}
	s.Equal(acc.Name, *acc.GetName())
	s.Equal(acc.Login, *acc.GetUsername())
	s.Equal(acc.Email, *acc.GetEmail())
	s.Equal(acc.HTMLURL, *acc.GetLink())
	s.Equal(acc.CreatedAt, *acc.GetCreatedAt())
	s.Equal(acc.UpdatedAt, *acc.GetLastSeen())
	s.Equal(acc.Followers, *acc.GetFollowerCount())
	s.Equal(acc.PublicRepos+acc.TotalPrivateRepos, *acc.GetProjectCount())
	s.Nil(acc.GetIsAdmin())
}
