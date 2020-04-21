package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/sanity-io/litter"
	"github.com/usvc/dev/internal/constants"
	"github.com/usvc/dev/pkg/utils"
)

func GetAccount(hostname, accessToken string) (*APIv4UserResponse, error) {
	gitlabHostname := hostname
	if len(hostname) == 0 {
		gitlabHostname = constants.DefaultGitlabHostname
	}
	targetURL, urlParseError := url.Parse(fmt.Sprintf("https://%s/api/v4/user", gitlabHostname))
	if urlParseError != nil {
		return nil, urlParseError
	}
	responseObject, requestError := utils.HTTPGet(*targetURL, map[string]string{
		"Content-Type":  "application/json",
		"PRIVATE-TOKEN": accessToken,
	})
	if requestError != nil {
		return nil, requestError
	}
	defer responseObject.Body.Close()
	responseBody, bodyReadError := ioutil.ReadAll(responseObject.Body)
	if bodyReadError != nil {
		return nil, bodyReadError
	}
	var response APIv4UserResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		litter.Dump(string(responseBody))
		return nil, unmarshalError
	}
	return &response, nil
}
