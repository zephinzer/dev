package network

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DownloadEventTests struct {
	suite.Suite
}

func TestDownloadEvents(t *testing.T) {
	suite.Run(t, &DownloadEventTests{})
}

func (s *DownloadEventTests) Test_String() {
	event := DownloadEvent{
		State:        DownloadState("__state"),
		URL:          "__url",
		FilePath:     "__file_path",
		TempFilePath: "__temp_file_path",
		Status: &DownloadStatus{
			TotalBytes:     1,
			ProcessedBytes: 1,
		},
	}
	asString := event.String()
	s.Contains(asString, string(event.State))
	s.Contains(asString, event.URL)
	s.Contains(asString, event.FilePath)
	s.Contains(asString, event.TempFilePath)
	s.Contains(asString, fmt.Sprintf("%v%%", event.Status.GetPercentage()))
}
