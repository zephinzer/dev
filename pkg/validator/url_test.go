package validator

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidatorTests struct {
	FilePath string
	suite.Suite
}

func TestValidator(t *testing.T) {
	suite.Run(t, &ValidatorTests{})
	log.Printf("RegexForHTTP: %s", RegexForHTTP)
	log.Printf("RegexForSSH: %s", RegexForSSH)
}

func (s *ValidatorTests) TestIsHTTPUrl() {
	regex := regexp.MustCompile(RegexForHTTP)
	givenURL := "http://gitlab.com/usvc/utils/dev.git"
	s.True(IsHTTPUrl(givenURL))
	s.Equal(
		[]string{"http", "", "", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com/usvc/utils/dev.git"
	s.True(IsHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com:12345/usvc/utils/dev.git"
	s.True(IsHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username@gitlab.com:12345/usvc/utils/dev.git"
	s.True(IsHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "username", "", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username:password@gitlab.com:12345/usvc/utils/dev.git"
	s.True(IsHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "username", "password", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
}

func (s *ValidatorTests) TestIsSSHURL() {
	regex := regexp.MustCompile(RegexForSSH)
	givenURL := "git@gitlab.com:usvc/utils/dev.git"
	s.True(IsSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:22/usvc/utils/dev.git"
	s.True(IsSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "22", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:22a/usvc/utils/dev.git"
	s.True(IsSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "22a", "usvc/utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:a22/usvc/utils/dev.git"
	s.True(IsSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "a22", "usvc/utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
}

func (s *ValidatorTests) TestParseURL_HTTP() {
	given := "https://gitlab.com/usvc/utils/dev.git"
	expected := &URL{
		Schema:   "https",
		Hostname: "gitlab.com",
		User:     "usvc",
		Path:     "utils/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP_insecure() {
	given := "http://gitlab.com/usvc/utils/dev.git"
	expected := &URL{
		Schema:   "http",
		Hostname: "gitlab.com",
		User:     "usvc",
		Path:     "utils/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP_withCredentials() {
	given := "https://username:password@gitlab.com/usvc/utils/dev.git"
	expected := &URL{
		Schema:   "https",
		Username: "username",
		Password: "password",
		Hostname: "gitlab.com",
		User:     "usvc",
		Path:     "utils/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP_withPorts() {
	given := "https://gitlab.com:443/usvc/utils/dev.git"
	expected := &URL{
		Schema:   "https",
		Hostname: "gitlab.com",
		Port:     "443",
		User:     "usvc",
		Path:     "utils/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_SSH() {
	given := "git@gitlab.com:usvc/utils/dev.git"
	expected := &URL{
		Username: "git",
		Hostname: "gitlab.com",
		User:     "usvc",
		Path:     "utils/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_SSH_withPort() {
	given := "git@gitlab.com:8888/usvc/utils/dev.git"
	expected := &URL{
		Username: "git",
		Hostname: "gitlab.com",
		Port:     "8888",
		User:     "usvc",
		Path:     "utils/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}
