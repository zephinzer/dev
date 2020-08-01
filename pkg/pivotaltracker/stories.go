package pivotaltracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/zephinzer/dev/internal/constants"
	"github.com/zephinzer/dev/pkg/utils/request"
)

// GetStories returns a user's stories
func GetStories(accessToken string, inProjectID string, since ...time.Time) (*APIv5StoriesResponse, error) {
	accountInfo, accountInfoError := GetAccount(accessToken)
	if accountInfoError != nil {
		return nil, accountInfoError
	}
	dateSinceFilter := time.Now().Add(-time.Hour * 24 * 365)
	if len(since) > 0 {
		dateSinceFilter = since[0]
	}
	responseObject, requestError := request.Get(request.GetOptions{
		URL: fmt.Sprintf(
			"https://www.pivotaltracker.com/services/v5/projects/%s/stories",
			inProjectID,
		),
		Headers: map[string]string{
			"Content-Type":   "application/json",
			"X-TrackerToken": accessToken,
		},
		Queries: map[string]string{
			"filter": fmt.Sprintf(
				"(mywork:%s OR is:following) AND -state:accepted AND -state:planned AND updated_after:\"%s\"",
				accountInfo.Username,
				dateSinceFilter.Format(constants.PivotalTrackerAPITimeFormat),
			),
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
	var response APIv5StoriesResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}

// APIv5StoriesResponse defines the reponse structure for a request made to the
// endpoint at https://www.pivotaltracker.com/services/v5/projects/{project_id}/stories
type APIv5StoriesResponse []APIStory

// String converts the notifications object into a CLI-friendly block of text
func (s APIv5StoriesResponse) String(format ...string) string {
	var output strings.Builder
	for i := 0; i < len(s); i++ {
		format = append(format, "")
		switch format[0] {
		case "md":
			fallthrough
		case "markdown":
			output.WriteString(fmt.Sprintf("%v. %s", i+1, s[i].String(format...)))
			output.Write([]byte{'\n'})
		default:
			output.WriteString(s[i].String(format...))
			output.Write([]byte{'\n', '\n'})
		}
	}
	return output.String()
}

// APIStory stores data about a pivotal tracker story as returned by its API
type APIStory struct {
	Kind          string     `json:"kind"`
	ID            int        `json:"id"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	StoryType     string     `json:"story_type"`
	Name          string     `json:"name"`
	Description   string     `json:"description,omitempty"`
	CurrentState  string     `json:"current_state"`
	RequestedByID int        `json:"requested_by_id"`
	URL           string     `json:"url"`
	ProjectID     int        `json:"project_id"`
	OwnerIds      []int      `json:"owner_ids"`
	Labels        []APILabel `json:"labels"`
	OwnedByID     int        `json:"owned_by_id,omitempty"`
	Estimate      int        `json:"estimate,omitempty"`
}

// String converts the story object into a CLI-friendly block of text
func (s APIStory) String(format ...string) string {
	tag := s.StoryType
	tagIcon := "📌"
	switch tag {
	case "feature":
		tagIcon = "🌟"
	case "chore":
		tagIcon = "⚙️"
	case "bug":
		tagIcon = "🐞"
	}
	message := s.Name
	link := s.URL
	state := s.CurrentState
	datetime := s.UpdatedAt
	timestamp, err := time.Parse(constants.PivotalTrackerAPITimeFormat, datetime)
	if err == nil {
		datetime = humanize.Time(timestamp)
	}
	format = append(format, "")
	switch format[0] {
	case "md":
		fallthrough
	case "markdown":
		return fmt.Sprintf("[%s (%s was %s %s)](%s)", message, tag, state, datetime, link)
	default:
		return fmt.Sprintf("%s %s (link: %s)\n- %s is in %s state and last updated %s", tagIcon, message, link, tag, state, datetime)
	}
}
