package pivotaltracker

import (
	"fmt"

	"github.com/usvc/dev/internal/db"
)

const (
	TableName = "pivotaltracker"
)

var SQLite3Migrations = []string{
	fmt.Sprintf("ALTER TABLE `%s` ADD `notification_id` VARCHAR(32) NOT NULL DEFAULT ''", TableName),
	fmt.Sprintf("ALTER TABLE `%s` ADD `message` TEXT NOT NULL DEFAULT ''", TableName),
	fmt.Sprintf("ALTER TABLE `%s` ADD `raw` TEXT NOT NULL DEFAULT ''", TableName),
}

func InitSQLite3Database(databasePath string) error {
	return db.ApplyMigrations(TableName, SQLite3Migrations, databasePath)
}
