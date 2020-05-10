package config

import (
	"io/ioutil"

	"github.com/usvc/dev/internal/link"
	"github.com/usvc/dev/pkg/software"
	"gopkg.in/yaml.v2"
)

// NewFromFile is a convenience function that reads the configuration
// in from a file at the provided :filePath
func NewFromFile(filePath string) (*File, error) {
	var config File

	configFile, readFileError := ioutil.ReadFile(filePath)
	if readFileError != nil {
		return nil, readFileError
	}

	unmarshalError := yaml.Unmarshal(configFile, &config)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &config, nil
}

type File struct {
	Dev       Dev                 `json:"dev" yaml:"dev"`
	Networks  Networks            `json:"networks" yaml:"networks"`
	Platforms Platforms           `json:"platforms" yaml:"platforms"`
	Softwares []software.Software `json:"softwares" yaml:"softwares"`
	Links     []link.Link         `json:"links" yaml:"links"`
}

// MergeWith merges the current File configuration instance with
// a second provided :other File configuration instance, intention for
// this is to merge a local configuration with a global configuration
func (f *File) MergeWith(other *File) {
	f.Dev = other.Dev

	seenMap := map[string]bool{}

	// copy over github configuration
	for _, githubAccount := range f.Platforms.Github.Accounts {
		seenMap["github"+githubAccount.AccessToken] = true
	}
	for _, githubAccount := range other.Platforms.Github.Accounts {
		if seenMap["github"+githubAccount.AccessToken] == true {
			continue
		}
		f.Platforms.Github.Accounts = append(f.Platforms.Github.Accounts, githubAccount)
		seenMap["github"+githubAccount.AccessToken] = true
	}

	// copy over gitlab configuration
	for _, gitlabAccount := range f.Platforms.Gitlab.Accounts {
		seenMap["gitlab"+gitlabAccount.AccessToken] = true
	}
	for _, gitlabAccount := range other.Platforms.Gitlab.Accounts {
		if seenMap["gitlab"+gitlabAccount.AccessToken] == true {
			continue
		}
		f.Platforms.Gitlab.Accounts = append(f.Platforms.Gitlab.Accounts, gitlabAccount)
		seenMap["gitlab"+gitlabAccount.AccessToken] = true
	}

	// copy over pivotal tracker configuration
	if len(f.Platforms.PivotalTracker.AccessToken) == 0 && len(other.Platforms.PivotalTracker.AccessToken) > 0 {
		f.Platforms.PivotalTracker.AccessToken = other.Platforms.PivotalTracker.AccessToken
	}
	for _, pivotalProject := range f.Platforms.PivotalTracker.Projects {
		seenMap["pivotal"+pivotalProject.ProjectID] = true
	}
	for _, pivotalProject := range other.Platforms.PivotalTracker.Projects {
		if seenMap["pivotal"+pivotalProject.ProjectID] == true {
			continue
		}
		f.Platforms.PivotalTracker.Projects = append(f.Platforms.PivotalTracker.Projects, pivotalProject)
		seenMap["pivotal"+pivotalProject.ProjectID] = true
	}

	// copy over trello configuration
	if len(f.Platforms.Trello.AccessToken) == 0 && len(other.Platforms.Trello.AccessToken) > 0 {
		f.Platforms.Trello.AccessToken = other.Platforms.Trello.AccessToken
	}
	if len(f.Platforms.Trello.AccessKey) == 0 && len(other.Platforms.Trello.AccessKey) > 0 {
		f.Platforms.Trello.AccessKey = other.Platforms.Trello.AccessKey
	}
	for _, trelloBoard := range f.Platforms.Trello.Boards {
		seenMap["trello"+trelloBoard.ID] = true
	}
	for _, trelloBoard := range other.Platforms.Trello.Boards {
		if seenMap["trello"+trelloBoard.ID] == true {
			continue
		}
		f.Platforms.Trello.Boards = append(f.Platforms.Trello.Boards, trelloBoard)
		seenMap["trello"+trelloBoard.ID] = true
	}

	// copy over softwares
	for _, software := range f.Softwares {
		seenMap["software"+software.Check.Command[0]] = true
	}
	for _, software := range other.Softwares {
		if seenMap["software"+software.Check.Command[0]] == true {
			continue
		}
		f.Softwares = append(f.Softwares, software)
		seenMap["software"+software.Check.Command[0]] = true
	}

	// copy over networks
	for _, network := range f.Networks {
		seenMap["network"+network.Check.URL] = true
	}
	for _, network := range other.Networks {
		if seenMap["network"+network.Check.URL] == true {
			continue
		}
		f.Networks = append(f.Networks, network)
		seenMap["network"+network.Check.URL] = true
	}
}
