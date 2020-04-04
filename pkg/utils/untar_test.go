package utils

import (
	"os"
	"path"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UntarTests struct {
	suite.Suite
	observedPath     string
	workingDirectory string
}

func TestUntar(t *testing.T) {
	suite.Run(t, &UntarTests{})
}

func (s *UntarTests) SetupTest() {
	var err error
	s.workingDirectory, err = os.Getwd()
	s.Nil(err)
	s.observedPath = "tests/zipfiles/observed"
	s.Nil(os.RemoveAll(s.observedPath))
	observedPath, err := os.Stat(s.observedPath)
	if os.IsNotExist(err) {
		s.Nil(os.MkdirAll(s.observedPath, os.ModePerm))
		observedPath, err = os.Stat(s.observedPath)
	}
	s.Nil(err)
	s.True(observedPath.IsDir())
}

func (s *UntarTests) TestUntar() {
	set01tarPath := "./tests/zipfiles/set01.tar.gz"
	set01path := path.Join(s.observedPath, "/set01")
	s.Nil(os.MkdirAll(set01path, os.ModePerm))
	var waiter sync.WaitGroup
	files := map[string]interface{}{}
	events := make(chan UntarEvent)
	go func() {
		for {
			e, ok := <-events
			if !ok {
				waiter.Done()
				return
			}
			if len(e.Path) > 0 {
				files[strings.ReplaceAll(e.Path, s.observedPath, ".")] = nil
			}
		}
	}()
	waiter.Add(1)
	Untar(UntarOptions{
		Events:     events,
		InputPath:  set01tarPath,
		OutputPath: set01path,
	})
	waiter.Wait()
	var observedPaths []string
	for key, _ := range files {
		observedPaths = append(observedPaths, key)
	}
	var expectedPaths = []string{"./set01/a", "./set01/b", "./set01/c", "./set01/d", "./set01/d/a", "./set01/d/b", "./set01/e", "./set01/e/a", "./set01/e/b", "./set01/f", "./set01/f/a", "./set01/f/b"}
	s.Equal(len(expectedPaths), len(observedPaths))
	for _, expectedPath := range expectedPaths {
		s.Contains(observedPaths, expectedPath)
	}
	s.Nil(os.RemoveAll(set01path))
}

func (s *UntarTests) TestUntar_noEvents() {
	defer func() {
		r := recover()
		s.Nil(r)
	}()
	set01path := path.Join(s.observedPath, "/set01")
	s.Nil(os.MkdirAll(set01path, os.ModePerm))
	errs := Untar(UntarOptions{
		InputPath:  "./tests/zipfiles/set01.tar.gz",
		OutputPath: set01path,
	})
	s.Nil(errs)
	s.Nil(os.RemoveAll(set01path))
}

func (s *UntarTests) TestUntar_invalidInputPath() {
	defer func() {
		r := recover()
		s.Nil(r)
	}()
	invalidInputPath := "`~!@#$%^&*()_+{}|-=[];':\\\",./<>?"
	set01path := path.Join(s.observedPath, "/set01")
	s.Nil(os.MkdirAll(set01path, os.ModePerm))
	errs := Untar(UntarOptions{
		InputPath:  invalidInputPath,
		OutputPath: set01path,
	})
	s.NotNil(errs)
	s.Len(errs, 1)
	s.Contains(errs[0].Error(), "no such file or directory")
	s.Nil(os.RemoveAll(set01path))
}

func (s *UntarTests) TestUntarStatus_GetPercentDoneByBytes() {
	us := UntarStatus{BytesTotal: 2, BytesProcessed: 1}
	s.Equal(float64(50), us.GetPercentDoneByBytes())
	us.BytesTotal = 3
	s.Less(us.GetPercentDoneByBytes(), float64(33.334))
	s.Greater(us.GetPercentDoneByBytes(), float64(33.333))
}

func (s *UntarTests) TestUntarStatus_GetPercentDoneByFiles() {
	us := UntarStatus{FilesTotalCount: 2, FilesProcessedCount: 1}
	s.Equal(float64(50), us.GetPercentDoneByFiles())
	us.FilesTotalCount = 3
	s.Less(us.GetPercentDoneByFiles(), float64(33.334))
	s.Greater(us.GetPercentDoneByFiles(), float64(33.333))
}
