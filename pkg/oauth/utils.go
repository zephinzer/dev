package oauth

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

// Open runs xdg-open on linux, open on macos, and rundll32.exe on windows
// with the provided :targetURI as an argument
func Open(targetURI string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", targetURI).Start()
	case "macos":
		return exec.Command("open", targetURI).Start()
	case "windows":
		return exec.Command(path.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe"), "url.dll,FileProtocolHandler", targetURI).Start()
	}
	return fmt.Errorf("unsupported platform '%s'", runtime.GOOS)
}
