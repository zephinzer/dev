package system

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

// OpenURIWithDefaultApplication runs xdg-open on linux, open on macos, and rundll32.exe on windows
// with the provided :targetURI as an argument
func OpenURIWithDefaultApplication(targetURI string) error {
	commandString, err := getDefaultOpenURICommand(runtime.GOOS, targetURI)
	if err != nil {
		return err
	}
	return exec.Command(commandString[0], commandString[1:]...).Start()
}

func getDefaultOpenURICommand(goos, targetURI string) ([]string, error) {
	switch goos {
	case "linux":
		return []string{"xdg-open", targetURI}, nil
	case "macos":
		return []string{"open", targetURI}, nil
	case "windows":
		return []string{path.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe"), "url.dll,FileProtocolHandler", targetURI}, nil
	}
	return nil, fmt.Errorf("unsupported platform '%s'", goos)
}
