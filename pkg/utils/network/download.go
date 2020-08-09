package network

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type DownloadState string

const (
	DefaultEventsUpdateInterval               = 500 * time.Millisecond
	DownloadStateStarting       DownloadState = "download_starting"
	DownloadStateReport         DownloadState = "download_report"
	DownloadStateError          DownloadState = "download_error"
	DownloadStateFailed         DownloadState = "download_failed"
	DownloadStateSuccess        DownloadState = "download_success"
)

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
	var eventsUpdateDuration time.Duration
	if options.EventsUpdateInterval == eventsUpdateDuration {
		eventsUpdateDuration = DefaultEventsUpdateInterval
	}
	ticker := time.NewTicker(eventsUpdateDuration)
	defer ticker.Stop()
	go func(tick <-chan time.Time) {
		defer func() {
			recover() // when channel is closed, a panic will happen
		}()
		for {
			<-tick
			options.Events <- DownloadEvent{DownloadStateReport, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
		}
	}(ticker.C)
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

	// open output file for writing (create if it doesn't exist)
	outputFile, err := os.OpenFile(tmpFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	options.Events <- DownloadEvent{DownloadStateStarting, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
	response, err := http.Get(options.URL)
	if err != nil {
		options.Events <- DownloadEvent{DownloadStateFailed, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
		return err
	}
	defer response.Body.Close()
	contentLength, err := strconv.Atoi(response.Header.Get("Content-Length"))
	if err != nil {
		options.Events <- DownloadEvent{DownloadStateFailed, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
		return err
	}
	downloadStatus.TotalBytes = uint64(contentLength)
	if _, err := io.Copy(outputFile, io.TeeReader(response.Body, &downloadStatus)); err != nil {
		options.Events <- DownloadEvent{DownloadStateFailed, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
		return err
	}

	if err := outputFile.Close(); err != nil {
		options.Events <- DownloadEvent{DownloadStateFailed, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
		return err
	}
	if err := os.Rename(tmpFilePath, options.FilePath); err != nil {
		options.Events <- DownloadEvent{DownloadStateFailed, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
		return err
	}
	options.Events <- DownloadEvent{DownloadStateSuccess, options.URL, options.FilePath, tmpFilePath, &downloadStatus}
	return nil
}
