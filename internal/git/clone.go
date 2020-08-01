package git

import (
	"fmt"

	extgit "github.com/go-git/go-git/v5"
	extgitransport "github.com/go-git/go-git/v5/plumbing/transport"
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
