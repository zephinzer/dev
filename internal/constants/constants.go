package constants

const (
	DefaultGitlabHostname      = "gitlab.com"
	DefaultPathToConfiguration = "./dev.yaml"
	DefaultPathToSQLite3DB     = "./dev.db"

	DateOnlyTimeFormat          = "2006-01-02"
	GithubAPITimeFormat         = "2006-01-02T15:04:05.99Z"
	GitlabAPITimeFormat         = "2006-01-02T15:04:05.999Z"
	PivotalTrackerAPITimeFormat = "2006-01-02T15:04:05Z"
)

var (
	AccountCanonicalNoun        = "account"
	AccountAliases              = []string{"accounts", "acc", "accs", "a"}
	CheckCanonicalVerb          = "check"
	CheckAliases                = []string{"c", "verify"}
	ClientCanonicalNoun         = "client"
	ClientAliases               = []string{"cl", "c"}
	ConfigurationCanonicalNoun  = "configuration"
	ConfigurationAliases        = []string{"config", "conf", "cf", "c"}
	DatabaseCanonicalNoun       = "database"
	DatabaseAliases             = []string{"db"}
	GetCanonicalNoun            = "get"
	GetAliases                  = []string{"retrieve", "g"}
	GitlabCanonicalNoun         = "gitlab"
	GitlabAliases               = []string{"gl"}
	GithubCanonicalNoun         = "github"
	GithubAliases               = []string{"gh"}
	InitialiseCanonicalNoun     = "initialise"
	InitialiseAliases           = []string{"initialize", "init", "i"}
	OpenCanonicalVerb           = "open"
	OpenAliases                 = []string{"op", "o"}
	NotificationsCanonicalNoun  = "notifications"
	NotificationsAliases        = []string{"notification", "notif", "notifs", "n"}
	PivotalTrackerCanonicalNoun = "pivotaltracker"
	PivotalTrackerAliases       = []string{"pivotal", "pt"}
	RepositoryCanonicalNoun     = "repository"
	RepositoryAliases           = []string{"repo", "rp", "r"}
	SoftwareCanonicalNoun       = "software"
	SoftwareAliases             = []string{"sw", "apps", "s"}
	StartCanonicalVerb          = "start"
	StartAliases                = []string{"st", "s"}
	WorkCanonicalNoun           = "work"
	WorkAliases                 = []string{"stories", "tasks", "tickets", "w"}
)
