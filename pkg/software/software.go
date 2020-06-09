package software

import "strings"

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

// Platforms is a list of platforms
type Platforms []Platform

const (
	// PlatformLinux is a universal string for linux-based operating systems
	// based on the possible values of runtime.GOOS
	PlatformLinux = "linux"
	// PlatformMac is a universal string for macos-based operating systems
	// based on the possible values of runtime.GOOS
	PlatformMac = "darwin"
	// PlatformWindows is a universal string for windows-based operating systems
	// based on the possible values of runtime.GOOS
	PlatformWindows = "windows"
)

// Platform should define an operating system
type Platform string

func (p Platform) String() string {
	switch strings.ToLower(string(p)) {
	case "win":
		fallthrough
	case "windows":
		return PlatformWindows
	case "mac":
		fallthrough
	case "macos":
		fallthrough
	case "darwin":
		return PlatformMac
	case "arch":
		fallthrough
	case "centos":
		fallthrough
	case "fedora":
		fallthrough
	case "rhel":
		fallthrough
	case "ubuntu":
		return PlatformLinux
	}
	return "unknown"
}
