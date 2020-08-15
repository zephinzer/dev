package gitlab

import (
	"time"
)

type APIv4Author struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	ID        int    `json:"id"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type APIv4Identities []APIv4Identity

type APIv4Identity struct {
	Provider  string `json:"provider"`
	ExternUID string `json:"extern_uid"`
}

type APIv4Project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	NameWithNamespace string `json:"name_with_namespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"path_with_namespace"`
}

type APIv4Target struct {
	ID                        int                  `json:"id"`
	Iid                       int                  `json:"iid"`
	ProjectID                 int                  `json:"project_id"`
	Title                     string               `json:"title"`
	Description               string               `json:"description"`
	State                     string               `json:"state"`
	CreatedAt                 time.Time            `json:"created_at"`
	UpdatedAt                 time.Time            `json:"updated_at"`
	TargetBranch              string               `json:"target_branch"`
	SourceBranch              string               `json:"source_branch"`
	Upvotes                   int                  `json:"upvotes"`
	Downvotes                 int                  `json:"downvotes"`
	Author                    APIv4Author          `json:"author"`
	Assignee                  APIv4Author          `json:"assignee"`
	SourceProjectID           int                  `json:"source_project_id"`
	TargetProjectID           int                  `json:"target_project_id"`
	Labels                    []interface{}        `json:"labels"`
	WorkInProgress            bool                 `json:"work_in_progress"`
	Milestone                 APIv4TargetMilestone `json:"milestone"`
	MergeWhenPipelineSucceeds bool                 `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string               `json:"merge_status"`
	Subscribed                bool                 `json:"subscribed"`
	UserNotesCount            int                  `json:"user_notes_count"`
}

type APIv4TargetMilestone struct {
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
