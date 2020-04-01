package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DownloadTests struct {
	suite.Suite
}

func TestDownload(t *testing.T) {
	suite.Run(t, &DownloadTests{})
}

func (s *DownloadTests) TestDownload() {
	Download(DownloadOptions{
		URL:      "https://github.com/usvc/db/releases/download/v0.0.4/db_darwin_amd64.sha256",
		FilePath: "./tests/downloads/aaab",
	})
}
