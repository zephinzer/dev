package gitlab

import (
	"testing"

	"github.com/stretchr/testify/suite"
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
