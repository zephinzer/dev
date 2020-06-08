package repository

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	pkgrepository "github.com/zephinzer/dev/pkg/repository"
	"github.com/zephinzer/dev/pkg/utils"
)

type Repository struct {
	pkgrepository.Repository
}

func (r *Repository) PromptForDescription() error {
	fmt.Printf("\033[1menter a description for '%s': \033[0m", r.URL)
	var answer string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer = scanner.Text()
	}
	if scanError := scanner.Err(); scanError != nil {
		return fmt.Errorf("an unexpected error occurred: %s", scanError)
	}
	r.Description = answer
	return nil
}

func (r *Repository) PromptForName() error {
	repoPath, getPathError := r.GetPath()
	if getPathError != nil {
		return fmt.Errorf("failed to get repository path: %s", getPathError)
	}
	defaultName := path.Base(repoPath)
	fmt.Printf("\033[1menter a name for '%s' (default: '%s'): \033[0m", r.URL, defaultName)
	var answer string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer = scanner.Text()
	}
	if scanError := scanner.Err(); scanError != nil {
		return fmt.Errorf("an unexpected error occurred: %s", scanError)
	}
	r.Name = answer
	if utils.IsEmptyString(r.Name) {
		r.Name = defaultName
	}
	return nil
}

func (r *Repository) PromptForWorkspaces() error {
	fmt.Printf("\033[1menter workspaces for '%s' (separate using commas): \033[0m", r.URL)
	var answer string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer = scanner.Text()
	}
	if scanError := scanner.Err(); scanError != nil {
		return fmt.Errorf("an unexpected error occurred: %s", scanError)
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
