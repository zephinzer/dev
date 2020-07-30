package network

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type DownloadTests struct {
	FilePath string
	suite.Suite
}

func TestDownload(t *testing.T) {
	suite.Run(t, &DownloadTests{})
}

func (s *DownloadTests) SetupTest() {
	s.FilePath = "./tests/downloads/" + time.Now().Format("download_test_20060102150405")
}

func (s *DownloadTests) TestDownload() {
	var receivedEventTypes []DownloadState
	var receivedStatuses []DownloadStatus
	var waiter sync.WaitGroup
	events := make(chan DownloadEvent, 16)
	go func() {
		for {
			if event, ok := <-events; ok {
				receivedEventTypes = append(receivedEventTypes, event.State)
				receivedStatuses = append(receivedStatuses, *event.Status)
			} else {
				waiter.Done()
				return
			}
		}
	}()
	waiter.Add(1)
	err := Download(DownloadOptions{
		Events:   events,
		URL:      "https://github.com/kubernetes/kubectl/archive/v0.19.0-alpha.1.zip",
		FilePath: s.FilePath,
	})
	waiter.Wait()
	s.Contains(receivedEventTypes, DownloadStateStarting, "%v should have contained %s", receivedEventTypes, DownloadStateStarting)
	if err == nil {
		s.Contains(receivedEventTypes, DownloadStateReport, "%v should have contained %s", receivedEventTypes, DownloadStateReport)
		s.Contains(receivedEventTypes, DownloadStateSuccess, "%v should have contained %s", receivedEventTypes, DownloadStateSuccess)
	}
}
