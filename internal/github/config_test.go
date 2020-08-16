package github

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTests struct {
	suite.Suite
}

func TestConfig(t *testing.T) {
	suite.Run(t, &ConfigTests{})
}

func (s ConfigTests) Test_GetSanitized() {
	c := Config{
		Accounts: AccountConfigs{
			{
				AccessToken: "__access_token_0",
			},
			{
				AccessToken: "__access_token_1",
			},
		},
	}
	sanitized := c.GetSanitized()
	for _, account := range sanitized.Accounts {
		s.NotContains(account.AccessToken, "__access_token")
	}
}

func (s ConfigTests) Test_MergeWith() {
	c1 := Config{
		Accounts: AccountConfigs{
			{AccessToken: "__access_token_0"},
			{AccessToken: "__access_token_1"},
		},
	}
	c2 := Config{
		Accounts: AccountConfigs{
			{AccessToken: "__access_token_0"},
		},
	}
	c1.MergeWith(c2)
	s.Len(c1.Accounts, 2)

	c3 := Config{
		Accounts: AccountConfigs{
			{AccessToken: "__access_token_2"},
		},
	}
	c3.MergeWith(c1)
	s.Len(c3.Accounts, 3)
}
