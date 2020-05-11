package pivotaltracker

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/pkg/utils"
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
		query.Add("created_after", since[0].Format(constants.PivotalTrackerAPITimeFormat))
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
		return nil, unmarshalError
	}
	return &response, nil
}
