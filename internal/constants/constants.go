package constants

const (
	// CanonicalRepositoryURL defines the working repository for `dev`'s code
	CanonicalRepositoryURL = "https://gitlab.com/zephinzer/dev"

	// RepositoryURLConfiguration defines a URL to the working repository for `dev`'s configuration
	RepositoryURLConfiguration = "https://gitlab.com/zephinzer/dev/-/blob/master/docs/configuration.md"

	// DefaultGitlabHostname is the default Gitlab hostname when it's not specified
	DefaultGitlabHostname = "gitlab.com"

	// DefaultNetworkCheckMethod is the default method to use when issuing HTTP calls
	DefaultNetworkCheckMethod = "GET"

	// DefaultPathToSQLite3DB is the default location of the database
	DefaultPathToSQLite3DB = "./dev.db"

	// DevTimeFormat is the printed time format for messages intended for logging from the `dev` application
	DevTimeFormat = "20060102_150405"

	// DevHumanTimeFormat is the printed time format for messages intended for human reading from the `dev` application
	DevHumanTimeFormat = "2006-01-02 15:04:05"

	// DateOnlyTimeFormat is the time format when printing just the date
	DateOnlyTimeFormat = "2006-01-02"

	// GithubAPITimeFormat defines the timestamp format of timestamps returned by the Github API
	GithubAPITimeFormat = "2006-01-02T15:04:05.99Z"

	// GitlabAPITimeFormat defines the timestamp format of timestamps returned by the Gitlab API
	GitlabAPITimeFormat = "2006-01-02T15:04:05.999Z"

	// PivotalTrackerAPITimeFormat defines the timestampformat of timestamps returned by the Pivotal Tracker API
	PivotalTrackerAPITimeFormat = "2006-01-02T15:04:05Z"

	// ExitOK indicates there are no errors
	ExitOK = 0

	// ExitErrorSystem indicates the error is because of the system
	ExitErrorSystem = 1

	// ExitErrorUser indicates the error is because of the user
	ExitErrorUser = 2

	// ExitErrorInput indicates the error is because of the user's input
	ExitErrorInput = 4

	// ExitErrorConfiguration indicates the error is because of the
	// configuration
	ExitErrorConfiguration = 8

	// ExitErrorApplication indicates the error is a logical bug with the
	// code
	ExitErrorApplication = 16

	// ExitErrorValidation indicates the error is related to validation
	ExitErrorValidation = 32

	// CheckSuccessFormat is the standardised format used to print messages
	// where the check was successful
	CheckSuccessFormat = "\033[32m✅\033[0m \033[32m%s\033[0m "

	// CheckFailureFormat is the standardised format used to print messsages
	// where the check failed
	CheckFailureFormat = "\033[31m❌ \033[31m%s\033[0m "

	// CheckSkippedFormat is the standardised format used to print messages
	// where the check was skipped
	CheckSkippedFormat = "\033[33m👀\033[0m \033[33m%s\033[0m "
)

var (
	DefaultSystrayIcon = SystrayIconLight

	AccountCanonicalNoun        = "account"
	AccountAliases              = []string{"accounts", "acct", "accts", "acc", "accs", "a"}
	AddCanonicalVerb            = "add"
	AddAliases                  = []string{"+"}
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
	TrelloCanonicalNoun         = "trello"
	TrelloAliases               = []string{"tr"}
	WorkCanonicalNoun           = "work"
	WorkAliases                 = []string{"stories", "tasks", "tickets", "w"}
	WorkspaceCanonicalNoun      = "workspace"
	WorkspaceAliases            = []string{"ws"}
)
