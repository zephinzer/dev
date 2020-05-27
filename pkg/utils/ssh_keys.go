package utils

import (
	"fmt"
	"io/ioutil"
	"path"

	"golang.org/x/crypto/ssh"
)

func GetLocalSSHKeys(keysDirectory string) ([]string, map[string]string, error) {
	directoryList, readDirError := ioutil.ReadDir(keysDirectory)
	if readDirError != nil {
		return nil, nil, fmt.Errorf("failed to list directory at '%s': %s", keysDirectory, readDirError)
	}
	privateKeys := []string{}
	authorizedKeys := map[string]string{}
	for _, directoryListing := range directoryList {
		if !directoryListing.IsDir() && directoryListing.Name() != "known_hosts" {
			possibleKeyPath := path.Join(keysDirectory, directoryListing.Name())
			possibleKeyContent, readFileError := ioutil.ReadFile(possibleKeyPath)
			if readFileError != nil {
				continue
			}
			_, parsePrivateKeyError := ssh.ParsePrivateKey(possibleKeyContent)
			if parsePrivateKeyError == nil {
				privateKeys = append(privateKeys, possibleKeyPath)
				continue
			}
			_, comment, _, _, parseAuthorizedKeyError := ssh.ParseAuthorizedKey(possibleKeyContent)
			if parseAuthorizedKeyError == nil {
				authorizedKeys[possibleKeyPath] = comment
				continue
			}
		}
	}
	return privateKeys, authorizedKeys, nil
}
