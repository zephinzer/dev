package pivotaltracker

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/sanity-io/litter"
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
		query.Add("created_after", since[0].Format("2006-01-02T15:04:05Z"))
		query.Add("updated_after", since[0].Format("2006-01-02T15:04:05Z"))
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
		litter.Dump(string(responseBody))
		return nil, unmarshalError
	}
	return &response, nil
}
