package constants

const (
	DefaultGitlabHostname      = "gitlab.com"
	DefaultPathToConfiguration = "./dev.yaml"
	DefaultPathToSQLite3DB     = "./dev.db"

	DevTimeFormat               = "20060102_150405"
	DevHumanTimeFormat          = "2006-01-02 15:04:05"
	DateOnlyTimeFormat          = "2006-01-02"
	GithubAPITimeFormat         = "2006-01-02T15:04:05.99Z"
	GitlabAPITimeFormat         = "2006-01-02T15:04:05.999Z"
	PivotalTrackerAPITimeFormat = "2006-01-02T15:04:05Z"

	CheckSuccessFormat = "\033[32m✅\033[0m \033[1m%s\033[0m "
	CheckFailureFormat = "\033[31m❌ \033[1m%s\033[0m "
)

var (
	DefaultSystrayIcon = SystrayIconLight

	AccountCanonicalNoun        = "account"
	AccountAliases              = []string{"accounts", "acc", "accs", "a"}
	CheckCanonicalVerb          = "check"
	CheckAliases                = []string{"c", "verify"}
	ClientCanonicalNoun         = "client"
	ClientAliases               = []string{"cl", "c"}
	DebugCanonicalVerb          = "debug"
	DebugAliases                = []string{"test"}
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
	GotoCanonicalVerb           = "goto"
	GotoAliases                 = []string{"go", "go2"}
	InitialiseCanonicalNoun     = "initialise"
	InitialiseAliases           = []string{"initialize", "init", "i"}
	InstallCanonicalVerb        = "install"
	InstallAliases              = []string{"setup"}
	NetworkCanonicalNoun        = "network"
	NetworkAliases              = []string{"networks", "net", "nets", "nw"}
	NotificationsCanonicalNoun  = "notifications"
	NotificationsAliases        = []string{"notification", "notif", "notifs", "n"}
	OpenCanonicalVerb           = "open"
	OpenAliases                 = []string{"op", "o"}
	PivotalTrackerCanonicalNoun = "pivotaltracker"
	PivotalTrackerAliases       = []string{"pivotal", "pt"}
	RepositoryCanonicalNoun     = "repository"
	RepositoryAliases           = []string{"repositories", "repo", "repos", "rp", "r"}
	ServerCanonicalNoun         = "server"
	ServerAliases               = []string{"syncer"}
	SoftwareCanonicalNoun       = "software"
	SoftwareAliases             = []string{"softwares", "sw", "apps", "s"}
	StartCanonicalVerb          = "start"
	StartAliases                = []string{"st", "s"}
	SystemInfoCanonicalNoun     = "sysinfo"
	SystemInfoAliases           = []string{"system-info", "system", "si"}
	TelegramCanonicalNoun       = "telegram"
	TelegramAliases             = []string{"tg", "t.me"}
	WorkCanonicalNoun           = "work"
	WorkAliases                 = []string{"stories", "tasks", "tickets", "w"}
)
