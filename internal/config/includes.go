package config

import (
	"fmt"

	"github.com/zephinzer/dev/pkg/utils"
)

// Includes is a list of paths to include
type Includes []IncludePath

// MergeWith adds non-duplicated strings from `o` into this instace
// of Includes
func (i *Includes) MergeWith(o Includes) {
	seen := map[string]bool{}
	for _, include := range *i {
		seen[string(include)] = true
	}
	for _, newInclude := range o {
		if val, ok := seen[string(newInclude)]; val && ok {
			continue
		}
		*i = append(*i, newInclude)
		seen[string(newInclude)] = true
	}
}

// IncludePath represents a file path to a configuration file
type IncludePath string

// GetConfig returns the configuration read from the path specified
// by this instance of IncludePath
func (i IncludePath) GetConfig() (*Config, error) {
	fullPath, resolvePathError := utils.ResolvePath(string(i))
	if resolvePathError != nil {
		return nil, fmt.Errorf("failed to resolve path '%s': %s", i, resolvePathError)
	}
	return NewFromFile(fullPath)
}
