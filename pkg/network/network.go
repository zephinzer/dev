package network

// Network defines the configuration structure for the `network` property
// in the configuration file
type Network struct {
	Name  string `json:"name" yaml:"name"`
	Check Check  `json:"check" yaml:"check"`
}
