package software

// Software represents a software that should be installed
// on the user's machine
type Software struct {
	Name  string `json:"name" yaml:"name"`
	Check Check  `json:"check" yaml:"check"`
}

type Check struct {
	// Command is the command to run including arguments
	Command []string `json:"command" yaml:"command"`
	// ExitCode is the expected exit code we can expect from the command running if defined
	ExitCode int `json:"exitCode" yaml:"exitCode"`
	// Stdout is the expected output on stdout if defined
	Stdout string `json:"stdout" yaml:"stdout"`
	// Stderr is the expected output on stderr if defined
	Stderr string `json:"stderr" yaml:"stderr"`
}
