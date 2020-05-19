package workspace

import (
	"fmt"
	"os"

	"github.com/zephinzer/dev/pkg/repository"
)

// Workspace represents a project-scoped set of code repos
type Workspace struct {
	Name         string                  `json:"name" yaml:"name"`
	Description  string                  `json:"description" yaml:"description"`
	Repositories []repository.Repository `json:"repositories" yaml:"repositories"`
}

func (ws Workspace) ToVSCode() (*VSCode, error) {
	homeDir, getHomeDirError := os.UserHomeDir()
	if getHomeDirError != nil {
		return nil, fmt.Errorf("unable to retrieve user's home directory: %s", getHomeDirError)
	}

	folders := []VSCodeFolder{}
	for _, repository := range ws.Repositories {
		repositoryPath, getPathError := repository.GetPath(homeDir)
		if getPathError != nil {
			return nil, fmt.Errorf("failed to retrieve path of repository '%s': %s", repository.Name, getPathError)
		}
		folders = append(folders, VSCodeFolder{
			Name: repository.Name,
			Path: repositoryPath,
		})
	}

	return &VSCode{Folders: folders}, nil
}
