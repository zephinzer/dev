package gitlab

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

type Config struct {
	Accounts Accounts `yaml:"accounts"`
}

type Accounts []Account

type Account struct {
	// Name is the user-defined label for this account
	Name string `yaml:"name"`
	// Hostname is the hostname of the Gitlab installation. If this is not provided
	// this should default to "gitlab.com"
	Hostname string `yaml:"hostname"`
	// AccessToken is your Gitlab Personal Access Token which can be generated
	// from this page at https://gitlab.com/profile/personal_access_tokens
	//
	// The required Scopes are as follows:
	// - read_user
	// - read_api
	// - read_repository
	// - read_registry
	AccessToken string `yaml:"accessToken"`
}

// APIv4TodoResponse defines the response structure for a call to the Gitlab
// API endpoint at https://gitlab/api/v4/todos
type APIv4TodoResponse []APIv4Todo

func (t APIv4TodoResponse) String() string {
	var output strings.Builder
	for i := 0; i < len(t); i++ {
		output.WriteString(t[i].String())
		output.Write([]byte{'\n', '\n'})
	}
	return output.String()
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
