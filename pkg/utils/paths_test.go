package utils

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PathsTests struct {
	suite.Suite
}

func TestPaths(t *testing.T) {
	suite.Run(t, &PathsTests{})
}

func (s *PathsTests) Test_IsPathType() {
	cwd, err := os.Getwd()
	s.Nil(err)
	if err != nil {
		return
	}
	fileInfo, err := os.Lstat(cwd)
	s.Nil(err)
	if err != nil {
		return
	}
	s.True(IsPathType(fileInfo, PathType("d")))
	s.False(IsPathType(fileInfo, PathType("f")))
	s.True(IsPathType(fileInfo, PathType("a")))
}

func (s *PathsTests) Test_PathExists() {
	cwd, err := os.Getwd()
	s.Nil(err)
	if err != nil {
		return
	}
	s.True(PathExists(PathType("d"), cwd))
	s.True(PathExists(PathType("dir"), cwd))
	s.True(PathExists(PathType("directory"), cwd))
	s.True(PathExists(PathType("folder"), cwd))
	s.False(PathExists(PathType("f"), cwd))
	s.False(PathExists(PathType("file"), cwd))
}

func (s *PathsTests) Test_ResolvePath() {
	userHomeDir, getUserHomeDirError := os.UserHomeDir()
	s.Nilf(getUserHomeDirError, "failed to get home directory: %s", getUserHomeDirError)
	if getUserHomeDirError != nil {
		return
	}
	currWorkDir, getCurrWorkDirError := os.Getwd()
	s.Nilf(getCurrWorkDirError, "failed to get working directory: %s", getCurrWorkDirError)
	if getCurrWorkDirError != nil {
		return
	}
	var testCases = map[string][]string{
		userHomeDir:                          []string{"~"},
		path.Join(userHomeDir, "/some/path"): []string{"~/some/path"},
		currWorkDir:                          []string{"."},
		path.Join(currWorkDir, "/some/path"): []string{"./some/path"},
	}
	for output, input := range testCases {
		observed, resolvePathError := ResolvePath(input...)
		s.Nilf(resolvePathError, "failed to resolve path: %s", resolvePathError)
		if resolvePathError != nil {
			continue
		}
		s.Equalf(output, observed, "expected ResolvePath(\"%s\") to resolve to \"%s\", but it did not", input, output)
	}
}

func (s *PathsTests) Test_ResolvePath_emptyInput() {
	wd, err := os.Getwd()
	s.Nil(err)
	if err != nil {
		return
	}
	resolvedPath, resolvePathError := ResolvePath("")
	s.Nil(resolvePathError)
	if resolvePathError != nil {
		return
	}
	s.Equal(wd, resolvedPath)
}
