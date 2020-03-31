package utils

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UnzipTests struct {
	suite.Suite
	observedPath string
}

func TestUnzip(t *testing.T) {
	suite.Run(t, &UnzipTests{})
}

func (s *UnzipTests) SetupTest() {
	s.observedPath = "./tests/zipfiles/observed"
	s.Nil(os.RemoveAll(s.observedPath))
	observedPath, err := os.Stat(s.observedPath)
	if os.IsNotExist(err) {
		s.Nil(os.MkdirAll(s.observedPath, os.ModePerm))
		observedPath, err = os.Stat(s.observedPath)
	}
	s.Nil(err)
	s.True(observedPath.IsDir())
}

func (s *UnzipTests) TestUnzip() {
	set01path := path.Join(s.observedPath, "/set01")
	s.Nil(os.MkdirAll(set01path, os.ModePerm))
	Unzip(UnzipOptions{
		InputPath:  "./tests/zipfiles/set01.zip",
		OutputPath: set01path,
	})
	s.Nil(os.RemoveAll(set01path))
}
