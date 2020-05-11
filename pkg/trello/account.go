package trello

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"github.com/usvc/dev/pkg/utils"
)

func GetAccount(accessKey, accessToken string) (*APIv1MemberResponse, error) {
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
	var response APIv1MemberResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}

type APIv1MemberResponse struct {
	ID              string   `json:"id" yaml:"id"`
	Username        string   `json:"username" yaml:"username"`
	FullName        string   `json:"fullName" yaml:"fullName"`
	Initials        string   `json:"initials" yaml:"initials"`
	URL             string   `json:"url" yaml:"url"`
	Email           string   `json:"email" yaml:"email"`
	BoardIDs        []string `json:"idBoards" yaml:"idBoards"`
	OrganizationIDs []string `json:"idOrganizations" yaml:"idOrganizations"`
	ActivityBlocked bool     `json:"activityBlocked" yaml:"activityBlocked"`
}
