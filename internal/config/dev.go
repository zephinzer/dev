package config

// Dev specifies the configurations available for the CLI tool itself
type Dev struct {
	Client   DevClient   `json:"client" yaml:"client"`
	Defaults DevDefaults `json:"defaults" yaml:"defaults"`
}

// DevClient holds configurations related to the CLI tool
type DevClient struct {
	Database      DevClientDatabase      `json:"database" yaml:"database"`
	Notifications DevClientNotifications `json:"notifications" yaml:"notifications"`
}

// DevClientDatabase holds configurations related to the data persistence
// mechanism of the CLI tool
type DevClientDatabase struct {
	Path string `json:"path" yaml:"path"`
}

// DevClientNotifications holds configurations related to the notifications
// mechanisms of the CLI tool
type DevClientNotifications struct {
	Telegram DevClientNotificationsTelegram `json:"telegram" yaml:"telegram"`
}

// DevClientNotificationsTelegram holds configurations related to the
// telegram integration for sending notifications
type DevClientNotificationsTelegram struct {
	Token string `json:"token" yaml:"token"`
	ID    string `json:"id" yaml:"id"`
}

// DevDefaults specifies default selections when running `dev <verb> <noum> <this>` commands
type DevDefaults struct {
	// Account defines the default platform to use for account
	Account string `json:"account" yaml:"account"`
	// Notifications defines the default platform to use for notifications
	Notifications string `json:"notifications" yaml:"notifications"`
	// Work defines the default platform to use for work
	Work string `json:"work" yaml:"work"`
}