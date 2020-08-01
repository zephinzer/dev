package pivotaltracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/pkg/utils/request"
)

// GetNotifs returns a user's current notifications
func GetNotifs(accessToken string, since ...time.Time) (*APIv5NotificationsResponse, error) {
	dateSinceFilter := time.Now().Add(-time.Hour * 24 * 365)
	if len(since) > 0 {
		dateSinceFilter = since[0]
	}
	responseObject, requestError := request.Get(request.GetOptions{
		URL: "https://www.pivotaltracker.com/services/v5/my/notifications",
		Headers: map[string]string{
			"Content-Type":   "application/json",
			"X-TrackerToken": accessToken,
		},
		Queries: map[string]string{
			"notification_types": ":all",
			"updated_after":      dateSinceFilter.Format(constants.PivotalTrackerAPITimeFormat),
		},
	})
	if requestError != nil {
		return nil, requestError
	}
	defer responseObject.Body.Close()
	responseBody, bodyReadError := ioutil.ReadAll(responseObject.Body)
	if bodyReadError != nil {
		return nil, bodyReadError
	}
	var response APIv5NotificationsResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %s\n\noriginal text: %s", unmarshalError, string(responseBody))
	}
	return &response, nil
}

// APIv5NotificationsResponse defines the response structure for a request made to
// the endpoint at https://www.pivotaltracker.com/services/v5/my/notifications
type APIv5NotificationsResponse []APINotification

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

// APINotificationReference defines the structure for an object reference in responses to API queries
type APINotificationReference struct {
	Kind string `json:"kind"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}
