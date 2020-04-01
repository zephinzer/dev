package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type DownloadEvent struct {
	URL          string
	FilePath     string
	TempFilePath string
	Status       *DownloadStatus
}

type DownloadStatus struct {
	TotalBytes     uint64
	ProcessedBytes uint64
}

// Write implements io.Writer
func (ds *DownloadStatus) Write(content []byte) (int, error) {
	contentLength := len(content)
	ds.ProcessedBytes += uint64(contentLength)
	return contentLength, nil
}

type DownloadOptions struct {
	Events               chan DownloadEvent
	EventsUpdateInterval time.Duration
	FilePath             string
	URL                  string
	Overwrite            bool
}

// Download downloads a file from the url in options.URL and places the
// content in the file at options.FilePath
//
// Made possible with guidance from: https://golangcode.com/download-a-file-with-progress/
func Download(options DownloadOptions) error {
	downloadStatus := DownloadStatus{}
	tmpFilePath := options.FilePath + ".download_" + time.Now().Format("20060102150405")
	if options.Events == nil {
		options.Events = make(chan DownloadEvent, 16)
		go func() {
			for {
				if _, ok := <-options.Events; !ok {
					return
				}
			}
		}()
	}
	defer close(options.Events)

	if err := os.MkdirAll(filepath.Dir(options.FilePath), os.ModePerm); err != nil {
		return err
	}

	// check that intended output file does not exist
	_, err := os.Stat(options.FilePath)
	if err == nil && !options.Overwrite {
		return fmt.Errorf("file at '%s' already exists", options.FilePath)
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	outputFile, err := os.OpenFile(tmpFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	response, err := http.Get(options.URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	contentLength, err := strconv.Atoi(response.Header.Get("Content-Length"))
	if err != nil {
		return err
	}
	downloadStatus.TotalBytes = uint64(contentLength)
	if _, err := io.Copy(outputFile, io.TeeReader(response.Body, &downloadStatus)); err != nil {
		return err
	}

	if err := outputFile.Close(); err != nil {
		return err
	}
	if err := os.Rename(tmpFilePath, options.FilePath); err != nil {
		return err
	}
	return nil
}
