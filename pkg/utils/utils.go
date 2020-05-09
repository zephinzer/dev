package utils

import (
	"os"
	"path"
)

func ContainsInt(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}

func convertPathToAbsolute(relativePathFragments ...string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	pathFragments := []string{cwd}
	pathFragments = append(pathFragments, relativePathFragments...)
	return path.Join(pathFragments...), nil
}
