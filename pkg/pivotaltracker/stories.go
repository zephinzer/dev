package pivotaltracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/sanity-io/litter"
	"github.com/usvc/dev/pkg/utils"
)

// GetStories returns a user's stories
func GetStories(accessToken string, inProjectID string, since ...time.Time) (*APIv5StoriesResponse, error) {
	accountInfo, accountInfoError := GetAccount(accessToken)
	if accountInfoError != nil {
		return nil, accountInfoError
	}
	targetURL, urlParseError := url.Parse(fmt.Sprintf(
		"https://www.pivotaltracker.com/services/v5/projects/%s/stories",
		inProjectID,
	))
	if urlParseError != nil {
		return nil, urlParseError
	}
	sinceThisTime := time.Now().Add(time.Duration(-1 * 31 * 24 * time.Hour))
	if len(since) > 0 {
		sinceThisTime = since[0]
	}
	query := targetURL.Query()
	query.Add("filter", fmt.Sprintf(
		"(mywork:%s OR is:following) AND -state:accepted AND -state:planned AND updated_after:\"%s\"",
		accountInfo.Username,
		sinceThisTime.Format("2006-01-02T15:04:05Z"),
	))
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
	var response APIv5StoriesResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		litter.Dump(string(responseBody))
		return nil, unmarshalError
	}
	return &response, nil
}
