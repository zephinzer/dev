package utils

import (
	"fmt"
	"path"
	"os"
	"os/exec"
	"runtime"
)

// OpenURIWithDefaultApplication runs xdg-open on linux, open on macos, and rundll32.exe on windows
// with the provided :targetURI as an argument
func OpenURIWithDefaultApplication(targetURI string) error {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", targetURI).Start()
	case "macos":
		exec.Command("open", targetURI).Start()
	case "windows":
		exec.Command(path.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe"), "url.dll,FileProtocolHandler", targetURI).Start()
	default:
		return fmt.Errorf("unsupported platform '%s'", runtime.GOOS)
	}
	return nil
}
