package pivotaltracker

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/zephinzer/dev/pkg/utils"
)

func GetAccount(accessToken string) (*APIv5MeResponse, error) {
	targetURL, urlParseError := url.Parse("https://www.pivotaltracker.com/services/v5/me")
	if urlParseError != nil {
		return nil, urlParseError
	}
	query := targetURL.Query()
	query.Add("fields", ":default")
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
	var response APIv5MeResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}
