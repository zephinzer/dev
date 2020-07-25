package cmdutils

import (
	"github.com/zephinzer/dev/internal/log"
	"github.com/zephinzer/dev/pkg/repository"
)

// FilterOutLoadedRepositories returns a slice of strings that has repository paths
// that do not already exist in the configuration :alreadyPresent argument
func FilterOutLoadedRepositories(
	repoURLsToAdd []string,
	alreadyPresent map[string]string,
) []string {
	finalRepoURLsToAdd := []string{}
	for _, repoURL := range repoURLsToAdd {
		repo := repository.Repository{URL: repoURL}
		repoPath, getPathError := repo.GetPath()
		if getPathError != nil {
			log.Warnf("failed to get repository path for '%s': %s", repo.URL, getPathError)
			continue
		}
		if value, ok := alreadyPresent[repoPath]; ok && len(value) > 0 {
			log.Warnf("repository '%s' already configured from '%s'", repoURL, value)
			continue
		}
		finalRepoURLsToAdd = append(finalRepoURLsToAdd, repoURL)
	}
	return finalRepoURLsToAdd
}
