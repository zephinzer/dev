package str

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	DefaultUsername = "git"
)

func ParseGitUrl(repoUrl string) (*GitUrl, error) {
	if strings.HasPrefix(repoUrl, "https://") { // is a https url
		parsedUrl, err := url.Parse(repoUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to parse https url: %s", err)
		}
		gitUrl := &GitUrl{}

		username := parsedUrl.User.Username()
		gitUrl.Username = &username

		if password, passwordSet := parsedUrl.User.Password(); passwordSet {
			gitUrl.AccessToken = &password
		}

		gitUrl.Hostname = parsedUrl.Hostname()
		port := parsedUrl.Port()
		if port != "" {
			if portNumber, err := strconv.Atoi(port); err == nil {
				gitUrl.Port = &portNumber
			}
		}

		gitUrl.Path = strings.Trim(parsedUrl.Path, "/")
		return gitUrl, nil
	}

	// assume its an ssh url

	gitUrl := &GitUrl{}
	sshText := "ssh"
	sshPrefix := sshText + "://"
	if strings.HasPrefix(repoUrl, sshPrefix) {
		gitUrl.Protocol = &sshText
		repoUrl = strings.ReplaceAll(repoUrl, sshPrefix, "")
	}

	if strings.Contains(repoUrl, "@") {
		sshUrlParts := strings.Split(repoUrl, "@")
		if len(sshUrlParts) > 2 {
			return nil, fmt.Errorf("failed to get a valid url: more than 2 '@' symbols detected")
		}
		gitUrl.Username = &sshUrlParts[0]
		repoUrl = sshUrlParts[1]
	}

	repoUrlParts := strings.Split(repoUrl, ":")
	gitUrl.Hostname = repoUrlParts[0]
	repoPath := repoUrlParts[1]

	repoPathParts := strings.Split(repoPath, "/")
	possiblePort := repoPathParts[0]
	if port, err := strconv.Atoi(possiblePort); err == nil {
		if port < 1<<16 && port > 0 {
			gitUrl.Port = &port
			repoPathParts = repoPathParts[1:]
		}
	}
	gitUrl.Path = strings.Join(repoPathParts, "/")

	return gitUrl, nil
}

type GitUrl struct {
	// Protocol is an optional protocol to specify that is
	// only populated when parsing a string into a GitUrl
	Protocol *string

	// Username is an optional username that defaults to
	// DefaultUsername for SSH URLs and omitted for HTTPS URLs
	Username *string

	// AccessToken is only used when retrieving the HTTPS URL
	AccessToken *string

	// Hostname is the hostname of the Git server
	Hostname string

	// Port is the optional port number where the Git server is
	// listening on
	Port *int

	// Path is the path to the repository
	Path string
}

func (g GitUrl) GetBrowserUrl() string {
	repoHost := g.Hostname
	if g.Port != nil {
		repoHost = fmt.Sprintf("%s:%v", repoHost, *g.Port)
	}
	repoPath := strings.Trim(g.Path, "/")
	if strings.HasSuffix(repoPath, ".git") {
		repoPath = repoPath[:len(repoPath)-len(".git")]
	}
	return fmt.Sprintf("https://%s/%s", repoHost, repoPath)
}

func (g GitUrl) GetHttpsUrl() string {
	authString := ""
	if g.Username != nil {
		authString = *g.Username
		if g.AccessToken != nil {
			authString += ":" + *g.AccessToken
		}
		authString = authString + "@"
	}
	repoHost := g.Hostname
	if g.Port != nil {
		repoHost = fmt.Sprintf("%s:%v", repoHost, *g.Port)
	}
	repoPath := strings.Trim(g.Path, "/")
	if !strings.HasSuffix(repoPath, ".git") {
		repoPath = repoPath + ".git"
	}
	return fmt.Sprintf("https://%s%s/%s", authString, repoHost, repoPath)
}

func (g GitUrl) GetSshUrl() string {
	username := "git"
	if g.Username != nil {
		username = *g.Username
	}
	repoHost := g.Hostname
	pathDelimiter := ":"
	if g.Port != nil {
		repoHost = fmt.Sprintf("%s:%v", repoHost, *g.Port)
		pathDelimiter = "/"
	}
	repoPath := strings.Trim(g.Path, "/")
	if !strings.HasSuffix(repoPath, ".git") {
		repoPath = repoPath + ".git"
	}
	return fmt.Sprintf("%s@%s%s%s", username, repoHost, pathDelimiter, repoPath)
}
