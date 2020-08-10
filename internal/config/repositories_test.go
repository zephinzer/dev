package config

import (
	"io/ioutil"
	"sort"
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

func (s RepositoriesTests) Test_GetWorkspaces() {
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

func (s RepositoriesTests) Test_MergeWith() {
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

func (s RepositoriesTests) Test_Repositories_verifySortInterface() {
	repos := Repositories{
		{
			Name: "__name_2",
			URL:  "__url",
			Path: "/path/2",
		},
		{
			Name: "__name_1",
			URL:  "__url",
			Path: "/path/2",
		},
		{
			Name: "__name_0",
			URL:  "__url",
			Path: "/path/0",
		},
		{
			Name: "__name_1",
			URL:  "__url",
			Path: "/path/1",
		},
	}
	sortUsingSort := repos
	sort.Sort(sortUsingSort)
	s.Equal("__name_0", sortUsingSort[0].Name)
	s.Equal("/path/0", sortUsingSort[0].Path)
	s.Equal("__name_1", sortUsingSort[1].Name)
	s.Equal("/path/1", sortUsingSort[1].Path)
	s.Equal("__name_1", sortUsingSort[2].Name)
	s.Equal("/path/2", sortUsingSort[2].Path)
	s.Equal("__name_2", sortUsingSort[3].Name)
	s.Equal("/path/2", sortUsingSort[3].Path)

	sortUsingDotSort := repos
	sortUsingDotSort.Sort()
	s.Equal("__name_0", sortUsingDotSort[0].Name)
	s.Equal("/path/0", sortUsingDotSort[0].Path)
	s.Equal("__name_1", sortUsingDotSort[1].Name)
	s.Equal("/path/1", sortUsingDotSort[1].Path)
	s.Equal("__name_1", sortUsingDotSort[2].Name)
	s.Equal("/path/2", sortUsingDotSort[2].Path)
	s.Equal("__name_2", sortUsingDotSort[3].Name)
	s.Equal("/path/2", sortUsingDotSort[3].Path)
}
