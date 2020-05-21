package pivotaltracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/pkg/utils"
)

// GetNotifs returns a user's current notifications
func GetNotifs(accessToken string, since ...time.Time) (*APIv5NotificationsResponse, error) {
	targetURL, urlParseError := url.Parse("https://www.pivotaltracker.com/services/v5/my/notifications")
	if urlParseError != nil {
		return nil, urlParseError
	}
	query := targetURL.Query()
	query.Add("notification_types", ":all")
	if len(since) > 0 {
		query.Add("updated_after", since[0].Format(constants.PivotalTrackerAPITimeFormat))
	}
	targetURL.RawQuery = query.Encode()
	responseObject, requestError := utils.HTTPGet(*targetURL, map[string]string{
		"Content-Type":   "application/json",
		"X-TrackerToken": accessToken,
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
