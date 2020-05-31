package gitlab

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountConfigTests struct {
	suite.Suite
}

func TestAccountConfig(t *testing.T) {
	suite.Run(t, &AccountConfigTests{})
}

func (s *AccountConfigTests) TestAccountConfigs_GetSanitized() {
	given := AccountConfigs{
		{
			Public: false,
		},
		{
			Name:        "name",
			Description: "description",
			AccessToken: "access token",
			Hostname:    "hostname",
			Public:      true,
		},
	}
	sanitized := given.GetSanitized()
	s.Len(sanitized, 1, "when the Public property is set to false, we shouldn't include it in the sanitized version")
	expected := given[1]
	observed := sanitized[0]
	s.Equal(expected.Name, observed.Name)
	s.Equal(expected.Description, observed.Description)
	s.Equal(expected.Hostname, observed.Hostname)
	s.Equal(expected.Public, observed.Public)
	s.NotEqual(expected.AccessToken, observed.AccessToken)
}

func (s *AccountConfigTests) TestAccountConfig_GetSanitized() {
	expected := AccountConfig{
		Name:        "name",
		Description: "description",
		AccessToken: "access token",
		Hostname:    "hostname",
		Public:      true,
	}
	observed := expected.GetSanitized()
	s.Equal(expected.Name, observed.Name)
	s.Equal(expected.Description, observed.Description)
	s.Equal(expected.Hostname, observed.Hostname)
	s.Equal(expected.Public, observed.Public)
	s.NotEqual(expected.AccessToken, observed.AccessToken)
}
