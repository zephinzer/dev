package prompt

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zephinzer/dev/tests"
)

type InputHelperTests struct {
	suite.Suite
}

func TestInputHelper(t *testing.T) {
	suite.Run(t, &InputHelperTests{})
}

func (s *InputHelperTests) Test_PrintAfterMessage() {
	expected := "__after_message: hello world"
	helper := InputHelper(InputOptions{
		AfterMessage: expected,
	})
	output, err := tests.CaptureStdout(func() error {
		helper.PrintAfterMessage()
		return nil
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Contains(output, expected, "should contain the untruncated indicated message")
	s.NotContains(output, "\n", "should not contain newlines, leave that to the consumer")
}

func (s *InputHelperTests) Test_PrintBeforeMessage() {
	expected := "__before_message: hello world"
	helper := InputHelper(InputOptions{
		BeforeMessage: expected,
	})
	output, err := tests.CaptureStdout(func() error {
		helper.PrintBeforeMessage()
		return nil
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Contains(output, expected, "should contain the untruncated indicated message")
	s.NotContains(output, "\n", "should not contain newlines, leave that to the consumer")
}

func (s *InputHelperTests) Test_PrintOptions() {
	helper := InputHelper(InputOptions{
		SerializedOptions: []string{"__opt_0", "__opt_1", "__opt_2"},
	})
	output, err := tests.CaptureStdout(func() error {
		helper.PrintOptions()
		return nil
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Contains(output, "1. __opt_0")
	s.Contains(output, "2. __opt_1")
	s.Contains(output, "3. __opt_2")
}

func (s *InputHelperTests) Test_ReadInput_empty() {
	input := bytes.NewBuffer([]byte(""))
	helper := InputHelper(InputOptions{
		Reader: input,
	})
	helper.ReadInput()
	s.Equal("", helper.GetData())
}

func (s *InputHelperTests) Test_ReadInput_string() {
	expected := "hello world"
	input := bytes.NewBuffer([]byte(expected))
	helper := InputHelper(InputOptions{
		Reader: input,
	})
	helper.ReadInput()
	s.Equal(expected, helper.GetData())
}

func (s *InputHelperTests) Test_ReadInput_string_multiline() {
	expected := "hello"
	unexpected := fmt.Sprintf("%s\nworld", expected)
	input := bytes.NewBuffer([]byte(unexpected))
	helper := InputHelper(InputOptions{
		Reader: input,
	})
	helper.ReadInput()
	s.Equal(expected, helper.GetData())
}
