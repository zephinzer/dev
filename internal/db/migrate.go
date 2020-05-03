package db

import (
	"database/sql"
	"fmt"

	"github.com/usvc/dev/internal/log"
)

func IsMigrationAppliedAndValid(tableName string, migrationIndex int, migrationScript string, connection *sql.DB) (bool, error) {
	row := connection.QueryRow(fmt.Sprintf("SELECT `index`, `script` FROM `%s_migrations` WHERE `index` = ? AND `script` = ?", tableName), migrationIndex, migrationScript)
	var storedIndex int
	var storedScript string
	var scanError error
	if scanError = row.Scan(&storedIndex, &storedScript); scanError != nil {
		if scanError == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to retrieve migration from `%s`'s migrations table: %s", tableName, scanError)
	}
	if storedIndex != migrationIndex || storedScript != migrationScript {
		return true, fmt.Errorf("migration '%s' at index %v does not seem equivalent to stored migration '%s' seen last at index %v", migrationScript, migrationIndex, storedScript, storedIndex)
	}
	return true, nil
}

func ApplyMigrations(tableName string, migrations []string, databasePath string) error {
	connection, newConnectionError := NewConnection(databasePath)
	defer connection.Close()
	if newConnectionError != nil {
		return newConnectionError
	}
	if initError := InitTable(tableName, connection); initError != nil {
		return initError
	}
	for index, script := range migrations {
		log.Debugf("checking and applying migration[%v] (script: '%s')...\n", index, script)
		applied, validityError := IsMigrationAppliedAndValid(tableName, index, script, connection)
		if validityError != nil {
			return fmt.Errorf("encountered validation error at migration[%v] (script: '%s'): %s", index, script, validityError)
		} else if applied {
			log.Tracef("migration[%v] already applied\n", index)
		} else {
			if migrationError := RunMigration(tableName, index, script, connection); migrationError != nil {
				return fmt.Errorf("encountered migration error at migration[%v] (script: '%s'): %s", index, script, migrationError)
			}
			log.Infof("migration[%v] successfully applied\n", index)
		}
	}
	return nil
}

func RunMigration(tableName string, migrationIndex int, migrationScript string, connection *sql.DB) error {
	_, execError := connection.Exec(migrationScript)
	if execError != nil {
		return fmt.Errorf("failed to run migration: %s", execError)
	}
	_, execError = connection.Exec(fmt.Sprintf("INSERT INTO `%s_migrations` (`index`, `script`) VALUES(?, ?)", tableName), migrationIndex, migrationScript)
	if execError != nil {
		return fmt.Errorf("failed to update migration table: %s", execError)
	}
	return nil
}
