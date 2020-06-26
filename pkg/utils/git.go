package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	gitconfig "github.com/go-git/go-git/v5/config"
	gitransport "github.com/go-git/go-git/v5/plumbing/transport"
)

func GetSshCloneUrlFromHttpLinkUrl(httpLinkUrl string) (string, error) {
	u, err := url.Parse(httpLinkUrl)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("git@%s:%s.git", u.Hostname(), strings.TrimLeft(u.Path, "/")), nil
}

func GetHttpCloneUrlFromHttpLink(httpLinkUrl string) (string, error) {
	return fmt.Sprintf("%s.git", httpLinkUrl), nil
}

func GetHttpLinkFromHttpCloneUrl(httpCloneUrl string) (string, error) {
	u, err := url.Parse(httpCloneUrl)
	if err != nil {
		return "", err
	}
	u2 := url.URL{
		Host:   u.Hostname(),
		Scheme: u.Scheme,
		Path:   u.Path[:len(u.Path)-len(".git")],
	}
	return u2.String(), nil
}

func GetHttpLinkFromSshCloneUrl(sshCloneUrl string) (string, error) {
	re := regexp.MustCompile(`(?P<scheme>ssh://)?(?P<username>[a-zA-Z_\.]*@){0,1}(?P<hostname>[a-zA-Z\-\.]*\:){0,1}(?P<path>[a-zA-Z0-9\.\-/]*).git`)
	var matches map[string]string
	if re.Match([]byte(sshCloneUrl)) {
		matches = map[string]string{}
		submatches := re.FindStringSubmatch(sshCloneUrl)
		names := re.SubexpNames()
		for i, name := range names {
			if i != 0 && name != "" {
				matches[name] = submatches[i]
			}
		}
		if scheme, ok := matches["scheme"]; ok {
			matches["scheme"] = strings.Trim(strings.Trim(scheme, "/"), ":")
		}
		if username, ok := matches["username"]; ok {
			matches["username"] = strings.Trim(username, "@")
		}
		if hostname, ok := matches["hostname"]; ok {
			matches["hostname"] = strings.Trim(hostname, ":")
		}
	}
	if matches == nil {
		return "", fmt.Errorf("unable to parse ssh clone url '%s'", sshCloneUrl)
	}
	u := url.URL{
		Host:   matches["hostname"],
		Path:   matches["path"],
		Scheme: "https",
	}
	return u.String(), nil
}

// GitClone does a `git clone` on the provided :cloneURL into the directory at
// :localPath. If the remote is a bare repository, a new repository will be
// initialised at :localPath with the `origin` remote set to :cloneURL
func GitClone(cloneURL, localPath string) error {
	_, cloneError := git.PlainClone(localPath, false, &git.CloneOptions{URL: cloneURL})
	if cloneError != nil {
		if cloneError == gitransport.ErrEmptyRemoteRepository {
			return GitInit(cloneURL, localPath)
		}
		return fmt.Errorf("failed to clone repository from url '%s': %s", cloneURL, cloneError)
	}
	return nil
}

func GitInit(cloneURL, localPath string) error {
	initialisedRepo, initError := git.PlainInit(localPath, false)
	if initError != nil {
		return fmt.Errorf("failed to initialise repository at '%s': %s", localPath, initError)
	}
	_, createRemoteError := initialisedRepo.CreateRemote(&gitconfig.RemoteConfig{
		Name: "origin",
		URLs: []string{cloneURL},
	})
	if createRemoteError != nil {
		return fmt.Errorf("failed to set remote origin for repository at '%s' tracked at '%s': %s", localPath, cloneURL, createRemoteError)
	}
	return nil
}
