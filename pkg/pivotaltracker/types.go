package pivotaltracker

import "time"

// Config defines the structure of the configuration that can be consumed for tracking
// notifications and events from the Pivotal Tracker platform
type Config struct {
	// Projects is the list of projects to monitor
	Projects Projects `yaml:"projects"`
	// AccessToken to be used for all listed projects if their own access token is not defined.
	// This value can be retrieved from the section named API Token at the URL:
	//   https://www.pivotaltracker.com/profile
	AccessToken string `yaml:"accessToken"`
}

// Projects defines a structure to store a list of Pivotal Tracker projects
type Projects []Project

// Project defines a Pivotal Tracker project
type Project struct {
	// Name is the custom label given to this project as the user would like to define it
	Name string `yaml:"name"`
	// ProjectID of the project in PivotalTracker that's obtainable by going to a board and getting it from the URL
	ProjectID string `yaml:"projectID"`
	// AccessToken for this project, if not specified, defaults to use parent access token.
	// This value can be retrieved from the section named API Token at the URL:L
	//   https://www.pivotaltracker.com/profile
	AccessToken string `yaml:"accessToken"`
}

type APIv5MeResponse struct {
	Accounts                   []APIAccount `json:"accounts"`
	APIToken                   string       `json:"api_token"`
	CreatedAt                  time.Time    `json:"created_at"`
	Email                      string       `json:"email"`
	HasGoogleIdentity          bool         `json:"has_google_identity"`
	ID                         int          `json:"id"`
	Initials                   string       `json:"initials"`
	Kind                       string       `json:"kind"`
	Name                       string       `json:"name"`
	Projects                   []APIProject `json:"projects"`
	ReceivesInAppNotifications bool         `json:"receives_in_app_notifications"`
	TimeZone                   APITimezone  `json:"time_zone"`
	UpdatedAt                  time.Time    `json:"updated_at"`
	Username                   string       `json:"username"`
}

type APIv5NotificationsResponse []APINotification

type APIAccount struct {
	Kind   string `json:"kind"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Plan   string `json:"plan"`
}

type APINotification struct {
	Kind               string                    `json:"kind"`
	ID                 int                       `json:"id"`
	Project            APINotificationReference  `json:"project"`
	Performer          APINotificationReference  `json:"performer"`
	Message            string                    `json:"message"`
	NotificationType   string                    `json:"notification_type"`
	NewAttachmentCount int                       `json:"new_attachment_count,omitempty"`
	Action             string                    `json:"action"`
	Story              APINotificationReference  `json:"story,omitempty"`
	CreatedAt          string                    `json:"created_at"`
	UpdatedAt          string                    `json:"updated_at"`
	Epic               *APINotificationReference `json:"epic,omitempty"`
	CommentID          int                       `json:"comment_id,omitempty"`
	ReadAt             string                    `json:"read_at,omitempty"`
}

type APINotificationReference struct {
	Kind string `json:"kind"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type APIProject struct {
	Kind         string    `json:"kind"`
	ID           int       `json:"id"`
	ProjectID    int       `json:"project_id"`
	ProjectName  string    `json:"project_name"`
	ProjectColor string    `json:"project_color"`
	Favorite     bool      `json:"favorite"`
	Role         string    `json:"role"`
	LastViewedAt time.Time `json:"last_viewed_at"`
}

type APITimezone struct {
	Kind      string `json:"kind"`
	OlsonName string `json:"olson_name"`
	Offset    string `json:"offset"`
}
