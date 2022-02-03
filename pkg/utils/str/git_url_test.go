package str

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GitUrlTest struct {
	suite.Suite
}

func TestGitUrl(t *testing.T) {
	suite.Run(t, &GitUrlTest{})
}

func (s GitUrlTest) Test_ParseGitUrl() {
	// perfect happy case for https
	observed, err := ParseGitUrl("https://github.com/zephinzer/dev.git")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Equal("zephinzer/dev.git", observed.Path)

	// imperfect happy case for https
	observed, err = ParseGitUrl("https://github.com/zephinzer/dev")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Equal("zephinzer/dev", observed.Path)

	// happy case for https with auth
	observed, err = ParseGitUrl("https://username:accesstoken@github.com/zephinzer/dev")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("username", *observed.Username)
	s.Equal("accesstoken", *observed.AccessToken)
	s.Equal("github.com", observed.Hostname)
	s.Equal("zephinzer/dev", observed.Path)

	// perfect happy case for ssh
	observed, err = ParseGitUrl("git@github.com:zephinzer/dev.git")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Equal("zephinzer/dev.git", observed.Path)

	// imperfect happy case for ssh
	observed, err = ParseGitUrl("git@github.com:zephinzer/dev")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Equal("zephinzer/dev", observed.Path)

	// happy case for ssh with port
	observed, err = ParseGitUrl("git@github.com:8080/zephinzer/dev")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Equal(8080, *observed.Port)
	s.Equal("zephinzer/dev", observed.Path)

	// happy case for ssh with username masquerading as a port
	observed, err = ParseGitUrl("git@github.com:65537/dev")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Nil(observed.Port)
	s.Equal("65537/dev", observed.Path)

	// weird case for ssh with username masquerading as a port
	observed, err = ParseGitUrl("git@github.com:-1/dev")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Nil(observed.Port)
	s.Equal("-1/dev", observed.Path)

	// extra happy case for ssh
	observed, err = ParseGitUrl("ssh://git@github.com:zephinzer/dev.git")
	s.Nil(err)
	s.NotNil(observed)
	s.Equal("github.com", observed.Hostname)
	s.Equal("zephinzer/dev.git", observed.Path)
}

func (s GitUrlTest) Test_GetBrowserUrl() {
	source := GitUrl{
		Hostname: "git.com",
		Path:     "a/b/c",
	}
	s.Equal("https://git.com/a/b/c", source.GetBrowserUrl())

	optionalPort := 12345
	source = GitUrl{
		Hostname: "git.com",
		Port:     &optionalPort,
		Path:     "a/b/c",
	}
	s.Equal("https://git.com:12345/a/b/c", source.GetBrowserUrl())

	authUsername := "username"
	source = GitUrl{
		Username: &authUsername,
		Hostname: "git.com",
		Path:     "a/b/c",
	}
	s.Equal("https://git.com/a/b/c", source.GetBrowserUrl())

	authPassword := "password"
	source = GitUrl{
		AccessToken: &authPassword,
		Username:    &authUsername,
		Hostname:    "git.com",
		Path:        "a/b/c",
	}
	s.Equal("https://git.com/a/b/c", source.GetBrowserUrl())
}

func (s GitUrlTest) Test_GetHttpsUrl() {
	source := GitUrl{
		Hostname: "git.com",
		Path:     "a/b/c",
	}
	s.Equal("https://git.com/a/b/c.git", source.GetHttpsUrl())

	optionalPort := 12345
	source = GitUrl{
		Hostname: "git.com",
		Port:     &optionalPort,
		Path:     "a/b/c",
	}
	s.Equal("https://git.com:12345/a/b/c.git", source.GetHttpsUrl())

	authUsername := "username"
	source = GitUrl{
		Username: &authUsername,
		Hostname: "git.com",
		Path:     "a/b/c",
	}
	s.Equal("https://username@git.com/a/b/c.git", source.GetHttpsUrl())

	authPassword := "password"
	source = GitUrl{
		AccessToken: &authPassword,
		Username:    &authUsername,
		Hostname:    "git.com",
		Path:        "a/b/c",
	}
	s.Equal("https://username:password@git.com/a/b/c.git", source.GetHttpsUrl())

	source = GitUrl{
		AccessToken: &authPassword,
		Username:    &authUsername,
		Hostname:    "git.com",
		Port:        &optionalPort,
		Path:        "a/b/c",
	}
	s.Equal("https://username:password@git.com:12345/a/b/c.git", source.GetHttpsUrl())
}

func (s GitUrlTest) Test_GetSshUrl() {
	source := GitUrl{
		Hostname: "git.com",
		Path:     "a/b/c",
	}
	s.Equal("git@git.com:a/b/c.git", source.GetSshUrl())

	optionalPort := 12345
	source = GitUrl{
		Hostname: "git.com",
		Port:     &optionalPort,
		Path:     "a/b/c",
	}
	s.Equal("git@git.com:12345/a/b/c.git", source.GetSshUrl())

	authUsername := "username"
	source = GitUrl{
		Username: &authUsername,
		Hostname: "git.com",
		Path:     "a/b/c",
	}
	s.Equal("username@git.com:a/b/c.git", source.GetSshUrl())

	authPassword := "password"
	source = GitUrl{
		AccessToken: &authPassword,
		Username:    &authUsername,
		Hostname:    "git.com",
		Path:        "a/b/c",
	}
	s.Equal("username@git.com:a/b/c.git", source.GetSshUrl())

	source = GitUrl{
		AccessToken: &authPassword,
		Username:    &authUsername,
		Hostname:    "git.com",
		Port:        &optionalPort,
		Path:        "a/b/c",
	}
	s.Equal("username@git.com:12345/a/b/c.git", source.GetSshUrl())

}
