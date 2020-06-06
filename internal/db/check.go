package db

import (
	"fmt"
	"os"

	"github.com/zephinzer/dev/pkg/utils"
)

// Check performs a check whether the file at path `atPath` is already
// an initialised database
func Check(atPath string) error {
	absolutePath, resolvePathError := utils.ResolvePath(atPath)
	if resolvePathError != nil {
		return fmt.Errorf("failed to resolve path '%s': %s", atPath, resolvePathError)
	}
	fileInfo, checkFileError := os.Lstat(absolutePath)
	if checkFileError != nil {
		if !os.IsNotExist(checkFileError) {
			return fmt.Errorf("error occurred during file checking: %s", checkFileError)
		}
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("path %s appears to already exist as a directory", absolutePath)
	}

	openedDB, openDBError := NewConnection(absolutePath)
	if openDBError != nil {
		return fmt.Errorf("failed to open sqlite3 database at %s: %s", absolutePath, openDBError)
	}
	defer openedDB.Close()

	initialised := "initialised"
	genesisQuery := "SELECT message FROM `devents` WHERE `message` = ?"
	row := openedDB.QueryRow(genesisQuery, initialised)
	if scanError := row.Scan(&initialised); scanError != nil {
		return fmt.Errorf("failed to check existence of database: %s", scanError)
	}
	if initialised != "initialised" {
		return fmt.Errorf("table `devents` exists but returned a weird result '%s'", initialised)
	}
	return nil
}
