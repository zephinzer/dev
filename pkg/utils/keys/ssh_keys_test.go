package keys

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetSSHTests struct {
	suite.Suite
}

func TestGetSSH(t *testing.T) {
	suite.Run(t, &GetSSHTests{})
}

func (s *GetSSHTests) Test_GetSSH() {
	keys, err := GetSSH("./tests")
	s.Nil(err)
	if err != nil {
		return
	}
	// this value of 8 is the number of key files in the ./tests directory
	s.Len(keys, 8)
	for _, key := range keys {
		s.NotContains(key.Path, "not_a_key")
		s.NotContains(key.Path, "note.md")
		s.NotNil(key.Content, "all keys should contain content")
		s.NotEmpty(key.Fingerprint, "all keys should have a unique identifying hash")
		s.NotPanicsf(func() {
			if _, err := os.Lstat(key.Path); err != nil {
				panic(err)
			}
		}, "key.Path, '%s', should be a file that's accessible", key.Path)
		if len(key.Name) > 0 {
			s.Contains(key.String(), key.Name)
		}
	}
}
