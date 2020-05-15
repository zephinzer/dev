package validator

import (
	"log"
	"net/url"
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
	log.Printf("RegexForGitHTTP: %s", RegexForGitHTTP)
	log.Printf("RegexForGitSSH: %s", RegexForGitSSH)
}

func (s *ValidatorTests) TestIsGitHTTPUrl() {
	regex := regexp.MustCompile(RegexForGitHTTP)
	givenURL := "http://gitlab.com/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"http", "", "", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com:12345/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username@gitlab.com:12345/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "username", "", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username:password@gitlab.com:12345/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "username", "password", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
}

func (s *ValidatorTests) TestIsGitSSHURL() {
	regex := regexp.MustCompile(RegexForGitSSH)
	givenURL := "git@gitlab.com:zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:22/zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "22", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:22a/zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "22a", "zephinzer/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:a22/zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "a22", "zephinzer/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
}

func (s *ValidatorTests) TestParseURL_GitHTTP() {
	given := "https://gitlab.com/zephinzer/dev.git"
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

func (s *ValidatorTests) TestParseURL_Custom() {
	given := "someapp://somedata/somepath"
	expected := &URL{
		Schema:   "someapp",
		Hostname: "somedata",
		Path:     "/somepath",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP() {
	given := "http://gitlab.com/zephinzer/dev"
	expected := &URL{
		Schema:   "http",
		Hostname: "gitlab.com",
		Path:     "/zephinzer/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP_withPort() {
	given := "http://gitlab.com:123456/zephinzer/dev"
	expected := &URL{
		Schema:   "http",
		Hostname: "gitlab.com",
		Port:     "123456",
		Path:     "/zephinzer/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP_withQueryParameters() {
	given := "http://gitlab.com/zephinzer/dev?a=1&b=hello world, love, dev"
	expected := &URL{
		Schema:   "http",
		Hostname: "gitlab.com",
		Path:     "/zephinzer/dev",
		Query:    "a=1&b=hello+world%2C+love%2C+dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	parsedQuery, parseQueryError := url.ParseQuery(observed.Query)
	s.Nil(parseQueryError)
	s.Equal("1", parsedQuery.Get("a"))
	s.Equal("hello world, love, dev", parsedQuery.Get("b"))
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP_withUsernameCredentials() {
	given := "http://username@gitlab.com/zephinzer/dev"
	expected := &URL{
		Schema:   "http",
		Hostname: "gitlab.com",
		Path:     "/zephinzer/dev",
		Username: "username",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_HTTP_withFullCredentials() {
	given := "http://username:password@gitlab.com/zephinzer/dev"
	expected := &URL{
		Schema:   "http",
		Hostname: "gitlab.com",
		Path:     "/zephinzer/dev",
		Username: "username",
		Password: "password",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
}

func (s *ValidatorTests) TestParseURL_GitHTTP_insecure() {
	given := "http://gitlab.com/zephinzer/dev.git"
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

func (s *ValidatorTests) TestParseURL_GitHTTP_withCredentials() {
	given := "https://username:password@gitlab.com/zephinzer/dev.git"
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

func (s *ValidatorTests) TestParseURL_GitHTTP_withPorts() {
	given := "https://gitlab.com:443/zephinzer/dev.git"
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

func (s *ValidatorTests) TestParseURL_GitSSH() {
	given := "git@gitlab.com:zephinzer/dev.git"
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

func (s *ValidatorTests) TestParseURL_GitSSH_withPort() {
	given := "git@gitlab.com:8888/zephinzer/dev.git"
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
