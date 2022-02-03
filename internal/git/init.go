package git

import (
	"fmt"

	extgit "github.com/go-git/go-git/v5"
	extgitconfig "github.com/go-git/go-git/v5/config"
)

// Init does a `git init` on the provided :localPath using the clone URL :cloneURL
// as the initial remote. If :defaultRemoteName is specified, the first string in the expansion
// will be used as the remote name; if unspecified, 'origin' is used
func Init(cloneURL, localPath string, defaultRemoteName ...string) error {
	initialisedRepo, initError := extgit.PlainInit(localPath, false)
	if initError != nil {
		return fmt.Errorf("failed to initialise repository at '%s': %s", localPath, initError)
	}
	remoteOfInterest := DefaultRemote
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
