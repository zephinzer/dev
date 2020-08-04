package gitlab

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
	baseConfig := Config{
		Accounts: AccountConfigs{
			{
				AccessToken: "a",
			},
			{
				AccessToken: "b",
			},
		},
	}
	sanitizedConfig := baseConfig.GetSanitized()
	for i := 0; i < len(sanitizedConfig.Accounts); i++ {
		s.Equal("[REDACTED]", sanitizedConfig.Accounts[i].AccessToken)
	}
}

func (s ConfigTests) Test_MergeWith() {
	baseConfig := Config{
		Accounts: AccountConfigs{
			{
				AccessToken: "a",
			},
			{
				AccessToken: "b",
			},
		},
	}
	expectedFullMerge := Config{ // should result in a full merge
		Accounts: AccountConfigs{
			{
				AccessToken: "c",
			},
			{
				AccessToken: "d",
			},
		},
	}
	expectedFullMerge.MergeWith(baseConfig)
	s.Len(expectedFullMerge.Accounts, 4)

	expectedHalfMerge := Config{ // should result in a half merge
		Accounts: AccountConfigs{
			{
				AccessToken: "a",
			},
			{
				AccessToken: "c",
			},
		},
	}
	expectedHalfMerge.MergeWith(baseConfig)
	s.Len(expectedHalfMerge.Accounts, 3)

	expectedNoMerge := Config{
		Accounts: AccountConfigs{
			{
				AccessToken: "a",
			},
			{
				AccessToken: "b",
			},
		},
	}
	expectedNoMerge.MergeWith(baseConfig)
	s.Len(expectedNoMerge.Accounts, 2)
}
