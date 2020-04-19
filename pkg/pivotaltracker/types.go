package pivotaltracker

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

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

// APIv5MeResponse defines the response structure for a request made to
// the endpoint at https://www.pivotaltracker.com/services/v5/me
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

// APIv5NotificationsResponse defines the response structure for a request made to
// the endpoint at https://www.pivotaltracker.com/services/v5/my/notifications
type APIv5NotificationsResponse []APINotification

// String converts the notifications object into a CLI-friendly block of text
func (n APIv5NotificationsResponse) String() string {
	var output strings.Builder
	for i := 0; i < len(n); i++ {
		output.WriteString(n[i].String())
		output.Write([]byte{'\n', '\n'})
	}
	return output.String()
}

// APIv5StoriesResponse defines the reponse structure for a request made to the
// endpoint at https://www.pivotaltracker.com/services/v5/projects/{project_id}/stories
type APIv5StoriesResponse []APIStory

// String converts the notifications object into a CLI-friendly block of text
func (s APIv5StoriesResponse) String() string {
	var output strings.Builder
	for i := 0; i < len(s); i++ {
		output.WriteString(s[i].String())
		output.Write([]byte{'\n', '\n'})
	}
	return output.String()
}

// APIStory stores data about a pivotal tracker story as returned by its API
type APIStory struct {
	Kind          string     `json:"kind"`
	ID            int        `json:"id"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	StoryType     string     `json:"story_type"`
	Name          string     `json:"name"`
	Description   string     `json:"description,omitempty"`
	CurrentState  string     `json:"current_state"`
	RequestedByID int        `json:"requested_by_id"`
	URL           string     `json:"url"`
	ProjectID     int        `json:"project_id"`
	OwnerIds      []int      `json:"owner_ids"`
	Labels        []APILabel `json:"labels"`
	OwnedByID     int        `json:"owned_by_id,omitempty"`
	Estimate      int        `json:"estimate,omitempty"`
}

// String converts the story object into a CLI-friendly block of text
func (s APIStory) String() string {
	tag := s.StoryType
	message := s.Name
	link := s.URL
	state := s.CurrentState
	datetime := s.UpdatedAt
	timestamp, err := time.Parse("2006-01-02T15:04:05Z", datetime)
	if err == nil {
		datetime = humanize.Time(timestamp)
	}
	return fmt.Sprintf("[%s] %s : %s\n- link: %s\n- when: %s", tag, state, message, link, datetime)
}

type APILabel struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// APIAccount defines the structure for an Account object in responses to API queries
type APIAccount struct {
	Kind   string `json:"kind"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Plan   string `json:"plan"`
}

// APINotification defines the structure for a Notification object in responses to API queries
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

// String converts the notification object into a CLI-friendly text
func (n APINotification) String() string {
	tag := n.Project.Name
	message := n.Message
	referenceType := n.Story.Kind
	referenceLabel := n.Story.Name
	referenceID := n.Story.ID
	datetime := n.UpdatedAt
	timestamp, err := time.Parse("2006-01-02T15:04:05Z", datetime)
	if err == nil {
		datetime = humanize.Time(timestamp)
	}
	return fmt.Sprintf("[%s] %s\n- %s: %s\n- link: https://www.pivotaltracker.com/story/show/%v\n- when: about %s", tag, message, referenceType, referenceLabel, referenceID, datetime)

}

// APINotificationReference defines the structure for an object reference in responses to API queries
type APINotificationReference struct {
	Kind string `json:"kind"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// APIProject defines the structure for a project object in a response to API queries
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
