package workspace

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/usvc/dev/internal/log"
	"github.com/usvc/dev/pkg/repository"
)

// Workspace represents a project-scoped set of code repos
type Workspace struct {
	Name         string                  `json:"name" yaml:"name"`
	Description  string                  `json:"description" yaml:"description"`
	Repositories []repository.Repository `json:"repositories" yaml:"repositories"`
}

type VSCode struct {
	Folders []VSCodeFolder `json:"folders"`
}

type VSCodeFolder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (ws Workspace) GetVSCode() (string, error) {
	homeDir, getHomeDirError := os.UserHomeDir()
	if getHomeDirError != nil {
		log.Errorf("unable to retrieve user's home directory: %s", getHomeDirError)
		os.Exit(1)
	}

	folders := []VSCodeFolder{}
	for _, repository := range ws.Repositories {
		repositoryPath, getPathError := repository.GetPath(homeDir)
		if getPathError != nil {
			return "", fmt.Errorf("failed to retrieve path of repository '%s': %s", repository.Name, getPathError)
		}
		folders = append(folders, VSCodeFolder{
			Name: repository.Name,
			Path: repositoryPath,
		})
	}
	vscodeWorkspace, marshalError := json.MarshalIndent(VSCode{Folders: folders}, "", "  ")
	if marshalError != nil {
		return "", fmt.Errorf("failed to generate JSON for vscode workspace: %s", marshalError)
	}
	return string(vscodeWorkspace) + "\n", nil
}
