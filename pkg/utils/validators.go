package utils

import "regexp"

type GitRemoteURLType string

const (
	HTTPRemoteType GitRemoteURLType = "http"
	SSHRemoteType  GitRemoteURLType = "ssh"
	NotRemoteType  GitRemoteURLType = "invalid"
)

var (
	RegexForHTTP = `^(?P<schema>http[s]*)\:\/\/(?:(?P<username>[a-zA-Z0-9\_\.\-]+)(?:\:(?P<password>.*))?@)*(?P<hostname>[a-zA-Z0-9\.\-\_]+)(?:\:(?P<port>\d+?))?\/(?P<user>[a-zA-Z0-9\.\-\_]+)\/(?P<path>[a-zA-Z0-9\/\.\-\_]+)\.git$`
	RegexForSSH  = `^(?P<username>[a-zA-Z0-9\-\_\.]+?)@(?P<hostname>[a-zA-Z0-9\-\_\.]+?):(?:(?P<port>\d*?)\/)?(?P<user>[a-zA-Z0-9\-\_\.]+)\/(?P<path>[a-zA-Z0-9\-\_\.\/]+?)\.git$`
)

func GetGitRemoteURLType(testing string) GitRemoteURLType {
	switch true {
	case regexp.MustCompile(RegexForSSH).MatchString(testing):
		return SSHRemoteType
	case regexp.MustCompile(RegexForHTTP).MatchString(testing):
		return HTTPRemoteType
	}
	return NotRemoteType
}

type GitRemoteStructure struct {
	Type     GitRemoteURLType
	Schema   string
	Username string
	Password string
	Hostname string
	Port     string
	User     string
	Path     string
}

func GetGitRemoteStructure(remoteURL string) GitRemoteStructure {
	switch GetGitRemoteURLType(remoteURL) {
	case HTTPRemoteType:
		regexp.MustCompile(RegexForHTTP)
	case SSHRemoteType:
		regexp.MustCompile(RegexForHTTP)
	}
	return GitRemoteStructure{}
}
