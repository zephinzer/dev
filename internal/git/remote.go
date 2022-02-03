package git

import (
	"fmt"
	"os"
	"path"
	"regexp"

	extgit "github.com/go-git/go-git/v5"
)

type Remote struct {
	Name string
	URL  string
}

// GetRemote retrieves the remote from the Git repository located at the path
// :fromGitRepositoryAt which should be a directory containing a .git directory. If
// :remoteNameMatcher is specified, the URL returned will be the URL of the remote
// whose name matches with the specified matcher. If :remoteNameMatcher is not specified,
// the URL of origin is returned
func GetRemote(fromGitRepositoryAt string, remoteNameMatcher ...string) (*Remote, error) {
	// validation checks
	gitPath := path.Join(fromGitRepositoryAt, "/.git")
	fileInfo, err := os.Lstat(gitPath)
	if err != nil {
		return nil, err
	} else if !fileInfo.IsDir() {
		return nil, fmt.Errorf("path at '%s' is not a valid git directory", gitPath)
	}

	remoteNameOfInterest := DefaultRemote
	if len(remoteNameMatcher) > 0 {
		remoteNameOfInterest = remoteNameMatcher[0]
	}

	// get the remote of interest
	repo, err := extgit.PlainOpen(fromGitRepositoryAt)
	if err != nil {
		return nil, err
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return nil, err
	}

	for _, remote := range remotes {
		remoteName := remote.Config().Name
		regex := regexp.MustCompile(remoteNameOfInterest)
		if regex.MatchString(remoteName) {
			return &Remote{Name: remoteName, URL: remote.Config().URLs[0]}, nil
		}
	}

	return nil, fmt.Errorf("failed to find remote with name that matches '%s'", remoteNameOfInterest)
}
