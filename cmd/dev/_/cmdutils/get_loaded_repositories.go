package cmdutils

import (
	"github.com/zephinzer/dev/internal/config"
	"github.com/zephinzer/dev/internal/log"
)

// GetLoadedRepositories retrieves repositories that have already been loaded and
// returns a `map[string]string` where the key value equals the local path of the
// repository and the value equals the full absolute path of the configuration file
func GetLoadedRepositories() map[string]string {
	alreadyPresentRepositories := map[string]string{}
	loadedConfigurationPaths := GetLoadedConfigFilePaths()
	for _, loadedConfigurationPath := range loadedConfigurationPaths {
		repos := config.Loaded[loadedConfigurationPath].Repositories
		for _, repo := range repos {
			repoPath, getPathError := repo.GetPath()
			if getPathError != nil {
				log.Warnf("failed to get repository path for '%s': %s", repo.URL, getPathError)
				continue
			}
			alreadyPresentRepositories[repoPath] = loadedConfigurationPath
		}
	}
	return alreadyPresentRepositories
}
