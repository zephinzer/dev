package config

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
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
			URL:        "git@github.com:zephinzer/dev.git",
			Name:       "repo 1",
			Workspaces: []string{"ab", "cd"},
		},
		{
			URL:        "git@github.com:zephinzer/dev.git",
			Name:       "repo 2",
			Workspaces: []string{"abc", "def"},
		},
		{
			URL:        "git@github.com:zephinzer/dev.git",
			Name:       "repo 3",
			Workspaces: []string{"ab", "def"},
		},
	}
	s.Contains(repos.GetWorkspaces(), "ab")
	s.Contains(repos.GetWorkspaces(), "cd")
	s.Contains(repos.GetWorkspaces(), "abc")
	s.Contains(repos.GetWorkspaces(), "def")
}

func (s RepositoriesTests) TestMergeWith() {
	reposA := Repositories{
		{
			Name: "repo1",
			URL:  "git@github.com:user1/repo1.git",
		},
		{
			Name: "repo2",
			URL:  "git@github.com:user1/repo2.git",
		},
	}
	reposB := Repositories{
		{
			Name: "repo1",
			URL:  "git@github.com:user2/repo1.git",
		},
		{
			Name: "repo2",
			URL:  "git@github.com:user2/repo2.git",
		},
	}
	s.Len(reposA, 2)
	reposA.MergeWith(reposB)
	s.Len(reposA, 4)
}

func (s RepositoriesTests) TestYAMLUnmarshal() {
	repositoryYAML, readFileError := ioutil.ReadFile("../../tests/config/repositories.yaml")
	s.Nil(readFileError)
	if readFileError != nil {
		return
	}
	c := Config{}
	unmarshalError := yaml.Unmarshal(repositoryYAML, &c)
	s.Nil(unmarshalError)
	if unmarshalError != nil {
		return
	}
	s.Len(c.Repositories, 4)
}
