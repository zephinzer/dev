package workspace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const VSCodeFileExtension = ".code-workspace"

// VSCode represents the visual studio code .code-workspace file
// format
type VSCode struct {
	// Folders represents a list of projects
	Folders []VSCodeFolder `json:"folders"`
}

// VSCodeFolder represents a single folder in the visual studio
// code workspace format
type VSCodeFolder struct {
	// Name is a logical user-defined label
	Name string `json:"name"`
	// Path is the absolute/relative path to the target project directory
	Path string `json:"path"`
}

// ToJSON returns a string that can be used as the .code-workspace visual
// studio code workspace file
func (vsc VSCode) ToJSON() (string, error) {
	asJSON, marshalError := json.MarshalIndent(vsc, "", "  ")
	if marshalError != nil {
		return "", fmt.Errorf("failed to marshal struct %v into JSON: %s", vsc, marshalError)
	}
	return string(asJSON), nil
}

// WriteTo writes the current instance of VSCode as JSON into a .code-workspace file
// located at `workspacePath`. If `overwrite` is not specified, the function will fail
// to complete if an existing file is found
func (vsc VSCode) WriteTo(workspacePath string, overwrite ...bool) error {
	isOverwritable := false
	if len(overwrite) > 0 {
		isOverwritable = overwrite[0]
	}

	pathToWorkspace := workspacePath

	if !path.IsAbs(pathToWorkspace) {
		workingDirectory, getWorkingDirectoryError := os.Getwd()
		if getWorkingDirectoryError != nil {
			return fmt.Errorf("failed to convert path '%s' to an absolute path: %s", pathToWorkspace, getWorkingDirectoryError)
		}
		pathToWorkspace = path.Join(workingDirectory, pathToWorkspace)
	}

	workspaceInfo, getWorkspaceInfoError := os.Stat(workspacePath)
	if getWorkspaceInfoError != nil {
		if !os.IsNotExist(getWorkspaceInfoError) {
			return fmt.Errorf("failed to get info about the path '%s': %s", workspacePath, getWorkspaceInfoError)
		}
	} else if workspaceInfo.IsDir() {
		return fmt.Errorf("failed to use path '%s': a directory already exists there", workspacePath)
	} else if !isOverwritable {
		return fmt.Errorf("failed to use path: '%s': a file already exists there", workspacePath)
	}

	workspaceData, toJSONError := vsc.ToJSON()
	if toJSONError != nil {
		return fmt.Errorf("failed to convert this instance of a VSCode workspace into JSON: %s", toJSONError)
	}
	if writeFileError := ioutil.WriteFile(workspacePath, []byte(workspaceData), os.ModePerm); writeFileError != nil {
		return fmt.Errorf("failed to write to the file at '%s': %s", workspacePath, writeFileError)
	}

	return nil
}
