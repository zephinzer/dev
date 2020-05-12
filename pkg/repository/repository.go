package repository

import (
	"fmt"
	"path"

	"github.com/usvc/dev/pkg/utils"
	"github.com/usvc/dev/pkg/validator"
)

// Repository represents a code repo
type Repository struct {
	// Name is a user-defined string to identify this repository
	Name string `json:"name" yaml:"name"`
	// Workspaces is a list of strings that represent the name of the
	// logical workspace this repository belongs to
	Workspaces []string `json:"workspaces" yaml:"workspaces"`
	// Description is a user-defined block of text about what this repository
	// is for in their context
	Description string `json:"description" yaml:"description"`
	// CloneURL is the URL to use to clone the repository; if the provided
	// URL does not terminate with `.git`, a best-guess should be made to
	// convert this to a proper git clone URL
	CloneURL string `json:"cloneURL" yaml:"cloneURL"`
	// Path is the user-defined location to store this repository
	Path string `json:"path" yaml:"path"`
}

// GetPath returns the path where the repository should be stored;
// if the `.Path` property is defined, it shall be used, otherwise
// the returned path will be derived from the hostname and path of
// the `.CloneURL`
func (r Repository) GetPath(rootPath ...string) (string, error) {
	storagePath := "."
	if len(rootPath) > 0 {
		storagePath = rootPath[0]
	}

	if validator.IsGitHTTPUrl(r.CloneURL) || validator.IsGitSSHUrl(r.CloneURL) {
		parsedURL, parseError := validator.ParseURL(r.CloneURL)
		if parseError != nil {
			return "", fmt.Errorf("failed to parse clone url '%s'", r.CloneURL)
		}
		return path.Join(storagePath, parsedURL.Hostname, parsedURL.User, parsedURL.Path), nil
	} else if _, parseError := validator.ParseURL(r.CloneURL); parseError != nil {
		return "", fmt.Errorf("failed to parse url '%s'", r.CloneURL)
	} else {
		cloneURL, getCloneURLError := utils.GetSshCloneUrlFromHttpLinkUrl(r.CloneURL)
		if getCloneURLError != nil {
			return "", fmt.Errorf("failed to convert '%s' to a git SSH clone URL", r.CloneURL)
		}
		parsedURL, parseError := validator.ParseURL(cloneURL)
		if parseError != nil {
			return "", fmt.Errorf("failed to parse clone url '%s'", cloneURL)
		}
		return path.Join(storagePath, parsedURL.Hostname, parsedURL.User, parsedURL.Path), nil
	}
}
