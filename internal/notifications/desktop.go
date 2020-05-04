package notifications

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func TriggerDesktop(title, message string, icon ...string) error {
	notificationIcon := ""
	if len(icon) > 0 {
		notificationIcon = icon[0]
	}
	if err := beeep.Notify(title, message, notificationIcon); err != nil {
		return fmt.Errorf("failed to trigger notification: %s", err)
	}
	return nil
}
