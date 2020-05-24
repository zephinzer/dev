package config

// Includes is a list of paths to include
type Includes []IncludePath

// IncludePath represents a file path to a configuration file
type IncludePath string

// GetConfig returns the configuration read from the path specified
// by this instance of IncludePath
func (i IncludePath) GetConfig() (*Config, error) {
	return NewFromFile(string(i))
}
