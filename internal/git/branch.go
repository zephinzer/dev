package git

import (
	"fmt"

	extgit "github.com/go-git/go-git/v5"
)

// GetCurrentBranch retrieves the branch we are currently on
// in a Git repository
func GetCurrentBranch(repoPath string) (string, error) {
	repo, err := extgit.PlainOpen(repoPath)
	if err != nil {
		return "", fmt.Errorf("failed to open repository at '%s': %s", repoPath, err)
	}
	head, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("failed to get info on repository at '%s': %s", repoPath, err)
	}
	return head.Name().Short(), nil
}
