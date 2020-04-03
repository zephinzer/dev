package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

type UntarState string

const (
	UntarStateStarting   UntarState = "untar_starting"
	UntarStateProcessing UntarState = "untar_processing"
	UntarStateError      UntarState = "untar_error"
	UntarStateOK         UntarState = "untar_ok"
	UntarStateStatus     UntarState = "untar_status"
	UntarStateSuccess    UntarState = "untar_success"
	UntarStateFailed     UntarState = "untar_failed"
)

// UntarEvent is an object that is passed to the events stream
// for the consumer to know what's going on inside
type UntarEvent struct {
	// State is a string code that indicates the underlying operation
	State UntarState
	// Path is a string indicating path to a file if it's non-empty
	Path string
	// Message is an arbitrary string
	Message string
	// Status should provide metadata on the underlying operation
	Status *UntarStatus
}

// UntarStatus stores the status of the untarring process and is
// returned through the UntarEvent object
type UntarStatus struct {
	// BytesTotal is the total number of bytes to extract
	BytesTotal int64
	// BytesProcessed is the number of bytes processed so far
	BytesProcessed int64
	// FilesTotalCount is the total number of files to extract
	FilesTotalCount int
	// FilesProcessedCount is the number of files processed so far
	FilesProcessedCount int
}

func (us UntarStatus) GetPercentDoneByBytes() float64 {
	return float64(us.BytesProcessed) / float64(us.BytesTotal) * 100
}

func (us UntarStatus) GetPercentDoneByFiles() float64 {
	return float64(us.FilesProcessedCount) / float64(us.FilesTotalCount) * 100
}

type UntarOptions struct {
	// Events, if populated, receives events for logging purposes
	Events chan UntarEvent
	// InputPath defines the path to the .zip file we want to untar
	InputPath string
	// OutputPath defines the path to a directory where the untarred files
	// should go
	OutputPath string
	// ReturnOnFileError indicates whether to return an error instantly
	// when an error is encountered
	ReturnOnFileError bool
}

// Untar takes the .zip file located at options.InputPath and untars its contents
// to options.OutputPath; returns an error if something unexpected happened, or
// nil if all is well
//
// Made possible with guidance from: https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
func Untar(options UntarOptions) []error {
	var err error
	status := UntarStatus{}

	if options.Events == nil {
		options.Events = make(chan UntarEvent, 16)
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
	pathToTar := options.InputPath
	if !path.IsAbs(options.InputPath) {
		if pathToTar, err = convertPathToAbsolute(options.InputPath); err != nil {
			options.Events <- UntarEvent{UntarStateFailed, "", fmt.Sprintf("failed to retrieve absolute path from '%s': %s", options.InputPath, err), &status}
			return []error{err}
		}
	}
	pathToExtractTo := options.OutputPath
	if !path.IsAbs(options.OutputPath) {
		if pathToExtractTo, err = convertPathToAbsolute(options.OutputPath); err != nil {
			options.Events <- UntarEvent{UntarStateFailed, "", fmt.Sprintf("failed to retrieve absolute path from '%s': %s", options.OutputPath, err), &status}
			return []error{err}
		}
	}

	// configure the reader for zip files
	tarFile, err := os.OpenFile(pathToTar, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return []error{err}
	}
	defer tarFile.Close()

	gzipReader, err := gzip.NewReader(tarFile)
	if err != nil {
		options.Events <- UntarEvent{UntarStateFailed, "", fmt.Sprintf("failed to open tar file at '%s': %s", pathToTar, err), &status}
		return []error{err}
	}
	defer gzipReader.Close()

	// prepare for iteration
	errors := []error{}
	tarReader := tar.NewReader(gzipReader)
	for {
		tarFile, err := tarReader.Next()
		shouldBreak := false
		switch {
		case err == io.EOF:
			shouldBreak = true
		case err != nil:
			return []error{err}
		case tarFile == nil:
			continue
		}
		if shouldBreak {
			break
		}
		log.Println("??? 1", tarFile.Name)
		switch tarFile.Typeflag {
		case tar.TypeReg:
			status.BytesTotal += tarFile.Size
			status.FilesTotalCount += 1
		}
	}
	options.Events <- UntarEvent{UntarStateStarting, "", "processed status", &status}

	// reset the reader
	tarFile.Seek(0, 0)
	gzipReader.Reset(tarFile)
	tarReader = tar.NewReader(gzipReader)

	// go through all files in the tar file
	for {
		tarFile, err := tarReader.Next()
		shouldBreak := false
		switch {
		case err == io.EOF:
			shouldBreak = true
		case err != nil:
			return []error{err}
		case tarFile == nil:
			continue
		}
		if shouldBreak {
			break
		}
		log.Println("??? 2", tarFile.Name)
		absoluteOutputPath := path.Join(pathToExtractTo, tarFile.Name)
		options.Events <- UntarEvent{UntarStateProcessing, absoluteOutputPath, "", &status}
		switch tarFile.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(absoluteOutputPath, os.ModePerm)
			options.Events <- UntarEvent{UntarStateProcessing, absoluteOutputPath, "created dir", &status}
			continue
		case tar.TypeReg:
			// ensure the directory to the file has been created
			if os.MkdirAll(filepath.Dir(absoluteOutputPath), os.ModePerm); err != nil {
				options.Events <- UntarEvent{UntarStateError, absoluteOutputPath, err.Error(), &status}
				if options.ReturnOnFileError {
					return []error{err}
				}
				errors = append(errors, err)
				continue
			}
			// open the file for writing (create if doesn't exist)
			outputFile, err := os.OpenFile(absoluteOutputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
			if err != nil {
				options.Events <- UntarEvent{UntarStateError, absoluteOutputPath, err.Error(), &status}
				if options.ReturnOnFileError {
					options.Events <- UntarEvent{UntarStateFailed, absoluteOutputPath, err.Error(), &status}
					return []error{err}
				}
				errors = append(errors, err)
				continue
			}
			options.Events <- UntarEvent{UntarStateProcessing, absoluteOutputPath, "created/opened file", &status}

			size, copyErr := io.Copy(outputFile, tarReader)
			if outputErr := outputFile.Close(); err != nil {
				options.Events <- UntarEvent{UntarStateError, absoluteOutputPath, outputErr.Error(), &status}
				if options.ReturnOnFileError {
					options.Events <- UntarEvent{UntarStateFailed, absoluteOutputPath, outputErr.Error(), &status}
					return []error{outputErr}
				}
				errors = append(errors, outputErr)
				continue
			}

			if copyErr != nil {
				options.Events <- UntarEvent{UntarStateError, absoluteOutputPath, copyErr.Error(), &status}
				if options.ReturnOnFileError {
					options.Events <- UntarEvent{UntarStateFailed, absoluteOutputPath, copyErr.Error(), &status}
					return []error{copyErr}
				}
				errors = append(errors, copyErr)
				continue
			}
			status.BytesProcessed += size
			status.FilesProcessedCount++
			options.Events <- UntarEvent{UntarStateOK, absoluteOutputPath, fmt.Sprintf("extracted %v bytes", size), &status}
		}
	}

	if len(errors) > 0 {
		options.Events <- UntarEvent{UntarStateFailed, "", "", &status}
		return errors
	}
	options.Events <- UntarEvent{UntarStateSuccess, "", "", &status}
	return nil
}
