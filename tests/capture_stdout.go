package tests

import (
	"bytes"
	"io"
	"os"
)

// CaptureStdout captures prints to the standard output from the function
// :fromThis. :fromThis should return an error to indicate whether the function
// call was a success; if the error is non-empty, the returned string will be
// empty (but an empty string with an error indicates a blank input)
func CaptureStdout(fromThis func() error) (string, error) {
	originalStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	defer func() {
		os.Stdout = originalStdout
	}()
	err := fromThis()
	if err != nil {
		return "", err
	}
	output := make(chan string)
	defer close(output)
	go func() {
		var b bytes.Buffer
		io.Copy(&b, reader)
		output <- b.String()
	}()
	writer.Close()
	return <-output, nil
}
