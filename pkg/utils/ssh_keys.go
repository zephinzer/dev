package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"golang.org/x/crypto/ssh"
)

func GetLocalSSHKeys() ([]string, map[string]string, error) {
	switch runtime.GOOS {
	case "linux":
		fallthrough
	case "darwin":
		homeDir, getUserHomeDirError := os.UserHomeDir()
		if getUserHomeDirError != nil {
			return nil, nil, fmt.Errorf("failed to retrieve home directory of current user with id '%v': %s", os.Getuid(), getUserHomeDirError)
		}
		sshKeysDirectory := path.Join(homeDir, "/.ssh")
		directoryList, readDirError := ioutil.ReadDir(sshKeysDirectory)
		if readDirError != nil {
			return nil, nil, fmt.Errorf("failed to list directory at '%s': %s", sshKeysDirectory, readDirError)
		}
		privateKeys := []string{}
		authorizedKeys := map[string]string{}
		for _, directoryListing := range directoryList {
			if !directoryListing.IsDir() && directoryListing.Name() != "known_hosts" {
				possibleKeyPath := path.Join(sshKeysDirectory, directoryListing.Name())
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
	default:
		return nil, nil, fmt.Errorf("'%s' is an unsupported platform", runtime.GOOS)
	}
}
