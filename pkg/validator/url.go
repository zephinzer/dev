package validator

import (
	"fmt"
	stdlibURL "net/url"
	"path"
	"regexp"
	"strings"

	"github.com/zephinzer/dev/pkg/utils/str"
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
	// RegexForGitHTTP defines a regular expression with capture groups that can be used to parse/validate
	// HTTP(S) URLs
	RegexForGitHTTP = fmt.Sprintf(
		`^(?P<%s>http[s]*)\:\/\/`+ // schema
			`(?:(?P<%s>[a-zA-Z0-9\_\.\-]+)`+ // username (optional)
			`(?:\:(?P<%s>.*))?@)*`+ // password (optional)
			`(?P<%s>[a-zA-Z0-9\.\-\_]+)`+ // hostname
			`(?:\:(?P<%s>\d+?))?\/`+ // port (optional)
			`(?P<%s>[a-zA-Z0-9\.\-\_]+)`+ // user
			`(?P<%s>\/[a-zA-Z0-9\/\.\-\_]+)`+ // path
			`\.git$`, // mandatory postfix
		keySchema, keyUsername, keyPassword, keyHostname, keyPort, keyUser, keyPath,
	)

	// RegexForGitSSH defines a regular expression with capture groups that can be used to parse/validate
	// SSH URLs
	RegexForGitSSH = fmt.Sprintf(
		`^(?P<%s>[a-zA-Z0-9\-\_\.]+?)@`+ // username
			`(?P<%s>[a-zA-Z0-9\-\_\.]+?):`+ // hostname
			`(?:(?P<%s>\d*?)\/)?`+ // port (optional)
			`(?P<%s>[a-zA-Z0-9\-\_\.]+)`+ // user
			`(?P<%s>\/[a-zA-Z0-9\-\_\.\/]+?)`+
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
	Query    string
}

func (u URL) GetHTTPSString() string {
	var builder strings.Builder
	builder.WriteString("https://")
	if !str.IsEmpty(u.Username) && u.Username != "git" {
		builder.WriteString(u.Username)
		if !str.IsEmpty(u.Password) {
			builder.WriteString(fmt.Sprintf(":%s", u.Password))
		}
		builder.WriteByte('@')
	}
	builder.WriteString(u.Hostname)
	if !str.IsEmpty(u.Port) {
		builder.WriteString(fmt.Sprintf(":%s", u.Port))
	}
	urlPath := path.Join("/", u.User, u.Path)
	builder.WriteString(fmt.Sprintf("%s.git", urlPath))
	return builder.String()
}

func (u URL) GetSSHString() string {
	var builder strings.Builder
	builder.WriteString("git@")
	builder.WriteString(u.Hostname)
	builder.WriteByte(':')
	if !str.IsEmpty(u.Port) {
		builder.WriteString(fmt.Sprintf("%s/", u.Port))
	}
	urlPath := strings.Trim(path.Join(u.User, u.Path), "/")
	builder.WriteString(fmt.Sprintf("%s.git", urlPath))
	return builder.String()
}

// String returns a string output of the URL instance
func (u URL) String() string {
	var builder strings.Builder
	if !str.IsEmpty(u.Schema) {
		builder.WriteString(u.Schema)
	} else {
		builder.WriteString("https")
	}
	builder.WriteString("://")
	if !str.IsEmpty(u.Username) && u.Username != "git" {
		builder.WriteString(u.Username)
		if !str.IsEmpty(u.Password) {
			builder.WriteString(fmt.Sprintf(":%s", u.Password))
		}
		builder.WriteByte('@')
	}
	builder.WriteString(u.Hostname)
	if !str.IsEmpty(u.Port) {
		builder.WriteString(fmt.Sprintf(":%s", u.Port))
	}
	urlPath := fmt.Sprintf("%s", path.Join("/", u.User, u.Path))
	builder.WriteString(urlPath)
	if !str.IsEmpty(u.Query) {
		builder.WriteString(fmt.Sprintf("?%s", u.Query))
	}
	return builder.String()
}

// IsGitHTTPUrl is used for testing if a provided :test string is a valid HTTP(S) URL
func IsGitHTTPUrl(test string) bool {
	return regexp.MustCompile(RegexForGitHTTP).MatchString(test)
}

// IsGitSSHUrl is used for testing if a provided :test string is a valid SSH URL
func IsGitSSHUrl(test string) bool {
	return regexp.MustCompile(RegexForGitSSH).MatchString(test)
}

// ParseURL returns a new *URL based on the remoteURL
func ParseURL(remoteURL string) (*URL, error) {
	var selectedRegex string
	switch true {
	case IsGitHTTPUrl(remoteURL):
		selectedRegex = RegexForGitHTTP
	case IsGitSSHUrl(remoteURL):
		selectedRegex = RegexForGitSSH
	default:
		return parseURLWithStdlib(remoteURL)
	}
	compiledRegex := regexp.MustCompile(selectedRegex)
	fields := parseRegexIntoFields(compiledRegex, remoteURL)
	return serializeFieldsToURL(fields), nil
}

// parseURLWithStdlib parses the provided url using the default net/url
// library provided by Go for urls that don't match a Git URL
func parseURLWithStdlib(remoteURL string) (*URL, error) {
	parsedURL, parseURLError := stdlibURL.Parse(remoteURL)
	if parseURLError != nil {
		return nil, fmt.Errorf("failed to parse url '%s': %s", remoteURL, parseURLError)
	}
	password, passwordSet := parsedURL.User.Password()
	finalURL := URL{
		Hostname: parsedURL.Hostname(),
		Port:     parsedURL.Port(),
		Schema:   parsedURL.Scheme,
		Username: parsedURL.User.Username(),
		Path:     parsedURL.Path,
		Query:    parsedURL.Query().Encode(),
	}
	if passwordSet {
		finalURL.Password = password
	}
	return &finalURL, nil
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

// serializeFieldsToURL is a utility function to serialize a
// map of strings from regex capture groups into a *URL instance
func serializeFieldsToURL(fields map[string]string) *URL {
	url := URL{}
	if !str.IsEmpty(fields[keySchema]) {
		url.Schema = fields[keySchema]
	}
	if !str.IsEmpty(fields[keyUsername]) {
		url.Username = fields[keyUsername]
	}
	if !str.IsEmpty(fields[keyPassword]) {
		url.Password = fields[keyPassword]
	}
	if !str.IsEmpty(fields[keyHostname]) {
		url.Hostname = fields[keyHostname]
	}
	if !str.IsEmpty(fields[keyPort]) {
		url.Port = fields[keyPort]
	}
	if !str.IsEmpty(fields[keyUser]) {
		url.User = fields[keyUser]
	}
	if !str.IsEmpty(fields[keyPath]) {
		url.Path = fields[keyPath]
	}
	return &url
}
