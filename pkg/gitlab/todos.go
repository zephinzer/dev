package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/zephinzer/dev/pkg/utils/request"
)

// APIv4TodoResponse defines the response structure for a call to the Gitlab
// API endpoint at https://GITLAB_HOSTNAME/api/v4/todos
type APIv4TodoResponse []APIv4Todo

type APIv4Todo struct {
	ID         int          `json:"id"`
	Project    APIv4Project `json:"project"`
	Author     APIv4Author  `json:"author"`
	ActionName string       `json:"action_name"`
	TargetType string       `json:"target_type"`
	Target     APIv4Target  `json:"target"`
	TargetURL  string       `json:"target_url"`
	Body       string       `json:"body"`
	State      string       `json:"state"`
	CreatedAt  time.Time    `json:"created_at"`
}

func GetTodos(hostname, accessToken string, since ...time.Time) (*APIv4TodoResponse, error) {
	responseObject, requestError := request.Get(request.GetOptions{
		URL: fmt.Sprintf("https://%s/api/v4/todos", hostname),
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"PRIVATE-TOKEN": accessToken,
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
	var response APIv4TodoResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}
