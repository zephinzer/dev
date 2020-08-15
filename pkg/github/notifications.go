package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/pkg/utils/request"
)

// GetNotifications retrieves notifications from the Github API using the
// provided access token
//
// Documentation at: https://developer.github.com/v3/activity/notifications/
// To run this as a cURL:
// `curl -H "Authorization: token XXX" https://api.github.com/notifications`
func GetNotifications(client request.Doer, accessToken string, since ...time.Time) (*APIv3Notifications, error) {
	dateSinceFilter := time.Now().Add(-time.Hour * 24 * 365)
	if len(since) > 0 {
		dateSinceFilter = since[0]
	}
	requestObject, createRequestError := request.Create(request.CreateOptions{
		URL: "https://api.github.com/notifications",
		Headers: map[string][]string{
			"Accept":        {"application/vnd.github.v3+json"}, // as requested at https://developer.github.com/v3/#current-version
			"Authorization": {fmt.Sprintf("token %s", accessToken)},
		},
		Queries: map[string][]string{
			"participating": {"true"},
			"since":         {dateSinceFilter.Format(constants.GithubAPITimeFormat)},
		},
	})
	if createRequestError != nil {
		return nil, fmt.Errorf("failed to create request: %s", createRequestError)
	}
	responseObject, doError := client.Do(requestObject)
	if doError != nil {
		return nil, fmt.Errorf("failed to execute request: %s", doError)
	}
	defer responseObject.Body.Close()
	responseBody, bodyReadError := ioutil.ReadAll(responseObject.Body)
	if bodyReadError != nil {
		return nil, bodyReadError
	}
	var response APIv3Notifications
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s (raw object: '%s')", unmarshalError, string(responseBody))
	}
	return &response, nil
}
