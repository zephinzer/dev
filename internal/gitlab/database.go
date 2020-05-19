package gitlab

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zephinzer/dev/internal/db"
	pkggitlab "github.com/zephinzer/dev/pkg/gitlab"
)

const (
	TableName = "gl_notifs"
)

var SQLite3Migrations = []string{
	fmt.Sprintf("ALTER TABLE `%s` ADD `notification_id` VARCHAR(32) NOT NULL DEFAULT ''", TableName),
	fmt.Sprintf("ALTER TABLE `%s` ADD `message` TEXT NOT NULL DEFAULT ''", TableName),
	fmt.Sprintf("ALTER TABLE `%s` ADD `raw` TEXT NOT NULL DEFAULT ''", TableName),
	fmt.Sprintf("ALTER TABLE `%s` ADD `hostname` VARCHAR(256) NOT NULL DEFAULT ''", TableName),
}

func InitSQLite3Database(databasePath string) error {
	return db.ApplyMigrations(TableName, SQLite3Migrations, databasePath)
}

func InsertNotification(todo pkggitlab.APIv4Todo, hostname string, connection *sql.DB) error {
	notification := TodoSerializer(todo)
	notificationID := strconv.Itoa(todo.ID)
	notificationMessage := fmt.Sprintf("%s: %s", notification.GetTitle(), notification.GetMessage())
	notificationRaw, marshalError := json.Marshal(todo)
	if marshalError != nil {
		return marshalError
	}
	_, dbExecError := connection.Exec(
		fmt.Sprintf("INSERT INTO %s (notification_id, hostname, message, raw) VALUES (?, ?, ?, ?)", TableName),
		notificationID,
		hostname,
		notificationMessage,
		string(notificationRaw),
	)
	if dbExecError != nil {
		return dbExecError
	}
	return nil
}

func QueryNotification(todo pkggitlab.APIv4Todo, hostname string, connection *sql.DB) (bool, error) {
	notificationID := strconv.Itoa(todo.ID)
	row := connection.QueryRow(
		fmt.Sprintf("SELECT notification_id FROM %s WHERE notification_id = ? AND hostname = ?", TableName),
		notificationID,
		hostname,
	)
	var remoteNotificationID string
	if scanError := row.Scan(&remoteNotificationID); scanError != nil && scanError != sql.ErrNoRows {
		return false, scanError
	}
	if remoteNotificationID == notificationID {
		return true, nil
	}
	return false, nil
}
