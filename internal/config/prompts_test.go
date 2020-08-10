package config

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PromptsTests struct {
	suite.Suite
}

func (s PromptsTests) captureStdout(fromThis func() error) (string, error) {
	originalStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	defer func() {
		os.Stdout = originalStdout
	}()
	err := fromThis()
	if err != nil {
		return "", err
	}
	output := make(chan string)
	defer close(output)
	go func() {
		var b bytes.Buffer
		io.Copy(&b, reader)
		output <- b.String()
	}()
	writer.Close()
	return <-output, nil
}

func TestPrompts(t *testing.T) {
	suite.Run(t, &PromptsTests{})
}

func (s *PromptsTests) Test_PromptSelectLoadedConfiguration_onlyOneConfiguration() {
	var output string
	var err error
	input := bytes.NewBuffer([]byte("0"))
	s.captureStdout(func() error {
		Loaded = map[string]Config{
			"__config_0": Config{},
		}
		output, err = PromptSelectLoadedConfiguration("__test_prompt_message", input)
		s.Nil(err)
		return nil
	})
	s.Equal("__config_0", output)
}

func (s *PromptsTests) Test_PromptSelectLoadedConfiguration_moreThanOneConfiguration() {
	// init
	configName := []string{"__config_0", "__config_1"}
	Loaded = map[string]Config{
		configName[0]: Config{},
		configName[1]: Config{},
	}
	var input *bytes.Buffer
	var output string
	var err error

	// testing for selection of skipping
	input = bytes.NewBuffer([]byte("0"))
	s.captureStdout(func() error {
		output, err = PromptSelectLoadedConfiguration("__test_prompt_message", input)
		s.Nil(err)
		return nil
	})
	s.Equal("", output, "on input of 0, this should be skipped")
	input = bytes.NewBuffer([]byte(""))
	s.captureStdout(func() error {
		output, err = PromptSelectLoadedConfiguration("__test_prompt_message", input)
		s.Nil(err)
		return nil
	})
	s.Equal("", output, "on empty input, this should be skipped")

	// testing for selection of number 1
	input = bytes.NewBuffer([]byte("1"))
	s.captureStdout(func() error {
		output, err = PromptSelectLoadedConfiguration("__test_prompt_message", input)
		s.Nil(err)
		return nil
	})
	s.Equal(configName[0], output, "a one-based index should be used to select configurations")

	// testing for selection of number 2
	input = bytes.NewBuffer([]byte("2"))
	s.captureStdout(func() error {
		output, err = PromptSelectLoadedConfiguration("__test_prompt_message", input)
		s.Nil(err)
		return nil
	})
	s.Equal(configName[1], output, "a one-based index should be used to select configurations")
}
