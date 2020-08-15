package github

import (
	"time"
)

// APIv3Notifications is an alias for a slice of APIv3Notification type objects
type APIv3Notifications []APIv3Notification

// APIv3Notification represents the structure of a returned response
// from the Github API from a call to https://api.github.com/notifications
type APIv3Notification struct {
	ID              string                      `json:"id"`
	Unread          bool                        `json:"unread"`
	Reason          string                      `json:"reason"`
	UpdatedAt       time.Time                   `json:"updated_at"`
	LastReadAt      interface{}                 `json:"last_read_at"`
	Subject         APIv3NotificationSubject    `json:"subject"`
	Repository      APIv3NotificationRepository `json:"repository"`
	URL             string                      `json:"url"`
	SubscriptionURL string                      `json:"subscription_url"`
}

// APIv3NotificationSubject is the structure for the Subject field of the main
// APIv3Notification object
type APIv3NotificationSubject struct {
	Title            string `json:"title"`
	URL              string `json:"url"`
	LatestCommentURL string `json:"latest_comment_url"`
	Type             string `json:"type"`
}

// APIv3NotificationRepository is the structure for the Repository field of the main
// APIv3Notification object
type APIv3NotificationRepository struct {
	ID               int                              `json:"id"`
	NodeID           string                           `json:"node_id"`
	Name             string                           `json:"name"`
	FullName         string                           `json:"full_name"`
	Private          bool                             `json:"private"`
	Owner            APIv3NotificationRepositoryOwner `json:"owner"`
	HTMLURL          string                           `json:"html_url"`
	Description      interface{}                      `json:"description"`
	Fork             bool                             `json:"fork"`
	URL              string                           `json:"url"`
	ForksURL         string                           `json:"forks_url"`
	KeysURL          string                           `json:"keys_url"`
	CollaboratorsURL string                           `json:"collaborators_url"`
	TeamsURL         string                           `json:"teams_url"`
	HooksURL         string                           `json:"hooks_url"`
	IssueEventsURL   string                           `json:"issue_events_url"`
	EventsURL        string                           `json:"events_url"`
	AssigneesURL     string                           `json:"assignees_url"`
	BranchesURL      string                           `json:"branches_url"`
	TagsURL          string                           `json:"tags_url"`
	BlobsURL         string                           `json:"blobs_url"`
	GitTagsURL       string                           `json:"git_tags_url"`
	GitRefsURL       string                           `json:"git_refs_url"`
	TreesURL         string                           `json:"trees_url"`
	StatusesURL      string                           `json:"statuses_url"`
	LanguagesURL     string                           `json:"languages_url"`
	StargazersURL    string                           `json:"stargazers_url"`
	ContributorsURL  string                           `json:"contributors_url"`
	SubscribersURL   string                           `json:"subscribers_url"`
	SubscriptionURL  string                           `json:"subscription_url"`
	CommitsURL       string                           `json:"commits_url"`
	GitCommitsURL    string                           `json:"git_commits_url"`
	CommentsURL      string                           `json:"comments_url"`
	IssueCommentURL  string                           `json:"issue_comment_url"`
	ContentsURL      string                           `json:"contents_url"`
	CompareURL       string                           `json:"compare_url"`
	MergesURL        string                           `json:"merges_url"`
	ArchiveURL       string                           `json:"archive_url"`
	DownloadsURL     string                           `json:"downloads_url"`
	IssuesURL        string                           `json:"issues_url"`
	PullsURL         string                           `json:"pulls_url"`
	MilestonesURL    string                           `json:"milestones_url"`
	NotificationsURL string                           `json:"notifications_url"`
	LabelsURL        string                           `json:"labels_url"`
	ReleasesURL      string                           `json:"releases_url"`
	DeploymentsURL   string                           `json:"deployments_url"`
}

// APIv3NotificationRepositoryOwner is the structure for the Owner field of the main
// APIv3NotificationRepository object
type APIv3NotificationRepositoryOwner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// APIv3UserResponse defines the response structure from the
// https://api.github.com/user endpoint
type APIv3UserResponse struct {
	Login                   string                `json:"login"`
	ID                      int                   `json:"id"`
	NodeID                  string                `json:"node_id"`
	AvatarURL               string                `json:"avatar_url"`
	GravatarID              string                `json:"gravatar_id"`
	URL                     string                `json:"url"`
	HTMLURL                 string                `json:"html_url"`
	FollowersURL            string                `json:"followers_url"`
	FollowingURL            string                `json:"following_url"`
	GistsURL                string                `json:"gists_url"`
	StarredURL              string                `json:"starred_url"`
	SubscriptionsURL        string                `json:"subscriptions_url"`
	OrganizationsURL        string                `json:"organizations_url"`
	ReposURL                string                `json:"repos_url"`
	EventsURL               string                `json:"events_url"`
	ReceivedEventsURL       string                `json:"received_events_url"`
	Type                    string                `json:"type"`
	SiteAdmin               bool                  `json:"site_admin"`
	Name                    string                `json:"name"`
	Company                 string                `json:"company"`
	Blog                    string                `json:"blog"`
	Location                string                `json:"location"`
	Email                   interface{}           `json:"email"`
	Hireable                bool                  `json:"hireable"`
	Bio                     string                `json:"bio"`
	PublicRepos             int                   `json:"public_repos"`
	PublicGists             int                   `json:"public_gists"`
	Followers               int                   `json:"followers"`
	Following               int                   `json:"following"`
	CreatedAt               string                `json:"created_at"`
	UpdatedAt               string                `json:"updated_at"`
	PrivateGists            int                   `json:"private_gists"`
	TotalPrivateRepos       int                   `json:"total_private_repos"`
	OwnedPrivateRepos       int                   `json:"owned_private_repos"`
	DiskUsage               int                   `json:"disk_usage"`
	Collaborators           int                   `json:"collaborators"`
	TwoFactorAuthentication bool                  `json:"two_factor_authentication"`
	Plan                    APIv3UserResponsePlan `json:"plan"`
}

// APIv3UserResponsePlan defines the structure for the Plan property of the main
// APIv3UserResponse class
type APIv3UserResponsePlan struct {
	Name          string `json:"name"`
	Space         int    `json:"space"`
	Collaborators int    `json:"collaborators"`
	PrivateRepos  int    `json:"private_repos"`
}
