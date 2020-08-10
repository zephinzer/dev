package prompt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/tests"
)

type ForStringTests struct {
	suite.Suite
}

func TestForString(t *testing.T) {
	suite.Run(t, &ForStringTests{})
}

func (s *ForStringTests) Test_ForString() {
	var input string
	var err error
	output, err := tests.CaptureStdout(func() error {
		input, err = ForString(InputOptions{
			BeforeMessage: "__before_message",
			Reader:        bytes.NewBuffer([]byte("hello world")),
		})
		return err
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Contains(output, "__before_message", "should print the specified Beforemessage")
	s.Contains(input, "hello world")
}

func (s *ForStringTests) Test_ForString_withDefault() {
	var input string
	var err error
	_, err = tests.CaptureStdout(func() error {
		input, err = ForString(InputOptions{
			Reader: bytes.NewBuffer([]byte("")),
		}, "__default_string")
		return err
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Contains(input, "__default_string")
}
