package gitlab

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

func GetAccount(hostname, accessToken string) (*APIv4UserResponse, error) {
	responseObject, requestError := request.Get(request.GetOptions{
		URL: fmt.Sprintf("https://%s/api/v4/user", hostname),
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
	var response APIv4UserResponse
	unmarshalError := json.Unmarshal(responseBody, &response)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &response, nil
}

// APIv4UserResponse defines the response structure for a call to the Gitlab
// API endpoint at https://GITLAB_HOSTNAME/api/v4/user
type APIv4UserResponse struct {
	ID               int             `json:"id"`
	Username         string          `json:"username"`
	Email            string          `json:"email"`
	Name             string          `json:"name"`
	State            string          `json:"state"`
	AvatarURL        string          `json:"avatar_url"`
	WebURL           string          `json:"web_url"`
	CreatedAt        string          `json:"created_at"`
	IsAdmin          bool            `json:"is_admin"`
	Bio              interface{}     `json:"bio"`
	Location         interface{}     `json:"location"`
	PublicEmail      string          `json:"public_email"`
	Skype            string          `json:"skype"`
	Linkedin         string          `json:"linkedin"`
	Twitter          string          `json:"twitter"`
	WebsiteURL       string          `json:"website_url"`
	Organization     string          `json:"organization"`
	JobTitle         string          `json:"job_title"`
	LastSignInAt     string          `json:"last_sign_in_at"`
	ConfirmedAt      string          `json:"confirmed_at"`
	ThemeID          int             `json:"theme_id"`
	LastActivityOn   string          `json:"last_activity_on"`
	ColorSchemeID    int             `json:"color_scheme_id"`
	ProjectsLimit    int             `json:"projects_limit"`
	CurrentSignInAt  string          `json:"current_sign_in_at"`
	Identities       APIv4Identities `json:"identities"`
	CanCreateGroup   bool            `json:"can_create_group"`
	CanCreateProject bool            `json:"can_create_project"`
	TwoFactorEnabled bool            `json:"two_factor_enabled"`
	External         bool            `json:"external"`
	PrivateProfile   bool            `json:"private_profile"`
	CurrentSignInIP  string          `json:"current_sign_in_ip"`
	LastSignInIP     string          `json:"last_sign_in_ip"`
}

func (u APIv4UserResponse) String() string {
	var output strings.Builder
	output.WriteString("gitlab account information\n")
	output.WriteString(fmt.Sprintf("username      : %s\n", u.Username))
	output.WriteString(fmt.Sprintf("real name     : %s\n", u.Name))
	if len(u.Email) == 0 {
		u.Email = "(hidden)"
	}
	output.WriteString(fmt.Sprintf("account email : %s\n", u.Email))
	if len(u.PublicEmail) == 0 {
		u.PublicEmail = "(hidden)"
	}
	output.WriteString(fmt.Sprintf("public email  : %s\n", u.PublicEmail))
	output.WriteString(fmt.Sprintf("2fa-enabled   : %v\n", u.TwoFactorEnabled))
	output.WriteString(fmt.Sprintf("link          : %s\n", u.WebURL))
	if createdAt, err := time.Parse(constants.GitlabAPITimeFormat, u.CreatedAt); err != nil {
		output.WriteString(fmt.Sprintf("created at    : %s\n", u.CreatedAt))
	} else {
		output.WriteString(fmt.Sprintf("created at    : %s (about %s)\n", u.CreatedAt, humanize.Time(createdAt)))
	}
	if lastActive, err := time.Parse(constants.DateOnlyTimeFormat, u.LastActivityOn); err != nil {
		output.WriteString(fmt.Sprintf("last active   : %s\n", u.LastActivityOn))
	} else {
		output.WriteString(fmt.Sprintf("last active   : %s (about %s)\n", u.LastActivityOn, humanize.Time(lastActive)))
	}
	output.WriteString(fmt.Sprintf("is admin      : %v\n", u.IsAdmin))
	output.WriteString(fmt.Sprintf("o/identities  : %v\n", len(u.Identities)))
	for _, identity := range u.Identities {
		output.WriteString(fmt.Sprintf("  - %s (%s)\n", identity.Provider, identity.ExternUID))
	}
	return output.String()
}
