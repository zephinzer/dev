package utils

import (
	"archive/zip"
	"fmt"
	"io"
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

// UnzipEvent is an object that is passed to the events stream
// for the consumer to know what's going on inside
type UnzipEvent struct {
	// State is a string code that indicates the underlying operation
	State UnzipState
	// Path is a string indicating path to a file if it's non-empty
	Path string
	// Message is an arbitrary string
	Message string
	// Status should provide metadata on the underlying operation
	Status *UnzipStatus
}

// UnzipStatus stores the status of the unzipping process and is
// returned through the UnzipEvent object
type UnzipStatus struct {
	// BytesTotal is the total number of bytes to extract
	BytesTotal int64
	// BytesProcessed is the number of bytes processed so far
	BytesProcessed int64
	// FilesTotalCount is the total number of files to extract
	FilesTotalCount int
	// FilesProcessedCount is the number of files processed so far
	FilesProcessedCount int
}

func (us UnzipStatus) GetPercentDoneByBytes() float64 {
	return float64(us.BytesProcessed) / float64(us.BytesTotal) * 100
}

func (us UnzipStatus) GetPercentDoneByFiles() float64 {
	return float64(us.FilesProcessedCount) / float64(us.FilesTotalCount) * 100
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

	if options.Events == nil {
		options.Events = make(chan UnzipEvent, 16)
		go func() {
			for {
				if _, ok := <-options.Events; !ok {
					return
				}
			}
		}()
	}
	defer close(options.Events)

	// configure paths using absolute paths
	pathToZip := options.InputPath
	pathToExtractTo := options.OutputPath

	// configure the reader for zip files
	var zipReader *zip.ReadCloser
	if zipReader, err = zip.OpenReader(pathToZip); err != nil {
		options.Events <- UnzipEvent{UnzipStateError, "", fmt.Sprintf("failed to open zip file at '%s': %s", pathToZip, err), &status}
		return []error{err}
	}
	defer zipReader.Close()

	// prepare for iteration
	errors := []error{}
	for _, file := range zipReader.File {
		status.BytesTotal += file.FileInfo().Size()
	}
	status.FilesTotalCount = len(zipReader.File)
	options.Events <- UnzipEvent{UnzipStateStarting, "", "", &status}

	// go through all files in the zip file
	for _, file := range zipReader.File {
		status.FilesProcessedCount++
		absoluteOutputPath := path.Join(pathToExtractTo, file.Name)
		options.Events <- UnzipEvent{UnzipStateProcessing, absoluteOutputPath, "", &status}
		if file.FileInfo().IsDir() {
			os.MkdirAll(absoluteOutputPath, os.ModePerm)
			options.Events <- UnzipEvent{UnzipStateProcessing, absoluteOutputPath, "created dir", &status}
			continue
		}

		// ensure the directory to the file has been created
		if os.MkdirAll(filepath.Dir(absoluteOutputPath), os.ModePerm); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, err.Error(), &status}
			if options.ReturnOnFileError {
				return []error{err}
			}
			errors = append(errors, err)
			continue
		}

		// open the file for writing (create if doesn't exist)
		outputFile, err := os.OpenFile(absoluteOutputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, err.Error(), &status}
			if options.ReturnOnFileError {
				options.Events <- UnzipEvent{UnzipStateFailed, absoluteOutputPath, err.Error(), &status}
				return []error{err}
			}
			errors = append(errors, err)
			continue
		}
		options.Events <- UnzipEvent{UnzipStateProcessing, absoluteOutputPath, "created/opened file", &status}
		inputFile, err := file.Open()
		if err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, err.Error(), &status}
			if options.ReturnOnFileError {
				options.Events <- UnzipEvent{UnzipStateFailed, absoluteOutputPath, err.Error(), &status}
				return []error{err}
			}
			errors = append(errors, err)
			continue
		}

		size, copyErr := io.Copy(outputFile, inputFile)
		if outputErr := outputFile.Close(); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, outputErr.Error(), &status}
			if options.ReturnOnFileError {
				options.Events <- UnzipEvent{UnzipStateFailed, absoluteOutputPath, outputErr.Error(), &status}
				return []error{outputErr}
			}
			errors = append(errors, outputErr)
			continue
		}
		if inputErr := inputFile.Close(); err != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, inputErr.Error(), &status}
			if options.ReturnOnFileError {
				options.Events <- UnzipEvent{UnzipStateFailed, absoluteOutputPath, inputErr.Error(), &status}
				return []error{inputErr}
			}
			errors = append(errors, inputErr)
			continue
		}

		if copyErr != nil {
			options.Events <- UnzipEvent{UnzipStateError, absoluteOutputPath, copyErr.Error(), &status}
			if options.ReturnOnFileError {
				options.Events <- UnzipEvent{UnzipStateFailed, absoluteOutputPath, copyErr.Error(), &status}
				return []error{copyErr}
			}
			errors = append(errors, copyErr)
			continue
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
