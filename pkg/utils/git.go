package utils

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	gitconfig "github.com/go-git/go-git/v5/config"
	gitransport "github.com/go-git/go-git/v5/plumbing/transport"
)

// GitClone does a `git clone` on the provided :cloneURL into the directory at
// :localPath. If the remote is a bare repository, a new repository will be
// initialised at :localPath with the `origin` remote set to :cloneURL
func GitClone(cloneURL, localPath string) error {
	_, cloneError := git.PlainClone(localPath, false, &git.CloneOptions{URL: cloneURL})
	if cloneError != nil {
		if cloneError == gitransport.ErrEmptyRemoteRepository {
			return GitInit(cloneURL, localPath)
		}
		return fmt.Errorf("failed to clone repository from url '%s': %s", cloneURL, cloneError)
	}
	return nil
}

func GitInit(cloneURL, localPath string) error {
	initialisedRepo, initError := git.PlainInit(localPath, false)
	if initError != nil {
		return fmt.Errorf("failed to initialise repository at '%s': %s", localPath, initError)
	}
	_, createRemoteError := initialisedRepo.CreateRemote(&gitconfig.RemoteConfig{
		Name: "origin",
		URLs: []string{cloneURL},
	})
	if createRemoteError != nil {
		return fmt.Errorf("failed to set remote origin for repository at '%s' tracked at '%s': %s", localPath, cloneURL, createRemoteError)
	}
	return nil
}
