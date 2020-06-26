package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type OpenTests struct {
	suite.Suite
}

func TestOpen(t *testing.T) {
	suite.Run(t, &OpenTests{})
}

func (s *OpenTests) Test_getDefaultOpenURICommand() {
	targetURI := "a"
	command, err := getDefaultOpenURICommand("linux", targetURI)
	s.Nil(err)
	s.EqualValues([]string{"xdg-open", targetURI}, command)
	command, err = getDefaultOpenURICommand("macos", targetURI)
	s.Nil(err)
	s.EqualValues([]string{"open", targetURI}, command)
	command, err = getDefaultOpenURICommand("windows", targetURI)
	s.Nil(err)
	s.EqualValues([]string{"System32/rundll32.exe", "url.dll,FileProtocolHandler", targetURI}, command)
	command, err = getDefaultOpenURICommand("x", targetURI)
	s.NotNil(err)
}
