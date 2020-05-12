package db

import "database/sql"

// NewConnection creates a new database connection to the sqlite3 database located
// at `databasePath`
func NewConnection(databasePath string) (*sql.DB, error) {
	openedDB, openDBError := sql.Open("sqlite3", databasePath)
	if openDBError != nil {
		return nil, openDBError
	}
	return openedDB, nil
}
