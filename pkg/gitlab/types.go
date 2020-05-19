package gitlab

import (
	"time"
)

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
