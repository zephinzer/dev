package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/zephinzer/dev/pkg/utils/request"
)

func GetTodos(client request.Doer, hostname, accessToken string, since ...time.Time) (*APIv4TodoResponse, error) {
	requestObject, createRequestError := request.Create(request.CreateOptions{
		URL: fmt.Sprintf("https://%s/api/v4/todos", hostname),
		Headers: map[string][]string{
			"Content-Type":  {"application/json"},
			"PRIVATE-TOKEN": {accessToken},
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
		return nil, fmt.Errorf("failed to process response body: %s", bodyReadError)
	}
	var response APIv4TodoResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %s (raw object: '%s')", unmarshalError, string(responseBody))
	}
	return &response, nil
}
