package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/zephinzer/dev/pkg/utils"
)

func GetTodos(hostname, accessToken string, since ...time.Time) (*APIv4TodoResponse, error) {
	targetURL, urlParseError := url.Parse(fmt.Sprintf("https://%s/api/v4/todos", hostname))
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
	var response APIv4TodoResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}
