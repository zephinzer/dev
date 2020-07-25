package git

import (
	"fmt"
	"os"
	"path"
	"regexp"

	extgit "github.com/go-git/go-git/v5"
	extgitconfig "github.com/go-git/go-git/v5/config"
	extgitransport "github.com/go-git/go-git/v5/plumbing/transport"
)

const (
	DefaultRemoteOfInterest = "origin"
)

// Clone does a `git clone` on the provided :cloneURL into the directory at
// :localPath. If the remote is a bare repository, a new repository will be
// initialised at :localPath with the `origin` remote set to :cloneURL
func Clone(cloneURL, localPath string) error {
	_, cloneError := extgit.PlainClone(localPath, false, &extgit.CloneOptions{URL: cloneURL})
	if cloneError != nil {
		if cloneError == extgitransport.ErrEmptyRemoteRepository {
			return Init(cloneURL, localPath)
		}
		return fmt.Errorf("failed to clone repository from url '%s': %s", cloneURL, cloneError)
	}
	return nil
}

// Init does a `git init` on the provided :localPath using the clone URL :cloneURL
// as the initial remote. If :defaultRemoteName is specified, the first string in the expansion
// will be used as the remote name; if unspecified, 'origin' is used
func Init(cloneURL, localPath string, defaultRemoteName ...string) error {
	initialisedRepo, initError := extgit.PlainInit(localPath, false)
	if initError != nil {
		return fmt.Errorf("failed to initialise repository at '%s': %s", localPath, initError)
	}
	remoteOfInterest := DefaultRemoteOfInterest
	if len(defaultRemoteName) > 0 {
		remoteOfInterest = defaultRemoteName[0]
	}
	_, createRemoteError := initialisedRepo.CreateRemote(&extgitconfig.RemoteConfig{
		Name: remoteOfInterest,
		URLs: []string{cloneURL},
	})
	if createRemoteError != nil {
		return fmt.Errorf("failed to set remote origin for repository at '%s' tracked at '%s': %s", localPath, cloneURL, createRemoteError)
	}
	return nil
}

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

	remoteNameOfInterest := DefaultRemoteOfInterest
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
