package constants

const (
	DefaultGitlabHostname      = "gitlab.com"
	DefaultPathToConfiguration = "./dev.yaml"
	DefaultPathToSQLite3DB     = "./dev.db"

	DateOnlyTimeFormat          = "2006-01-02"
	GitlabAPITimeFormat         = "2006-01-02T15:04:05.999Z"
	PivotalTrackerAPITimeFormat = "2006-01-02T15:04:05Z"
)

var (
	GitlabCanonicalNoun         = "gitlab"
	GitlabAliases               = []string{"gl"}
	GithubCanonicalNoun         = "github"
	GithubAliases               = []string{"gh"}
	PivotalTrackerCanonicalNoun = "pivotaltracker"
	PivotalTrackerAliases       = []string{"pivotal", "pt"}
)
