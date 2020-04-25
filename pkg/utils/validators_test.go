package utils

import (
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
}

func (s *ValidatorTests) TestRegexForHTTP() {
	regex := regexp.MustCompile(RegexForHTTP)
	givenURL := "http://gitlab.com/usvc/utils/dev.git"
	s.True(regex.MatchString(givenURL))
	s.Equal(
		[]string{"http", "", "", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com/usvc/utils/dev.git"
	s.True(regex.MatchString(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://gitlab.com:12345/usvc/utils/dev.git"
	s.True(regex.MatchString(givenURL))
	s.Equal(
		[]string{"https", "", "", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username@gitlab.com:12345/usvc/utils/dev.git"
	s.True(regex.MatchString(givenURL))
	s.Equal(
		[]string{"https", "username", "", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
	givenURL = "https://username:password@gitlab.com:12345/usvc/utils/dev.git"
	s.True(regex.MatchString(givenURL))
	s.Equal(
		[]string{"https", "username", "password", "gitlab.com", "12345", "usvc", "utils/dev"},
		regex.FindStringSubmatch(givenURL)[1:],
	)
}

func (s *ValidatorTests) TestRegexForSSH() {
	regex := regexp.MustCompile(RegexForSSH)
	s.True(regex.MatchString("git@gitlab.com:usvc/utils/dev.git"))
	s.Equal(
		[]string{"git", "gitlab.com", "", "usvc", "utils/dev"},
		regex.FindStringSubmatch("git@gitlab.com:usvc/utils/dev.git")[1:],
	)
	s.True(regex.MatchString("git@gitlab.com:22/usvc/utils/dev.git"))
	s.Equal(
		[]string{"git", "gitlab.com", "22", "usvc", "utils/dev"},
		regex.FindStringSubmatch("git@gitlab.com:22/usvc/utils/dev.git")[1:],
	)
	s.True(regex.MatchString("git@gitlab.com:22a/usvc/utils/dev.git"))
	s.Equal(
		[]string{"git", "gitlab.com", "", "22a", "usvc/utils/dev"},
		regex.FindStringSubmatch("git@gitlab.com:22a/usvc/utils/dev.git")[1:],
	)
	s.True(regex.MatchString("git@gitlab.com:a22/usvc/utils/dev.git"))
	s.Equal(
		[]string{"git", "gitlab.com", "", "a22", "usvc/utils/dev"},
		regex.FindStringSubmatch("git@gitlab.com:a22/usvc/utils/dev.git")[1:],
	)
}
