package repository

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/zephinzer/dev/internal/prompt"
	pkgrepository "github.com/zephinzer/dev/pkg/repository"
	"github.com/zephinzer/dev/pkg/utils/defaults"
	"github.com/zephinzer/dev/pkg/utils/str"
)

// Repository wraps the canonical Repository class (in pkg/repository) and provides
// interface specific getters/setters for CLI usage
type Repository struct {
	pkgrepository.Repository
}

// PromptForDescription makes a request to the user to enter in a description for this
// Repository instance
func (r *Repository) PromptForDescription(useOtherReader ...io.Reader) error {
	// decide on defaults
	var reader io.Reader = defaults.GetIoReader(os.Stdin, useOtherReader...)

	// go for it
	var promptForStringErr error
	r.Description, promptForStringErr = prompt.ForString(prompt.InputOptions{
		BeforeMessage: fmt.Sprintf("enter a description for '%s': ", r.URL),
		Reader:        reader,
	})
	return promptForStringErr
}

// PromptForName makes a request to the user to enter in a name for this Repository
// instance. Defaults to the name of the directory if no name is entered
func (r *Repository) PromptForName(useOtherReader ...io.Reader) error {
	// init
	repoPath, getPathError := r.GetPath()
	if getPathError != nil {
		return fmt.Errorf("failed to get repository path: %s", getPathError)
	}

	// decide on defaults
	defaultName := path.Base(repoPath)
	var reader io.Reader = defaults.GetIoReader(os.Stdin, useOtherReader...)

	// go for it
	var promptForStringErr error
	if r.Name, promptForStringErr = prompt.ForString(prompt.InputOptions{
		BeforeMessage: fmt.Sprintf("enter a name for '%s' (default: '%s'): ", r.URL, defaultName),
		Reader:        reader,
	}); promptForStringErr != nil {
		return fmt.Errorf("failed to receive input for repository name: %s", promptForStringErr)
	}

	// assign defaults if not present
	if str.IsEmpty(r.Name) {
		r.Name = defaultName
	}
	return nil
}

// PromptForWorkspaces makes a request to the user to enter in a comma-separated string that
// indicates which workspaces they would like to add the repository to. Defaults to not having
// a list of workspaces.
func (r *Repository) PromptForWorkspaces(useOtherReader ...io.Reader) error {
	//init
	var reader io.Reader = defaults.GetIoReader(os.Stdin, useOtherReader...)

	// go for it
	var response string
	var promptForStringErr error
	if response, promptForStringErr = prompt.ForString(prompt.InputOptions{
		BeforeMessage: fmt.Sprintf("enter workspaces for '%s' (separate using commas): ", r.URL),
		Reader:        reader,
	}); promptForStringErr != nil {
		return fmt.Errorf("failed to receive input for workspaces: %s", promptForStringErr)
	} else if str.IsEmpty(response) {
		return nil
	}

	// post-processing into a slice
	workspaces := strings.Split(response, ",")
	for i := 0; i < len(workspaces); i++ {
		workspaces[i] = strings.TrimSpace(workspaces[i])
	}
	r.Workspaces = workspaces
	return nil
}

// SetDescription is a setter method for the .Description property
func (r *Repository) SetDescription(repoDescription string) {
	r.Description = repoDescription
}

// SetName is a setter method for the .Name property
func (r *Repository) SetName(repoName string) {
	r.Name = repoName
}

// SetURL is a setter method for the .URL property
func (r *Repository) SetURL(repoURL string) {
	r.URL = repoURL
}

// ToRepository returns this Repository isntance as a canonical Repository
func (r *Repository) ToRepository() pkgrepository.Repository {
	return pkgrepository.Repository{
		Name:        r.Name,
		Description: r.Description,
		URL:         r.URL,
		Workspaces:  r.Workspaces,
	}
}
