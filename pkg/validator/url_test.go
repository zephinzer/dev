package validator

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
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
		[]string{"http", "", "", "gitlab.com", "", "zephinzer", "/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "", "zephinzer", "/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com:12345/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "12345", "zephinzer", "/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username@gitlab.com:12345/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "username", "", "gitlab.com", "12345", "zephinzer", "/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username:password@gitlab.com:12345/zephinzer/dev.git"
	s.True(IsGitHTTPUrl(givenURL))
	s.Equal(
		[]string{"https", "username", "password", "gitlab.com", "12345", "zephinzer", "/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
}

func (s *ValidatorTests) TestIsGitSSHURL() {
	regex := regexp.MustCompile(RegexForGitSSH)
	givenURL := "git@gitlab.com:zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "zephinzer", "/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:22/zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "22", "zephinzer", "/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:22a/zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "22a", "/zephinzer/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "git@gitlab.com:a22/zephinzer/dev.git"
	s.True(IsGitSSHUrl(givenURL))
	s.Equal(
		[]string{"git", "gitlab.com", "", "a22", "/zephinzer/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
}

func (s *ValidatorTests) TestParseURL_GitHTTP() {
	expectedString := "https://gitlab.com/zephinzer/dev"
	given := expectedString + ".git"
	expected := &URL{
		Schema:   "https",
		Hostname: "gitlab.com",
		User:     "zephinzer",
		Path:     "/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
	s.EqualValues(expectedString, observed.String())
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
	s.EqualValues(given, observed.String())
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
	s.EqualValues(given, observed.String())
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
	s.EqualValues(given, observed.String())
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
	s.EqualValues("http://gitlab.com/zephinzer/dev?a=1&b=hello+world%2C+love%2C+dev", observed.String())
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
	s.EqualValues(given, observed.String())
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
	s.EqualValues(given, observed.String())
}

func (s *ValidatorTests) TestParseURL_GitHTTP_insecure() {
	expectedString := "http://gitlab.com/zephinzer/dev"
	given := expectedString + ".git"
	expected := &URL{
		Schema:   "http",
		Hostname: "gitlab.com",
		User:     "zephinzer",
		Path:     "/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
	s.EqualValues(expectedString, observed.String())
}

func (s *ValidatorTests) TestParseURL_GitHTTP_withCredentials() {
	expectedString := "https://username:password@gitlab.com/zephinzer/dev"
	given := expectedString + ".git"
	expected := &URL{
		Schema:   "https",
		Username: "username",
		Password: "password",
		Hostname: "gitlab.com",
		User:     "zephinzer",
		Path:     "/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
	s.EqualValues(expectedString, observed.String())
}

func (s *ValidatorTests) TestParseURL_GitHTTP_withPorts() {
	expectedString := "https://gitlab.com:443/zephinzer/dev"
	given := expectedString + ".git"
	expected := &URL{
		Schema:   "https",
		Hostname: "gitlab.com",
		Port:     "443",
		User:     "zephinzer",
		Path:     "/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
	s.EqualValues(expectedString, observed.String())
}

func (s *ValidatorTests) TestParseURL_GitSSH() {
	expectedString := "git@gitlab.com:zephinzer/dev"
	given := expectedString + ".git"
	expected := &URL{
		Username: "git",
		Hostname: "gitlab.com",
		User:     "zephinzer",
		Path:     "/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
	s.EqualValues(expectedString, observed.String())
}

func (s *ValidatorTests) TestParseURL_GitSSH_withPort() {
	expectedString := "git@gitlab.com:8888/zephinzer/dev"
	given := expectedString + ".git"
	expected := &URL{
		Username: "git",
		Hostname: "gitlab.com",
		Port:     "8888",
		User:     "zephinzer",
		Path:     "/dev",
	}
	observed, err := ParseURL(given)
	s.Nil(err)
	s.EqualValues(expected, observed)
	s.EqualValues(expectedString, observed.String())
}

func (s *ValidatorTests) Test_parseRegexIntoFields() {
	groups := []string{"group1", "group2"}
	var regexString strings.Builder
	for _, group := range groups {
		regexString.WriteString(fmt.Sprintf("(?P<%s>%s)", group, group))
	}
	groupsString := strings.Join(groups, "")
	observed := parseRegexIntoFields(regexp.MustCompile(regexString.String()), groupsString)
	observedLength := 0
	for _, group := range groups {
		s.Equal(group, observed[group])
		observedLength++
	}
	s.Equal(len(groups), observedLength)
}

func (s *ValidatorTests) Test_serializeFieldsToURL() {
	testSchema := map[string]string{
		keySchema:   keySchema,
		keyUsername: keyUsername,
		keyPassword: keyPassword,
		keyHostname: keyHostname,
		keyPort:     keyPort,
		keyUser:     keyUser,
		keyPath:     keyPath,
	}
	observedURL := serializeFieldsToURL(testSchema)
	s.Equal(keySchema, observedURL.Schema)
	s.Equal(keyUsername, observedURL.Username)
	s.Equal(keyPassword, observedURL.Password)
	s.Equal(keyHostname, observedURL.Hostname)
	s.Equal(keyPort, observedURL.Port)
	s.Equal(keyUser, observedURL.User)
	s.Equal(keyPath, observedURL.Path)
}
