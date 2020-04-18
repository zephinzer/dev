package software

// Software represents a software that should be installed
// on the user's machine
type Software struct {
	Name  string `yaml:"name"`
	Check Check  `yaml:"check"`
}

type Check struct {
	// Command is the command to run including arguments
	Command []string `yaml:"command"`
	// ExitCode is the expected exit code we can expect from the command running if defined
	ExitCode int `yaml:"exitCode"`
	// Stdout is the expected output on stdout if defined
	Stdout string `yaml:"stdout"`
	// Stderr is the expected output on stderr if defined
	Stderr string `yaml:"stderr"`
}
