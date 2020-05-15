package pivotaltracker

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zephinzer/dev/internal/db"
	pkgpivotaltracker "github.com/zephinzer/dev/pkg/pivotaltracker"
)

const (
	TableName = "pt_notifs"
)

var SQLite3Migrations = []string{
	fmt.Sprintf("ALTER TABLE `%s` ADD `notification_id` VARCHAR(32) NOT NULL DEFAULT ''", TableName),
	fmt.Sprintf("ALTER TABLE `%s` ADD `message` TEXT NOT NULL DEFAULT ''", TableName),
	fmt.Sprintf("ALTER TABLE `%s` ADD `raw` TEXT NOT NULL DEFAULT ''", TableName),
}

func InitSQLite3Database(databasePath string) error {
	return db.ApplyMigrations(TableName, SQLite3Migrations, databasePath)
}

func InsertNotification(notification pkgpivotaltracker.APINotification, connection *sql.DB) error {
	notificationID := strconv.Itoa(notification.ID)
	notificationMessage := notification.String()
	notificationRaw, marshalError := json.Marshal(notification)
	if marshalError != nil {
		return marshalError
	}
	_, dbExecError := connection.Exec("INSERT INTO pt_notifs (notification_id, message, raw) VALUES (?, ?, ?)", notificationID, notificationMessage, string(notificationRaw))
	if dbExecError != nil {
		return dbExecError
	}
	return nil
}

func QueryNotification(notification pkgpivotaltracker.APINotification, connection *sql.DB) (bool, error) {
	notificationID := strconv.Itoa(notification.ID)
	row := connection.QueryRow("SELECT notification_id FROM pt_notifs WHERE notification_id = ?", notificationID)
	var remoteNotificationID string
	if scanError := row.Scan(&remoteNotificationID); scanError != nil && scanError != sql.ErrNoRows {
		return false, scanError
	}
	if remoteNotificationID == notificationID {
		return true, nil
	}
	return false, nil
}
