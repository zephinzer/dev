package github

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/internal/constants"
)

type AccountConfigTests struct {
	suite.Suite
}

func TestAccountConfig(t *testing.T) {
	suite.Run(t, &AccountConfigTests{})
}

func (s AccountConfigTests) Test_AccountConfigs_GetSanitized() {
	var accs AccountConfigs = AccountConfigs{
		{AccessToken: "__access_token_0"},
		{AccessToken: "__access_token_1"},
		{AccessToken: "__access_token_2"},
	}
	sanitizeds := accs.GetSanitized()
	for _, sanitized := range sanitizeds {
		s.NotContains(sanitized.AccessToken, "__access_token", ".AccessToken field should not be what it was")
		s.Equal(sanitized.AccessToken, constants.DefaultRedactedString, ".AccessToken field should have been masked with the default redacted string")
	}
}

func (s AccountConfigTests) Test_AccountConfig_GetSanitized() {
	accconf := AccountConfig{
		Name:        "__name",
		Description: "__description",
		AccessToken: "__access_token",
		Public:      true,
	}
	sanitized := accconf.GetSanitized()
	s.Equal("__name", sanitized.Name, ".Name field should not be changed")
	s.Equal("__description", sanitized.Description, ".Description field should not be changed")
	s.NotEqual("__access_token", sanitized.AccessToken, ".AccessToken field should have been masked")
	s.Equal(constants.DefaultRedactedString, sanitized.AccessToken, ".AccessToken field should have been masked with the default redacted string")
	s.True(sanitized.Public, ".Public field should not be changed")

}
