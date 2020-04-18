package pivotaltracker

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/usvc/dev/pkg/utils"
)

func GetNotifs(accessToken string) (*APIv5NotificationsResponse, error) {
	targetURL, urlParseError := url.Parse("https://www.pivotaltracker.com/services/v5/my/notifications")
	if urlParseError != nil {
		return nil, urlParseError
	}
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
