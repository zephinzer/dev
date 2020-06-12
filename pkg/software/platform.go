package software

import "strings"

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
