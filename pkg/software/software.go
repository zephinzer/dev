package software

// Software represents a software that should be installed
// on the user's machine
type Software struct {
	Name  string `json:"name" yaml:"name"`
	Check Check  `json:"check" yaml:"check"`
}
