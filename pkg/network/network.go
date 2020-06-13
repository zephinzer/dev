package network

// Network defines the configuration structure for the `network` property
// in the configuration file
type Network struct {
	Name            string `json:"name" yaml:"name,omitempty"`
	RegistrationURL string `json:"registrationUrl" yaml:"registrationUrl,omitempty"`
	Check           Check  `json:"check" yaml:"check,omitempty"`
}
