package cmdutils

import (
	"os"
	"runtime"

	"github.com/zephinzer/dev/internal/log"
)

// exitLogger defines an interface that allows for printing stuff back
// to the user
type exitLogger func(string, ...interface{})

// exit exits the current process with status code :exitCode, printing the message
// :message using the provided `exitLogger` to log at the appropriate level. if the
// :stackTrace variable is defined and set to `true`, the stack trace of the process
// will be dumped
func exit(print exitLogger, message string, exitCode int, stackTrace ...bool) {
	print(message)
	if len(stackTrace) > 0 && stackTrace[0] {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		print("stack trace follows\n\n%s", string(buf))
	}
	os.Exit(exitCode)
}

// ExitWithError exits the application gracefully with the provided
// exit status :exitCode, printing both message and stack trace at the
// `error` level
func ExitWithError(message string, exitCode int) {
	exit(log.Errorf, message, exitCode, true)
}

// ExitWithProblem exits the application with the provided exit status
// :exitCode, printing the provided message :message before exiting
func ExitWithProblem(message string, exitCode int) {
	exit(log.Errorf, message, exitCode)
}

// ExitWithMessage exits the application using the provided :exitCode,
// printing the message :message at the `info` level before that
func ExitWithMessage(message string, exitCode int) {
	exit(log.Infof, message, exitCode)
}

// ExitWithStackTrace exits the application with exit status 0 while
// printing both message and stack trace at the 'trace' level
func ExitWithStackTrace(message string) {
	exit(log.Tracef, message, 0)
}
