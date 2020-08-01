package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/zephinzer/dev/pkg/utils/request"
)

func GetAccount(accessToken string) (*APIv3UserResponse, error) {
	responseObject, requestError := request.Get(request.GetOptions{
		URL: "https://api.github.com/user",
		Headers: map[string]string{
			"Accept":        "application/vnd.github.v3+json", // as requested at https://developer.github.com/v3/#current-version
			"Authorization": fmt.Sprintf("token %s", accessToken),
		},
	})
	if requestError != nil {
		return nil, requestError
	}
	defer responseObject.Body.Close()
	responseBody, bodyReadError := ioutil.ReadAll(responseObject.Body)
	if bodyReadError != nil {
		return nil, bodyReadError
	}
	var response APIv3UserResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}
