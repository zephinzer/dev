package validator

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	keySchema   = "schema"
	keyUsername = "username"
	keyPassword = "password"
	keyHostname = "hostname"
	keyPort     = "port"
	keyUser     = "user"
	keyPath     = "path"
)

var (
	// RegexForHTTP defines a regular expression with capture groups that can be used to parse/validate
	// HTTP(S) URLs
	RegexForHTTP = fmt.Sprintf(
		`^(?P<%s>http[s]*)\:\/\/`+ // schema
			`(?:(?P<%s>[a-zA-Z0-9\_\.\-]+)`+ // username (optional)
			`(?:\:(?P<%s>.*))?@)*`+ // password (optional)
			`(?P<%s>[a-zA-Z0-9\.\-\_]+)`+ // hostname
			`(?:\:(?P<%s>\d+?))?\/`+ // port (optional)
			`(?P<%s>[a-zA-Z0-9\.\-\_]+)\/`+ // user
			`(?P<%s>[a-zA-Z0-9\/\.\-\_]+)`+ // path
			`\.git$`, // mandatory postfix
		keySchema, keyUsername, keyPassword, keyHostname, keyPort, keyUser, keyPath,
	)

	// RegexForSSH defines a regular expression with capture groups that can be used to parse/validate
	// SSH URLs
	RegexForSSH = fmt.Sprintf(
		`^(?P<%s>[a-zA-Z0-9\-\_\.]+?)@`+ // username
			`(?P<%s>[a-zA-Z0-9\-\_\.]+?):`+ // hostname
			`(?:(?P<%s>\d*?)\/)?`+ // port (optional)
			`(?P<%s>[a-zA-Z0-9\-\_\.]+)\/`+ // user
			`(?P<%s>[a-zA-Z0-9\-\_\.\/]+?)`+
			`\.git$`,
		keyUsername, keyHostname, keyPort, keyUser, keyPath,
	)
)

// URL defines the structure returnable by ParseURL
type URL struct {
	Schema   string
	Username string
	Password string
	Hostname string
	Port     string
	User     string
	Path     string
}

// String returns a string output of the URL instance
func (u URL) String() string {
	var builder strings.Builder
	if u.Schema != "http" {
		builder.WriteString("https://")
	}
	builder.WriteString(u.Hostname)
	if u.Port != "" {
		builder.WriteString(fmt.Sprintf(":%s", u.Port))
	}
	builder.WriteString(fmt.Sprintf("/%s/%s", u.User, u.Path))
	return builder.String()
}

// IsHTTPUrl is used for testing if a provided :test string is a valid HTTP(S) URL
func IsHTTPUrl(test string) bool {
	return regexp.MustCompile(RegexForHTTP).MatchString(test)
}

// IsSSHUrl is used for testing if a provided :test string is a valid SSH URL
func IsSSHUrl(test string) bool {
	return regexp.MustCompile(RegexForSSH).MatchString(test)
}

// ParseURL returns a new *URL based on the remoteURL
func ParseURL(remoteURL string) (*URL, error) {
	var selectedRegex string
	switch true {
	case IsHTTPUrl(remoteURL):
		selectedRegex = RegexForHTTP
	case IsSSHUrl(remoteURL):
		selectedRegex = RegexForSSH
	default:
		return nil, fmt.Errorf("provided url '%s' is not a valid git remote url", remoteURL)
	}
	compiledRegex := regexp.MustCompile(selectedRegex)
	fields := parseRegexIntoFields(compiledRegex, remoteURL)
	return serializeFieldsToURL(fields), nil
}

// parseRegexIntoFields is a utility function that is called by ParseURL to create a
// map of strings from the matches of a compiled regular expression
func parseRegexIntoFields(compiledRegex *regexp.Regexp, parse string) map[string]string {
	fields := map[string]string{}
	keys := compiledRegex.SubexpNames()
	values := compiledRegex.FindStringSubmatch(parse)
	for i := 1; i < len(keys); i++ {
		fields[keys[i]] = values[i]
	}
	return fields
}

// serializeFieldsToURL is a utility function called by ParseURL to serialize a
// map of strings into a *URL instance
func serializeFieldsToURL(fields map[string]string) *URL {
	url := URL{}
	if len(fields[keySchema]) > 0 {
		url.Schema = fields[keySchema]
	}
	if len(fields[keyUsername]) > 0 {
		url.Username = fields[keyUsername]
	}
	if len(fields[keyPassword]) > 0 {
		url.Password = fields[keyPassword]
	}
	if len(fields[keyHostname]) > 0 {
		url.Hostname = fields[keyHostname]
	}
	if len(fields[keyPort]) > 0 {
		url.Port = fields[keyPort]
	}
	if len(fields[keyUser]) > 0 {
		url.User = fields[keyUser]
	}
	if len(fields[keyPath]) > 0 {
		url.Path = fields[keyPath]
	}
	return &url
}
