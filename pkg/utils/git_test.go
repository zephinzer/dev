package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GitTests struct {
	suite.Suite
}

func TestGit(t *testing.T) {
	suite.Run(t, &GitTests{})
}

func (s *GitTests) Test_GetSshCloneUrlFromHttpLinkUrl() {
	observed, err := GetSshCloneUrlFromHttpLinkUrl("https://github.com/usvc/go-db")
	s.Nil(err)
	s.Equal("git@github.com:usvc/go-db.git", observed)
}

func (s *GitTests) Test_GetHttpCloneUrlFromHttpLink() {
	observed, err := GetHttpCloneUrlFromHttpLink("https://github.com/usvc/go-db")
	s.Nil(err)
	s.Equal("https://github.com/usvc/go-db.git", observed)
}

func (s *GitTests) Test_GetHttpLinkFromHttpCloneUrl() {
	observed, err := GetHttpLinkFromHttpCloneUrl("https://github.com/usvc/go-db.git")
	s.Nil(err)
	s.Equal("https://github.com/usvc/go-db", observed)
}

func (s *GitTests) Test_GetHttpLinkFromSshCloneUrl() {
	observed, err := GetHttpLinkFromSshCloneUrl("git@github.com:usvc/go-db.git")
	s.Nil(err)
	s.Equal("https://github.com/usvc/go-db", observed)
}
