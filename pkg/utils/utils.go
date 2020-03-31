package utils

import (
	"os"
	"path"
)

func convertPathToAbsolute(relativePathFragments ...string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	pathFragments := []string{cwd}
	pathFragments = append(pathFragments, relativePathFragments...)
	return path.Join(pathFragments...), nil
}
