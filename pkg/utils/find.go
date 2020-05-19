package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// FindParentContainingChildDirectory continously ascends from the initial
// directory path at `startingFrom` and checks for the existence of a child
// directory named `name` for `levels[0]` levels.
//
// On successfully finding such a child directory, it returns the path of the
// parent directory as the first argument, on failing to find, it returns an
// empty string. In both cases, the returned `error`-typed argument will be nil.
//
// If the function failed to complete, the second `error`-typed argument will
// be non-nil.
func FindParentContainingChildDirectory(name, startingFrom string, levels ...int) (string, error) {
	childDirectoryOfInterest := name

	// resolve to absolute if not already
	directoryPath := startingFrom
	if !path.IsAbs(startingFrom) {
		cwd, getCwdError := os.Getwd()
		if getCwdError != nil {
			return "", fmt.Errorf("failed to retrieve working directory: %s", getCwdError)
		}
		directoryPath = path.Join(cwd, startingFrom)
	}

	// prepare number of directories to ascend to
	searchDepth := -1
	if len(levels) > 0 {
		searchDepth = levels[0]
	}

	splitPath := strings.Split(directoryPath, string(filepath.Separator))
	for len(splitPath) >= 2 && searchDepth != 0 {
		fileListing, readDirError := ioutil.ReadDir(directoryPath)
		if readDirError != nil {
			return "", fmt.Errorf("failed to read directory at '%s': %s", directoryPath, readDirError)
		}
		for _, file := range fileListing {
			if file.IsDir() && file.Name() == childDirectoryOfInterest {
				return directoryPath, nil
			}
		}
		directoryPath = path.Dir(directoryPath)
		splitPath = strings.Split(directoryPath, string(filepath.Separator))
		searchDepth--
	}

	return "", nil
}
