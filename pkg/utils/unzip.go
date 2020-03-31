package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

type UnzipState string

const (
	UnzipStateStarting   UnzipState = "unzip_starting"
	UnzipStateProcessing UnzipState = "unzip_processing"
	UnzipStateError      UnzipState = "unzip_error"
	UnzipStateOK         UnzipState = "unzip_ok"
	UnzipStateStatus     UnzipState = "unzip_status"
	UnzipStateSuccess    UnzipState = "unzip_success"
	UnzipStateFailed     UnzipState = "unzip_failed"
)

type UnzipEvent struct {
	State   UnzipState
	Path    string
	Message string
	Status  *UnzipStatus
}

type UnzipStatus struct {
	BytesTotal          int64
	BytesProcessed      int64
	FilesTotalCount     int
	FilesProcessedCount int
}

type UnzipOptions struct {
	// Events, if populated, receives events for logging purposes
	Events chan UnzipEvent
	// InputPath defines the path to the .zip file we want to unzip
	InputPath string
	// OutputPath defines the path to a directory where the unzipped files
	// should go
	OutputPath string
	// ReturnOnFileError indicates whether to return an error instantly
	// when an error is encountered
	ReturnOnFileError bool
}

// Unzip takes the .zip file located at options.InputPath and unzips its contents
// to options.OutputPath; returns an error if something unexpected happened, or
// nil if all is well
//
// Made possible with guidance from: https://golangcode.com/unzip-files-in-go/
func Unzip(options UnzipOptions) []error {
	var err error
	status := UnzipStatus{}

	stopProcessingEvents := make(chan struct{}, 1)
	if options.Events == nil {
		options.Events = make(chan UnzipEvent, 16)
		go func() {
			for {
				e, ok := <-options.Events
				if !ok {
					stopProcessingEvents <- struct{}{}
					return
				}
				percentDone := int(float64(e.Status.BytesProcessed) / float64(e.Status.BytesTotal) * 100)
				log.Printf("%v%% done [%v/%v files processed]", percentDone, e.Status.FilesProcessedCount, e.Status.FilesTotalCount)
			}
		}()
	}
	defer close(options.Events)

	pathToZip := options.InputPath
	if !path.IsAbs(options.InputPath) {
		var err error
		if pathToZip, err = convertPathToAbsolute(options.InputPath); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, "", err.Error(), &status}
			return []error{err}
		}
	}

	var zipReader *zip.ReadCloser
	if zipReader, err = zip.OpenReader(pathToZip); err != nil {
		options.Events <- UnzipEvent{UnzipStateError, "", err.Error(), &status}
		return []error{err}
	}
	defer zipReader.Close()

	pathToExtractTo := options.OutputPath
	if !path.IsAbs(options.OutputPath) {
		var err error
		if pathToExtractTo, err = convertPathToAbsolute(options.OutputPath); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, "", err.Error(), &status}
			return []error{err}
		}
	}

	errors := []error{}
	for _, file := range zipReader.File {
		status.BytesTotal += file.FileInfo().Size()
	}
	status.FilesTotalCount = len(zipReader.File)
	options.Events <- UnzipEvent{UnzipStateStarting, "", "", &status}
	for _, file := range zipReader.File {
		status.FilesProcessedCount++
		absoluteOutputPath := path.Join(pathToExtractTo, file.Name)
		options.Events <- UnzipEvent{UnzipStateProcessing, absoluteOutputPath, "", &status}
		if file.FileInfo().IsDir() {
			os.MkdirAll(absoluteOutputPath, os.ModePerm)
			options.Events <- UnzipEvent{UnzipStateProcessing, absoluteOutputPath, "created dir", &status}
			continue
		}
		if os.MkdirAll(filepath.Dir(absoluteOutputPath), os.ModePerm); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, err.Error(), &status}
			if options.ReturnOnFileError {
				return []error{err}
			} else {
				errors = append(errors, err)
				continue
			}
		}
		outputFile, err := os.OpenFile(absoluteOutputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, err.Error(), &status}
			if options.ReturnOnFileError {
				return []error{err}
			} else {
				errors = append(errors, err)
				continue
			}
		}
		options.Events <- UnzipEvent{UnzipStateProcessing, absoluteOutputPath, "created/opened file", &status}
		inputFile, err := file.Open()
		if err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, err.Error(), &status}
			if options.ReturnOnFileError {
				return []error{err}
			} else {
				errors = append(errors, err)
				continue
			}
		}

		size, copyErr := io.Copy(outputFile, inputFile)
		if outputErr := outputFile.Close(); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, outputErr.Error(), &status}
			if options.ReturnOnFileError {
				return []error{outputErr}
			} else {
				errors = append(errors, outputErr)
				continue
			}
		}
		if inputErr := inputFile.Close(); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, inputErr.Error(), &status}
			if options.ReturnOnFileError {
				return []error{inputErr}
			} else {
				errors = append(errors, inputErr)
				continue
			}
		}

		if copyErr != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, copyErr.Error(), &status}
			if options.ReturnOnFileError {
				return []error{copyErr}
			} else {
				errors = append(errors, copyErr)
				continue
			}
		}
		status.BytesProcessed += size
		options.Events <- UnzipEvent{UnzipStateOK, absoluteOutputPath, fmt.Sprintf("extracted %v bytes", size), &status}
	}

	if len(errors) > 0 {
		options.Events <- UnzipEvent{UnzipStateFailed, "", "", &status}
		return errors
	}
	options.Events <- UnzipEvent{UnzipStateSuccess, "", "", &status}
	return nil
}
