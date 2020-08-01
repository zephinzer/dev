package repository

import (
	"fmt"
	"path"

	"github.com/zephinzer/dev/pkg/utils/str"
	"github.com/zephinzer/dev/pkg/validator"
)

// Repository represents a code repo
type Repository struct {
	// Description is a user-defined block of text about what this repository
	// is for in their context
	Description string `json:"description" yaml:"description,omitempty"`
	// Name is a user-defined string to identify this repository
	Name string `json:"name" yaml:"name,omitempty"`
	// Path is the user-defined location to store this repository
	Path string `json:"path" yaml:"path,omitempty"`
	// URL is the URL to use to clone the repository; if the provided
	// URL does not terminate with `.git`, a best-guess should be made to
	// convert this to a proper git clone URL
	URL string `json:"url" yaml:"url,omitempty"`
	// Workspaces is a list of strings that represent the name of the
	// logical workspace this repository belongs to
	Workspaces []string `json:"workspaces" yaml:"workspaces,omitempty,flow"`
}

// GetPath returns the path where the repository should be stored;
// if the `.Path` property is defined, it shall be used, otherwise
// the returned path will be derived from the hostname and path of
// the `.URL`
func (r Repository) GetPath(rootPath ...string) (string, error) {
	// if :rootPath is defined, use that as the root instead of the current directory
	storagePath := "."
	if len(rootPath) > 0 {
		storagePath = rootPath[0]
	}

	// if .Path is already defined, juse use that
	if len(r.Path) > 0 {
		return path.Join(storagePath, r.Path), nil
	}

	if validator.IsGitHTTPUrl(r.URL) || validator.IsGitSSHUrl(r.URL) {
		parsedURL, parseError := validator.ParseURL(r.URL)
		if parseError != nil {
			return "", fmt.Errorf("failed to parse clone url '%s'", r.URL)
		}
		return path.Join(storagePath, parsedURL.Hostname, parsedURL.User, parsedURL.Path), nil
	} else if _, parseError := validator.ParseURL(r.URL); parseError != nil {
		return "", fmt.Errorf("failed to parse url '%s': %s", r.URL, parseError)
	} else {
		URL, getURLError := str.GetSshCloneUrlFromHttpLinkUrl(r.URL)
		if getURLError != nil {
			return "", fmt.Errorf("failed to convert '%s' to a git SSH clone URL", r.URL)
		}
		parsedURL, parseError := validator.ParseURL(URL)
		if parseError != nil {
			return "", fmt.Errorf("failed to parse clone url '%s'", URL)
		}
		return path.Join(storagePath, parsedURL.Hostname, parsedURL.User, parsedURL.Path), nil
	}
}

func (r Repository) GetWebsiteURL() (string, error) {
	switch true {
	case validator.IsGitHTTPUrl(r.URL):
		link, err := str.GetHttpLinkFromHttpCloneUrl(r.URL)
		if err != nil {
			return "", err
		}
		return link, nil
	case validator.IsGitSSHUrl(r.URL):
		link, err := str.GetHttpLinkFromSshCloneUrl(r.URL)
		if err != nil {
			return "", err
		}
		return link, nil
	}
	parsedURL, parseURLError := validator.ParseURL(r.URL)
	if parseURLError != nil {
		return "", parseURLError
	}
	return parsedURL.String(), nil
}
