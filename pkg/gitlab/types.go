package gitlab

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/usvc/dev/internal/constants"
)

type Config struct {
	Accounts Accounts `json:"accounts" yaml:"accounts"`
}

type Accounts []Account

type Account struct {
	// Name is the user-defined label for this account
	Name string `json:"name" yaml:"name"`
	// Hostname is the hostname of the Gitlab installation. If this is not provided
	// this should default to "gitlab.com"
	Hostname string `json:"hostname" yaml:"hostname"`
	// AccessToken is your Gitlab Personal Access Token which can be generated
	// from this page at https://gitlab.com/profile/personal_access_tokens
	//
	// The required Scopes are as follows:
	// - read_user
	// - read_api
	// - read_repository
	// - read_registry
	AccessToken string `json:"accessToken" yaml:"accessToken"`
}

// APIv4TodoResponse defines the response structure for a call to the Gitlab
// API endpoint at https://GITLAB_HOSTNAME/api/v4/todos
type APIv4TodoResponse []APIv4Todo

func (t APIv4TodoResponse) String() string {
	var output strings.Builder
	for i := 0; i < len(t); i++ {
		output.WriteString(t[i].String())
		output.Write([]byte{'\n', '\n'})
	}
	return output.String()
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

type APIv4Identities []APIv4Identity

type APIv4Identity struct {
	Provider  string `json:"provider"`
	ExternUID string `json:"extern_uid"`
}

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

func (t APIv4Todo) String() string {
	tag := t.Project.NameWithNamespace
	title := t.Target.Title
	message := t.Body + "\n"
	if t.Body == title {
		message = ""
	}
	state := t.State
	link := t.TargetURL
	createdAt := humanize.Time(t.CreatedAt)
	return fmt.Sprintf("[%s] %s\n%s> state: %s\n> link: %s\n> when: about %v", tag, title, message, state, link, createdAt)
}

type APIv4Project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
}

type APIv4Author struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	ID        int    `json:"id"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type APIv4Target struct {
	ID                        int            `json:"id"`
	Iid                       int            `json:"iid"`
	ProjectID                 int            `json:"project_id"`
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	State                     string         `json:"state"`
	CreatedAt                 time.Time      `json:"created_at"`
	UpdatedAt                 time.Time      `json:"updated_at"`
	TargetBranch              string         `json:"target_branch"`
	SourceBranch              string         `json:"source_branch"`
	Upvotes                   int            `json:"upvotes"`
	Downvotes                 int            `json:"downvotes"`
	Author                    APIv4Author    `json:"author"`
	Assignee                  APIv4Author    `json:"assignee"`
	SourceProjectID           int            `json:"source_project_id"`
	TargetProjectID           int            `json:"target_project_id"`
	Labels                    []interface{}  `json:"labels"`
	WorkInProgress            bool           `json:"work_in_progress"`
	Milestone                 APIv4Milestone `json:"milestone"`
	MergeWhenPipelineSucceeds bool           `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string         `json:"merge_status"`
	Subscribed                bool           `json:"subscribed"`
	UserNotesCount            int            `json:"user_notes_count"`
}

type APIv4Milestone struct {
	ID          int         `json:"id"`
	Iid         int         `json:"iid"`
	ProjectID   int         `json:"project_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	State       string      `json:"state"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DueDate     interface{} `json:"due_date"`
}
