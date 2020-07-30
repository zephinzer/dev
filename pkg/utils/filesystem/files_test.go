package filesystem

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FilesTests struct {
	suite.Suite
}

func TestFiles(t *testing.T) {
	suite.Run(t, &FilesTests{})
}

func (s *FilesTests) TestEnsureDirectoryExists() {
	ensureDirectoryExistsError := EnsureDirectoryExists("./tests/should-exist")
	s.Nil(ensureDirectoryExistsError)
}

func (s *FilesTests) TestIsDirectoryEmpty() {
	ensureDirectoryExistsError := EnsureDirectoryExists("./tests/is-directory-empty")
	s.Nil(ensureDirectoryExistsError)
	if ensureDirectoryExistsError != nil {
		return
	}
	isDirectoryEmpty, isDirectoryEmptyError := IsDirectoryEmpty("./tests/is-directory-empty")
	s.Nil(isDirectoryEmptyError)
	if isDirectoryEmptyError != nil {
		return
	}
	s.True(isDirectoryEmpty)
}

func (s *FilesTests) TestIsDirectoryEmpty_whenNotEmpty() {
	testDirPath := "./tests/is-directory-empty-not-empty"
	ensureDirectoryExistsError := EnsureDirectoryExists(testDirPath)
	s.Nil(ensureDirectoryExistsError)
	if ensureDirectoryExistsError != nil {
		return
	}
	writeFileError := ioutil.WriteFile(path.Join(testDirPath, "/some-file"), []byte("testing is directory empty when it's not"), os.ModePerm)
	s.Nil(writeFileError)
	if writeFileError != nil {
		return
	}
	isDirectoryEmpty, isDirectoryEmptyError := IsDirectoryEmpty(testDirPath)
	s.Nil(isDirectoryEmptyError)
	if isDirectoryEmptyError != nil {
		return
	}
	s.False(isDirectoryEmpty)
}
