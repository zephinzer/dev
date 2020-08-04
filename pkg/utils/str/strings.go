package str

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// StringKeyGetter is the signature for a function used in .Dedupe
type StringKeyGetter func(string) string

// Dedupe de-duplicates the provided :values slice and returns the de-duplicated
// slice
func Dedupe(values []string, keyGetter ...StringKeyGetter) []string {
	getKey := func(input string) string { return input }
	if len(keyGetter) > 0 {
		getKey = keyGetter[0]
	}
	output := []string{}
	seen := map[string]bool{}
	for _, value := range values {
		key := getKey(value)
		if value, ok := seen[key]; ok && value {
			continue
		}
		output = append(output, value)
		seen[key] = true
	}
	return output
}

// GetSshCloneUrlFromHttpLinkUrl returns the Git SSH clone URL given a HTTP link URL,
// for example if the repository link is https://github.com/zephinzer/dev, the returned
// clone URL will be git@github.com:zephinzer/dev.git
func GetSshCloneUrlFromHttpLinkUrl(httpLinkUrl string) (string, error) {
	u, err := url.Parse(httpLinkUrl)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("git@%s:%s.git", u.Hostname(), strings.TrimLeft(u.Path, "/")), nil
}

// GetHttpCloneUrlFromHttpLink returns the Git HTTP clone URL given a HTTP link URL,
// for example if the repository link is https://github.com/zephinzer/dev, the returned
// clone URL will be https://github.com/zephinzer/dev.git
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

func IsEmpty(test string) bool {
	return test == *new(string)
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
