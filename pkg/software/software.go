package software

// Software represents a software that should be installed on the user's machine
type Software struct {
	Name        string `json:"name" yaml:"name,omitempty"`
	Description string `json:"description" yaml:"description,omitempty"`
	// Platforms when specified, restricts the operations to run only
	// when the check is being run on the allowed operating systems
	Platforms Platforms `json:"platforms" yaml:"platforms,omitempty"`
	Check     Check     `json:"check" yaml:"check,omitempty"`
	Install   Install   `json:"install" yaml:"install,omitempty"`
}

// Install defines installation instructions for the software
type Install struct {
	Link string `json:"link" yaml:"link"`
}
