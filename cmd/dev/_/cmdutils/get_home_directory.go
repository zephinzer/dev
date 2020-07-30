package cmdutils

import (
	"os"

	"github.com/zephinzer/dev/internal/log"
)

func GetHomeDirectory() string {
	homeDir, getHomeDirError := os.UserHomeDir()
	if getHomeDirError != nil {
		log.Errorf("failed to retrieve user's home directory: %s", getHomeDirError)
		os.Exit(1)
	}
	return homeDir
}
