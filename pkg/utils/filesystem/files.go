package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
)

// EnsureDirectoryExists ensures that a directory exists at
// :pathToDirectory, returning an error only if a directory
// cannot be ensured to exist there
func EnsureDirectoryExists(pathToDirectory string) error {
	if mkdirError := os.MkdirAll(pathToDirectory, os.ModePerm); mkdirError != nil {
		return fmt.Errorf("failed to ensure directory exists at '%s': %s", pathToDirectory, mkdirError)
	}
	return nil
}

// IsDirectoryEmpty returns whether the provided directory located
// at :pathToDirectory is empty or not. If the error return is not
// nil, this indicates a system-type error
func IsDirectoryEmpty(pathToDirectory string) (bool, error) {
	directoryListing, readdirError := ioutil.ReadDir(pathToDirectory)
	if readdirError != nil {
		return false, fmt.Errorf("failed to read directory at %s", pathToDirectory)
	}
	if len(directoryListing) > 0 {
		listOfFiles := []string{}
		for _, file := range directoryListing {
			listOfFiles = append(listOfFiles, file.Name())
		}
		return false, nil
	}
	return true, nil
}
