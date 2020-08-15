package constants

import "time"

const (
	// DefaultAPICallTimeout is the default timeout to apply for external API calls
	DefaultAPICallTimeout = time.Second * 10

	// DefaultGitlabHostname is the default Gitlab hostname when it's not specified
	DefaultGitlabHostname = "gitlab.com"

	// DefaultNetworkCheckMethod is the default method to use when issuing HTTP calls
	DefaultNetworkCheckMethod = "GET"

	// DefaultPathToSQLite3DB is the default location of the database
	DefaultPathToSQLite3DB = "~/dev.db"

	// DefaultRedactedString is the default string used to replace access tokens/secrets
	DefaultRedactedString = "[REDACTED]"
)
