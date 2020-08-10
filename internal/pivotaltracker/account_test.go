package pivotaltracker

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/internal/types"
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
