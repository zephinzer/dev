package cmdutils

import (
	"os"
	"runtime"

	"github.com/zephinzer/dev/internal/log"
)

// exitLogger defines an interface that allows for printing stuff back
// to the user
type exitLogger func(string, ...interface{})

// exitWithStackTrace exits the current process with status code :exitCode, printing
// the message :message using the provided `exitLogger` to log at the appropriate level
func exitWithStackTrace(print exitLogger, message string, exitCode int) {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	print("exiting with message '%s', stack trace follows\n\n%s", message, string(buf))
	os.Exit(exitCode)
}

// ExitWithError exits the application gracefully with the provided
// exit status :exitCode, printing both message and stack trace at the
// `error` level
func ExitWithError(message string, exitCode int) {
	exitWithStackTrace(log.Errorf, message, exitCode)
}

// ExitWithStackTrace exits the application gracefully with exit
// status 0 while printing both message and stack trace at the 'trace'
// level
func ExitWithStackTrace(message string) {
	exitWithStackTrace(log.Tracef, message, 0)
}
