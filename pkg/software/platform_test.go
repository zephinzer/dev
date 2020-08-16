package software

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PlatformTests struct {
	suite.Suite
}

func TestPlatform(t *testing.T) {
	suite.Run(t, &PlatformTests{})
}

func (s PlatformTests) Test_Platform_String() {
	windowsFields := []string{
		"bill",
		"gates",
		"win",
		PlatformWindows,
	}
	for _, field := range windowsFields {
		s.Equal(PlatformWindows, Platform(field).String())
	}
	macosFields := []string{
		"steve",
		"jobs",
		"mac",
		"macos",
		PlatformMac,
	}
	for _, field := range macosFields {
		s.Equal(PlatformMac, Platform(field).String())
	}
	linuxFields := []string{
		"linus",
		"torvalds",
		"arch",
		"centos",
		"fedora",
		"rhel",
		"ubuntu",
		PlatformLinux,
	}
	for _, field := range linuxFields {
		s.Equal(PlatformLinux, Platform(field).String())
	}
	s.Equal("unknown", Platform("x").String())
}
