package repository

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

type RepositoryTests struct {
	suite.Suite
}

func (s RepositoryTests) toYAML(repo Repository) string {
	repoYAML, marshalError := yaml.Marshal(repo)
	s.Nil(marshalError)
	if marshalError != nil {
		return ""
	}
	return string(repoYAML)
}

func TestRepository(t *testing.T) {
	suite.Run(t, &RepositoryTests{})
}

func (s *RepositoryTests) TestYAMLMarshal_zeroValue() {
	repo := Repository{}
	repoYAML := s.toYAML(repo)
	if len(repoYAML) == 0 {
		return
	}
	s.Equal("{}\n", repoYAML)
}

func (s *RepositoryTests) TestYAMLMarshal_description() {
	repo := Repository{Description: "__test_description"}
	repoYAML := s.toYAML(repo)
	if len(repoYAML) == 0 {
		return
	}
	s.Equal("description: __test_description\n", repoYAML)
}

func (s *RepositoryTests) TestYAMLMarshal_name() {
	repo := Repository{Name: "__test_name"}
	repoYAML := s.toYAML(repo)
	if len(repoYAML) == 0 {
		return
	}
	s.Equal("name: __test_name\n", repoYAML)
}

func (s *RepositoryTests) TestYAMLMarshal_path() {
	repo := Repository{Path: "__test_path"}
	repoYAML := s.toYAML(repo)
	if len(repoYAML) == 0 {
		return
	}
	s.Equal("path: __test_path\n", repoYAML)
}

func (s *RepositoryTests) TestYAMLMarshal_url() {
	repo := Repository{URL: "__test_url"}
	repoYAML := s.toYAML(repo)
	if len(repoYAML) == 0 {
		return
	}
	s.Equal("url: __test_url\n", repoYAML)
}

func (s *RepositoryTests) TestYAMLMarshal_workspaces() {
	repo := Repository{Workspaces: []string{"__test_workspaces"}}
	repoYAML := s.toYAML(repo)
	if len(repoYAML) == 0 {
		return
	}
	s.Equal("workspaces:\n- __test_workspaces\n", repoYAML)
}

func (s *RepositoryTests) TestGetPath() {
	expectedPath := "gitlab.com/zephinzer/dev"
	equivalentPathRepoURLs := []string{
		"git@gitlab.com:zephinzer/dev.git",
		"https://gitlab.com/zephinzer/dev.git",
		"https://username@gitlab.com/zephinzer/dev.git",
		"https://username:password@gitlab.com/zephinzer/dev.git",
		"https://gitlab.com/zephinzer/dev",
	}
	for _, u := range equivalentPathRepoURLs {
		repo := Repository{URL: u}
		repoPath, getPathError := repo.GetPath()
		s.Nil(getPathError)
		if getPathError != nil {
			return
		}
		s.Equal(expectedPath, repoPath)
	}
}

func (s *RepositoryTests) TestGetWebsiteURL() {
	expectedURL := "https://gitlab.com/zephinzer/dev"
	equivalentPathRepoURLs := []string{
		"git@gitlab.com:zephinzer/dev.git",
		"https://gitlab.com/zephinzer/dev.git",
		"https://username@gitlab.com/zephinzer/dev.git",
		"https://username:password@gitlab.com/zephinzer/dev.git",
		"https://gitlab.com/zephinzer/dev",
	}
	for _, u := range equivalentPathRepoURLs {
		repo := Repository{URL: u}
		repoWebsiteURL, getWebsiteURLError := repo.GetWebsiteURL()
		s.Nil(getWebsiteURLError)
		if getWebsiteURLError != nil {
			return
		}
		s.Equal(expectedURL, repoWebsiteURL)
	}
}
