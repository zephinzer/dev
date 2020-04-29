package trello

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/sanity-io/litter"
	"github.com/usvc/dev/pkg/utils"
)

func GetAccount(accessKey, accessToken string) (*APIv1MeResponse, error) {
	targetURL, urlParseError := url.Parse("https://api.trello.com/1/members/me")
	if urlParseError != nil {
		return nil, urlParseError
	}
	query := targetURL.Query()
	query.Add("key", accessKey)
	query.Add("token", accessToken)
	targetURL.RawQuery = query.Encode()
	responseObject, requestError := utils.HTTPGet(*targetURL, map[string]string{})
	if requestError != nil {
		return nil, requestError
	}
	defer responseObject.Body.Close()
	responseBody, bodyReadError := ioutil.ReadAll(responseObject.Body)
	if bodyReadError != nil {
		return nil, bodyReadError
	}
	var response APIv1MeResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		litter.Dump(string(responseBody))
		return nil, unmarshalError
	}
	return &response, nil
}
