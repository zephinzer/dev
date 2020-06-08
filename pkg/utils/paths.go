package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type PathType string

func (pt PathType) String() string {
	switch strings.ToLower(string(pt)) {
	case "f":
		fallthrough
	case "file":
		return "file"
	case "d":
		fallthrough
	case "dir":
		fallthrough
	case "directory":
		fallthrough
	case "folder":
		return "directory"
	}
	return "any"
}

func IsPathType(pathInfo os.FileInfo, pathType PathType) bool {
	switch pathType.String() {
	case "file":
		return !pathInfo.IsDir()
	case "directory":
		return pathInfo.IsDir()
	case "any":
		return true
	}
	return false
}

func PathExists(asType PathType, pathFragments ...string) (bool, error) {
	fullPath := path.Join(pathFragments...)

	pathInfo, statError := os.Stat(fullPath)
	if statError != nil {
		if os.IsNotExist(statError) {
			return false, nil
		}
		return false, fmt.Errorf("failed to get status of path '%s': %s", fullPath, statError)
	}

	return IsPathType(pathInfo, asType), nil
}

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

	// resolve symlink if it is one
	fileInfo, lstatError := os.Lstat(fullPath)
	if lstatError != nil && !os.IsNotExist(lstatError) {
		return "", fmt.Errorf("failed to stat file at '%s': %s", fullPath, lstatError)
	}
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		resolvedSymlinkPath, _ := os.Readlink(fullPath)
		if len(resolvedSymlinkPath) > 0 {
			return ResolvePath(resolvedSymlinkPath)
		}
	}

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
