package filesystem

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/pkg/utils/str"
)

type FindTests struct {
	suite.Suite
}

func TestFind(t *testing.T) {
	suite.Run(t, &FindTests{})
}

func (s *FindTests) TestFindParentContainingChildDirectory_inSameDirectory() {
	currentWorkingDirectory, getWdError := os.Getwd()
	s.Nil(getWdError, "failed to get working directory, skipping rest of test case")
	if getWdError != nil {
		return
	}
	workingDirectoryFileListing, readdirError := ioutil.ReadDir(currentWorkingDirectory)
	s.Nil(readdirError, "failed to get working directory file listing")
	if readdirError != nil {
		return
	}
	var theChosenOne string
	for _, file := range workingDirectoryFileListing {
		if file.IsDir() {
			theChosenOne = file.Name()
			break
		}
	}
	s.Truef(len(theChosenOne) > 0, "no files could be found in the working directory %s", currentWorkingDirectory)
	if str.IsEmpty(theChosenOne) {
		return
	}
	chosenDirectory, err := FindParentContainingChildDirectory(theChosenOne, currentWorkingDirectory)
	s.Nil(err)
	if err != nil {
		return
	}
	s.Equal(currentWorkingDirectory, chosenDirectory)
}

func (s *FindTests) TestFindParentContainingChildDirectory_notFound() {
	currentWorkingDirectory, getWdError := os.Getwd()
	s.Nil(getWdError, "failed to get working directory, skipping rest of test case")
	if getWdError != nil {
		return
	}
	workingDirectoryFileListing, readdirError := ioutil.ReadDir(currentWorkingDirectory)
	s.Nil(readdirError, "failed to get working directory file listing")
	if readdirError != nil {
		return
	}
	theChosenOne := "____probably-not-a-directory-name____"
	for _, file := range workingDirectoryFileListing {
		s.NotEqualf(file.Name(), theChosenOne, "expected %s to not be found, but it was", theChosenOne)
		if file.Name() == theChosenOne {
			return
		}
	}
	chosenDirectory, err := FindParentContainingChildDirectory(theChosenOne, currentWorkingDirectory)
	s.Nil(err)
	if err != nil {
		return
	}
	s.Len(chosenDirectory, 0)
}
