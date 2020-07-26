package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GitTests struct {
	suite.Suite
}

func TestGit(t *testing.T) {
	suite.Run(t, &GitTests{})
}
