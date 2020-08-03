package repository

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
	pkgrepository "github.com/zephinzer/dev/pkg/repository"
)

type RepositoryTests struct {
	suite.Suite
}

func captureStdout(fromThis func() error) (string, error) {
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

func TestRepository(t *testing.T) {
	suite.Run(t, &RepositoryTests{})
}

func (s *RepositoryTests) Test_PromptForDescription() {
	repo := Repository{pkgrepository.Repository{URL: "__url"}}
	input := bytes.NewBuffer([]byte("__description"))
	outputString, err := captureStdout(func() error {
		return repo.PromptForDescription(input)
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Contains(outputString, "enter a description for '__url': ")
	s.Equal("__description", repo.Description)
}

func (s *RepositoryTests) Test_PromptForName() {
	expectedURL := "__url"
	expectedPath := "/path/to/somewhere"
	expectedBasePath := path.Base(expectedPath)
	repo := Repository{pkgrepository.Repository{Path: expectedPath, URL: expectedURL}}
	input := bytes.NewBuffer([]byte("__name"))
	outputString, err := captureStdout(func() error {
		return repo.PromptForName(input)
	})
	s.Nil(err)
	if err != nil {
		return
	}

	s.Contains(outputString, fmt.Sprintf("enter a name for '%s' (default: '%s')", expectedURL, expectedBasePath))
	s.Equal("__name", repo.Name)
}

func (s *RepositoryTests) Test_PromptForWorkspace() {
	repo := Repository{pkgrepository.Repository{URL: "__url"}}
	input := bytes.NewBuffer([]byte("__workspace1 , __workspace2, __workspace3"))
	outputString, err := captureStdout(func() error {
		return repo.PromptForWorkspaces(input)
	})
	s.Nil(err)
	if err != nil {
		return
	}
	s.Contains(outputString, "enter workspaces for '__url' (separate using commas):")
	s.Equal([]string{"__workspace1", "__workspace2", "__workspace3"}, repo.Workspaces)
}

func (s *RepositoryTests) Test_SetDescription() {
	expectedText := "__description with spaces"
	repo := Repository{}
	s.Len(repo.Description, 0)
	repo.SetDescription(expectedText)
	s.Len(repo.Description, len(expectedText))
}

func (s *RepositoryTests) Test_SetName() {
	expectedText := "__name"
	repo := Repository{}
	s.Len(repo.Name, 0)
	repo.SetName(expectedText)
	s.Len(repo.Name, len(expectedText))
}

func (s *RepositoryTests) Test_SetURL() {
	expectedText := "__url"
	repo := Repository{}
	s.Len(repo.URL, 0)
	repo.SetURL(expectedText)
	s.Len(repo.URL, len(expectedText))
}