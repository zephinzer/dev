package pivotaltracker

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/internal/types"
	pkg "github.com/zephinzer/dev/pkg/pivotaltracker"
)

type AccountTests struct {
	suite.Suite
}

func TestAccount(t *testing.T) {
	suite.Run(t, &AccountTests{})
}

func (s AccountTests) Test_Account_fitsInterface() {
	var acc types.Account
	s.NotPanics(func() {
		acc = Account{}
	})
	s.NotNil(acc)
}

func (s AccountTests) Test_Account_getters() {
	account := Account{
		Name:      "__name",
		Username:  "__username",
		Email:     "__email",
		CreatedAt: "__created_at",
		UpdatedAt: "__updated_at",
		Projects:  []pkg.APIProject{{}, {}},
	}
	s.Equal("__name", *account.GetName())
	s.Equal("__username", *account.GetUsername())
	s.Equal("__email", *account.GetEmail())
	s.Equal("__created_at", *account.GetCreatedAt())
	s.Equal("__updated_at", *account.GetLastSeen())
	s.Equal(2, *account.GetProjectCount())
	s.Nil(account.GetLink())
	s.Nil(account.GetIsAdmin())
	s.Nil(account.GetFollowerCount())
}
