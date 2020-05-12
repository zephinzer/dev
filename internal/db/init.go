package db

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Init initialises a local disk sqlite3 database for storage of incoming data
func Init(atPath string) error {
	directory := filepath.Dir(atPath)
	if !path.IsAbs(atPath) {
		cwd, getWorkingDirectoryError := os.Getwd()
		if getWorkingDirectoryError != nil {
			return getWorkingDirectoryError
		}
		directory = path.Join(cwd, directory)
	}
	dirInfo, checkDirectoryError := os.Lstat(directory)
	if checkDirectoryError == nil {
		if !dirInfo.IsDir() {
			return fmt.Errorf("path %s appears to already exist as a file", directory)
		}
	} else if checkDirectoryError != os.ErrNotExist {
		return fmt.Errorf("failed to run checks on path %s: %s", directory, checkDirectoryError)
	}
	if makeDirectoryError := os.MkdirAll(directory, os.ModePerm); makeDirectoryError != nil {
		return fmt.Errorf("failed to create directories leading up to %s: %s", directory, makeDirectoryError)
	}

	filename := filepath.Base(atPath)
	fullPath := path.Join(directory, filename)
	fullPathInfo, checkFileError := os.Lstat(fullPath)
	if checkFileError == nil {
		if fullPathInfo.IsDir() {
			return fmt.Errorf("path %s appears to already exist as a directory", fullPath)
		}
		return fmt.Errorf("path %s appears to already exist as a file", fullPath)
	} else if !os.IsNotExist(checkFileError) {
		return fmt.Errorf("error occurred during file checking: %s", checkFileError)
	}

	databaseFile, createFileError := os.Create(fullPath)
	if createFileError != nil {
		return fmt.Errorf("failed to create a file at %s: %s", fullPath, createFileError)
	}
	defer databaseFile.Close()
	openedDB, openDBError := NewConnection(fullPath)
	if openDBError != nil {
		return fmt.Errorf("failed to open sqlite3 database at %s: %s", fullPath, openDBError)
	}
	defer openedDB.Close()
	initQuery := "CREATE TABLE `devents` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `message` TEXT, `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"
	_, execError := openedDB.Exec(initQuery)
	if execError != nil {
		return fmt.Errorf("failed to execute initialisation query '%s': %s", initQuery, execError)
	}
	genesisQuery := "INSERT INTO `devents` (`message`) VALUES('initialised');"
	_, execError = openedDB.Exec(genesisQuery)
	if execError != nil {
		return fmt.Errorf("failed to execute insert first `dev` event with query '%s': %s", genesisQuery, execError)
	}
	if closeDBError := openedDB.Close(); closeDBError != nil {
		return fmt.Errorf("failed to close the database connection: %s", closeDBError)
	}
	if closeFileError := databaseFile.Close(); closeFileError != nil {
		return fmt.Errorf("failed to release the file handler: %s", closeFileError)
	}
	return nil
}

// InitTable creates 2 tables, one named `tableName` and the other named
// `tableName`_migrations using the provided `connection`; this way of doing things
// distributes the migrations so that each table is independently migratable
// and hence independently removable
func InitTable(tableName string, connection *sql.DB) error {
	var execError error
	_, execError = connection.Exec(fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS `%s_migrations` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `index` INTEGER NOT NULL, `script` TEXT NOT NULL, `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP);",
		tableName,
	))
	if execError != nil {
		return fmt.Errorf("failed to create `%s`'s migrations table: %s", tableName, execError)
	}
	_, execError = connection.Exec(fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS `%s` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP);",
		tableName,
	))
	if execError != nil {
		return fmt.Errorf("failed to create `%s`'s logical table: %s", tableName, execError)
	}
	return nil
}
