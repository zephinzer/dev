package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GitRemoteTests struct {
	suite.Suite
}

func TestGitRemote(t *testing.T) {
	suite.Run(t, &GitRemoteTests{})
}

func (s *GitRemoteTests) Test_GetSshCloneUrlFromHttpLinkUrl() {
	observed, err := GetSshCloneUrlFromHttpLinkUrl("https://github.com/usvc/go-db")
	s.Nil(err)
	s.Equal("git@github.com:usvc/go-db.git", observed)
}

func (s *GitRemoteTests) Test_GetHttpCloneUrlFromHttpLink() {
	observed, err := GetHttpCloneUrlFromHttpLink("https://github.com/usvc/go-db")
	s.Nil(err)
	s.Equal("https://github.com/usvc/go-db.git", observed)
}

func (s *GitRemoteTests) Test_GetHttpLinkFromHttpCloneUrl() {
	observed, err := GetHttpLinkFromHttpCloneUrl("https://github.com/usvc/go-db.git")
	s.Nil(err)
	s.Equal("https://github.com/usvc/go-db", observed)
}

func (s *GitRemoteTests) Test_GetHttpLinkFromSshCloneUrl() {
	observed, err := GetHttpLinkFromSshCloneUrl("git@github.com:usvc/go-db.git")
	s.Nil(err)
	s.Equal("https://github.com/usvc/go-db", observed)
}
