package constants

const (
	DefaultGitlabHostname      = "gitlab.com"
	DefaultPathToConfiguration = "./dev.yaml"
	DefaultPathToSQLite3DB     = "./dev.db"

	DateOnlyTimeFormat          = "2006-01-02"
	GithubAPITimeFormat         = "2006-01-02T15:04:05.99Z"
	GitlabAPITimeFormat         = "2006-01-02T15:04:05.999Z"
	PivotalTrackerAPITimeFormat = "2006-01-02T15:04:05Z"

	CheckSuccessFormat = "\033[32m✅\033[0m \033[1m%s\033[0m "
	CheckFailureFormat = "\033[31m❌ \033[1m%s\033[0m "
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
	GetCanonicalVerb            = "get"
	GetAliases                  = []string{"retrieve", "g"}
	GitlabCanonicalNoun         = "gitlab"
	GitlabAliases               = []string{"gl"}
	GithubCanonicalNoun         = "github"
	GithubAliases               = []string{"gh"}
	InitialiseCanonicalNoun     = "initialise"
	InitialiseAliases           = []string{"initialize", "init", "i"}
	NetworkCanonicalNoun        = "network"
	NetworkAliases              = []string{"networks", "net", "nets", "nw"}
	NotificationsCanonicalNoun  = "notifications"
	NotificationsAliases        = []string{"notification", "notif", "notifs", "n"}
	OpenCanonicalVerb           = "open"
	OpenAliases                 = []string{"op", "o"}
	PivotalTrackerCanonicalNoun = "pivotaltracker"
	PivotalTrackerAliases       = []string{"pivotal", "pt"}
	RepositoryCanonicalNoun     = "repository"
	RepositoryAliases           = []string{"repo", "rp", "r"}
	ServerCanonicalNoun         = "server"
	ServerAliases               = []string{"syncer"}
	SoftwareCanonicalNoun       = "software"
	SoftwareAliases             = []string{"sw", "apps", "s"}
	StartCanonicalVerb          = "start"
	StartAliases                = []string{"st", "s"}
	WorkCanonicalNoun           = "work"
	WorkAliases                 = []string{"stories", "tasks", "tickets", "w"}
)
