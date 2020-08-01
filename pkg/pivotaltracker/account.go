package pivotaltracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/zephinzer/dev/pkg/utils/request"
)

func GetAccount(accessToken string) (*APIv5AccountResponse, error) {
	responseObject, requestError := request.Get(request.GetOptions{
		URL: "https://www.pivotaltracker.com/services/v5/me",
		Headers: map[string]string{
			"Content-Type":   "application/json",
			"X-TrackerToken": accessToken,
		},
		Queries: map[string]string{
			"fields": ":default",
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
	var response APIv5AccountResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}

// APIv5AccountResponse defines the response structure for a request made to
// the endpoint at https://www.pivotaltracker.com/services/v5/me
type APIv5AccountResponse struct {
	Accounts                   []APIAccount `json:"accounts"`
	APIToken                   string       `json:"api_token"`
	CreatedAt                  string       `json:"created_at"`
	Email                      string       `json:"email"`
	HasGoogleIdentity          bool         `json:"has_google_identity"`
	ID                         int          `json:"id"`
	Initials                   string       `json:"initials"`
	Kind                       string       `json:"kind"`
	Name                       string       `json:"name"`
	Projects                   []APIProject `json:"projects"`
	ReceivesInAppNotifications bool         `json:"receives_in_app_notifications"`
	TimeZone                   APITimezone  `json:"time_zone"`
	UpdatedAt                  string       `json:"updated_at"`
	Username                   string       `json:"username"`
}

func (m APIv5AccountResponse) String(format ...string) string {
	var me strings.Builder
	// provide a default in case there is none
	format = append(format, "")
	switch format[0] {
	case "md":
		fallthrough
	case "markdown":
		me.WriteString("## pivotal tracker account information\n\n")
		me.WriteString("| field | value |\n")
		me.WriteString("| --- | --- |\n")
		me.WriteString(fmt.Sprintf("| account id | %v |\n", m.ID))
		me.WriteString(fmt.Sprintf("| username   | %s (%s) |\n", m.Username, m.Initials))
		me.WriteString(fmt.Sprintf("| real name  | %s |\n", m.Name))
		me.WriteString(fmt.Sprintf("| projects   | %v |\n\n", len(m.Projects)))
		me.WriteString("## pivotal tracker projects\n\n")
		for index, project := range m.Projects {
			me.WriteString(fmt.Sprintf("%v. %s in [%s (id: %v)](https://www.pivotaltracker.com/n/projects/%v)\n", index+1, project.Role, project.ProjectName, project.ProjectID, project.ProjectID))
		}
		return me.String()
	default:
		me.WriteString("pivotal tracker account information\n")
		me.WriteString(fmt.Sprintf("account id : %v\n", m.ID))
		me.WriteString(fmt.Sprintf("username   : %s (%s)\n", m.Username, m.Initials))
		me.WriteString(fmt.Sprintf("real name  : %s\n", m.Name))
		me.WriteString(fmt.Sprintf("projects   : %v\n", len(m.Projects)))
		for _, project := range m.Projects {
			me.WriteString(fmt.Sprintf("  - %s in %s (id: %v) - https://www.pivotaltracker.com/n/projects/%v\n", project.Role, project.ProjectName, project.ProjectID, project.ProjectID))
		}
		return me.String()
	}
}
