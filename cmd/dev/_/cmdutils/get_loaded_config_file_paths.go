package cmdutils

import "github.com/zephinzer/dev/internal/config"

// GetLoadedConfigFilePaths returns a slice of strings where each string is an
// absolute path to a loaded configuration file
func GetLoadedConfigFilePaths() []string {
	output := []string{}
	for key := range config.Loaded {
		output = append(output, key)
	}
	return output
}
