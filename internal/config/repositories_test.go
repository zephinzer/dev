package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RepositoriesTests struct {
	suite.Suite
}

func TestRepositories(t *testing.T) {
	suite.Run(t, &RepositoriesTests{})
}

func (s RepositoriesTests) TestGetWorkspaces() {
	repos := Repositories{
		{
			CloneURL:   "git@github.com:zephinzer/dev.git",
			Name:       "repo 1",
			Workspaces: []string{"ab", "cd"},
		},
		{
			CloneURL:   "git@github.com:zephinzer/dev.git",
			Name:       "repo 2",
			Workspaces: []string{"abc", "def"},
		},
		{
			CloneURL:   "git@github.com:zephinzer/dev.git",
			Name:       "repo 3",
			Workspaces: []string{"ab", "def"},
		},
	}
	s.EqualValues([]string{"ab", "cd", "abc", "def"}, repos.GetWorkspaces())
}

func (s RepositoriesTests) TestMergeWith() {
	reposA := Repositories{
		{
			Name:     "repo1",
			CloneURL: "git@github.com:user1/repo1.git",
		},
		{
			Name:     "repo2",
			CloneURL: "git@github.com:user1/repo2.git",
		},
	}
	reposB := Repositories{
		{
			Name:     "repo1",
			CloneURL: "git@github.com:user2/repo1.git",
		},
		{
			Name:     "repo2",
			CloneURL: "git@github.com:user2/repo2.git",
		},
	}
	s.Len(reposA, 2)
	reposA.MergeWith(reposB)
	s.Len(reposA, 4)
}
