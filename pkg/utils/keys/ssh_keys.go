package keys

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"golang.org/x/crypto/ssh"
)

// Type is a string symbol that indicates the type of key we're dealing with
type Type string

const (
	// TypeSSHPrivate indicates the key is a private key
	TypeSSHPrivate Type = "KEY_SSH_PRIVATE"
	// TypeSSHpublic indicates the key is a public key
	TypeSSHPublic Type = "KEY_SSH_PUBLIC"
	// TypePGP indicates a Pretty Good Privacy key
	TypePGP Type = "KEY_PGP"
	// TypeGPG is an alias for TypePGP (heh)
	TypeGPG Type = TypePGP
)

// Key represents a logical key
type Key struct {
	// Content contains the content of the key
	Content []byte
	// Fingerprint contains a unique string value that identifies this key
	Fingerprint string
	// IsPasswordProtected defines if the key is protected by a password
	IsPasswordProtected bool
	// Name contains the name of the key (applies to only TypeSSHPublic keys if it's specified)
	Name string
	// Path contains a usable relative path to the key
	Path string
	// Type contains the type of key
	Type Type
}

func (k Key) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s %s %s", k.Type, k.Fingerprint, k.Path))
	if len(k.Name) > 0 {
		builder.WriteString(fmt.Sprintf(" (%s)", k.Name))
	}
	return builder.String()
}

func GetSSH(keysDirectory string) ([]Key, error) {
	directoryList, readDirError := ioutil.ReadDir(keysDirectory)
	if readDirError != nil {
		return nil, fmt.Errorf("failed to list directory at '%s': %s", keysDirectory, readDirError)
	}
	keys := []Key{}
	for _, directoryListing := range directoryList {
		if !directoryListing.IsDir() {
			var err error
			key := Key{}
			key.Path = path.Join(keysDirectory, directoryListing.Name())
			if key.Content, err = ioutil.ReadFile(key.Path); err != nil {
				continue
			}
			var authorizedKey ssh.PublicKey
			if authorizedKey, key.Name, _, _, err = ssh.ParseAuthorizedKey(key.Content); err == nil {
				key.Type = TypeSSHPublic
				key.Fingerprint = ssh.FingerprintSHA256(authorizedKey)
				keys = append(keys, key)
				continue
			} else if _, err = ssh.ParsePrivateKey(key.Content); err == nil {
				key.Type = TypeSSHPrivate
			} else if strings.Contains(err.Error(), "this private key is passphrase protected") {
				key.Type = TypeSSHPrivate
				key.IsPasswordProtected = true
				err = nil
			}
			if err == nil {
				if len(key.Fingerprint) == 0 {
					hasher := sha256.New()
					if _, err = hasher.Write(key.Content); err == nil {
						key.Fingerprint = fmt.Sprintf("SHA256:%s", base64.URLEncoding.EncodeToString(hasher.Sum(nil)))
					}
				}
				keys = append(keys, key)
			}
		}
	}
	return keys, nil
}
