package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func ResolvePath(relativePathFragments ...string) (string, error) {
	fullPath := path.Join(relativePathFragments...)

	// do home path resolution
	if fullPath[0] == '~' {
		userHomeDir, getUserHomeDirError := os.UserHomeDir()
		if getUserHomeDirError != nil {
			return "", fmt.Errorf("failed to resolve the home directory: %s", getUserHomeDirError)
		}
		fullPath = strings.Replace(fullPath, "~", userHomeDir, 1)
	}

	// resolve to absolute path
	if path.IsAbs(fullPath) {
		return fullPath, nil
	}

	currentWorkingDir, getWdError := os.Getwd()
	if getWdError != nil {
		return "", fmt.Errorf("failed to get working directory: %s", getWdError)
	}
	fullPath = path.Join(currentWorkingDir, fullPath)

	return fullPath, nil
}
