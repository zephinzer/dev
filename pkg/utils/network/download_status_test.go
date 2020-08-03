package network

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DownloadStatusTests struct {
	suite.Suite
}

func TestDownloadStatus(t *testing.T) {
	suite.Run(t, &DownloadStatusTests{})
}

func (s *DownloadStatusTests) Test_GetPercentage() {
	downloadStatus := DownloadStatus{1, 1}
	s.EqualValues(100, downloadStatus.GetPercentage())
	downloadStatus = DownloadStatus{2, 1}
	s.EqualValues(50, downloadStatus.GetPercentage())
	downloadStatus = DownloadStatus{10, 1}
	s.EqualValues(10, downloadStatus.GetPercentage())
	downloadStatus = DownloadStatus{100, 1}
	s.EqualValues(1, downloadStatus.GetPercentage())
	downloadStatus = DownloadStatus{1000, 1}
	s.EqualValues(0.1, downloadStatus.GetPercentage())
}

func (s *DownloadStatusTests) Test_Write() {
	downloadStatus := DownloadStatus{100, 1}
	s.EqualValues(1, downloadStatus.GetPercentage())
	contentLength, err := downloadStatus.Write([]byte("123456789"))
	s.Nil(err)
	s.EqualValues(9, contentLength)
	s.EqualValues(10, downloadStatus.GetPercentage())
}
