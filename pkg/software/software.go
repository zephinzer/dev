package software

// Software represents a software that should be installed
// on the user's machine
type Software struct {
	Name    string  `json:"name" yaml:"name"`
	Check   Check   `json:"check" yaml:"check"`
	Install Install `json:"install" yaml:"install"`
}

type Install struct {
	Link string `json:"link" yaml:"link"`
}
