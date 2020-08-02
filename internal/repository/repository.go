package repository

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	pkgrepository "github.com/zephinzer/dev/pkg/repository"
	"github.com/zephinzer/dev/pkg/utils/str"
)

type Repository struct {
	pkgrepository.Repository
}

func (r *Repository) getInput(using ...io.Reader) (string, error) {
	var answer string
	var fromReader io.Reader = os.Stdin
	if len(using) > 0 {
		fromReader = using[0]
	}
	scanner := bufio.NewScanner(fromReader)
	if scanner.Scan() {
		answer = scanner.Text()
	}
	if scanError := scanner.Err(); scanError != nil {
		return "", fmt.Errorf("failed to get input from tty: %s", scanError)
	}
	return answer, nil
}

func (r *Repository) PromptForDescription(reader ...io.Reader) error {
	fmt.Printf("\033[1menter a description for '%s': \033[0m", r.URL)
	var err error
	r.Description, err = r.getInput(reader...)
	return err
}

func (r *Repository) PromptForName(reader ...io.Reader) error {
	repoPath, getPathError := r.GetPath()
	if getPathError != nil {
		return fmt.Errorf("failed to get repository path: %s", getPathError)
	}
	defaultName := path.Base(repoPath)
	fmt.Printf("\033[1menter a name for '%s' (default: '%s'): \033[0m", r.URL, defaultName)
	var err error
	r.Name, err = r.getInput(reader...)
	if err != nil {
		return err
	}
	if str.IsEmpty(r.Name) {
		r.Name = defaultName
	}
	return nil
}

func (r *Repository) PromptForWorkspaces(reader ...io.Reader) error {
	fmt.Printf("\033[1menter workspaces for '%s' (separate using commas): \033[0m", r.URL)
	var answer string
	var err error
	answer, err = r.getInput(reader...)
	if err != nil {
		return err
	} else if len(answer) == 0 {
		return nil
	}
	workspaces := strings.Split(answer, ",")
	for i := 0; i < len(workspaces); i++ {
		workspaces[i] = strings.TrimSpace(workspaces[i])
	}
	r.Workspaces = workspaces
	return nil
}

func (r *Repository) SetDescription(repoDescription string) {
	r.Description = repoDescription
}

func (r *Repository) SetName(repoName string) {
	r.Name = repoName
}

func (r *Repository) SetURL(repoURL string) {
	r.URL = repoURL
}

func (r *Repository) ToRepository() pkgrepository.Repository {
	return pkgrepository.Repository{
		Name:        r.Name,
		Description: r.Description,
		URL:         r.URL,
		Workspaces:  r.Workspaces,
	}
}
