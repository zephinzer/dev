package software

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
)

// Check represents a functional check on whether the software exists.
// It works by running the command as specified in Command and comparing
// the output with the provided ExitCode, Stdout, and Stderr.
//
// Usage:
// 1. Use `.Run()` method to run the check, results will be stored in the `.observed` property
// 1. Use `.Verify()` method to verify that the `.observed` property matches the provided parameters
type Check struct {
	// Command is the command to run including arguments
	Command []string `json:"command" yaml:"command"`
	// ExitCode is the expected exit code we can expect from the command running if defined
	ExitCode int `json:"exitCode" yaml:"exitCode"`
	// Stdout is the expected output on stdout if defined
	Stdout *string `json:"stdout" yaml:"stdout"`
	// Stderr is the expected output on stderr if defined
	Stderr *string `json:"stderr" yaml:"stderr"`
	// observed contains the observed value after Run is executed
	observed *Check
}

// GetObserved returns the observed results, a nil indicates the check has
// never been run before
func (check Check) GetObserved() *Check {
	return check.observed
}

// Run executes the check, use check.Verify() to verify the results, if this
// function returns an error, it's a system-type error meaning it has nothing
// to do with the command itself, but with it's invocation on the host system
func (check *Check) Run() error {
	var command *exec.Cmd
	var stderr, stdout bytes.Buffer
	if len(check.Command) == 0 {
		return fmt.Errorf("no command specified")
	} else if len(check.Command) == 1 {
		command = exec.Command(check.Command[0])
	} else {
		command = exec.Command(check.Command[0], check.Command[1:]...)
	}
	command.Stdout = &stdout
	command.Stderr = &stderr
	if err := command.Start(); err != nil {
		return err
	}
	check.observed = &Check{
		Command: append([]string{"echo"}, strings.Join(check.Command, " ")),
	}
	if err := command.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				check.observed.ExitCode = status.ExitStatus()
			}
		} else {
			return err
		}
	}
	stderrOutput := strings.Trim(stderr.String(), "\r\n\t ")
	check.observed.Stderr = &stderrOutput
	stdoutOutput := strings.Trim(stdout.String(), "\r\n\t ")
	check.observed.Stdout = &stdoutOutput
	return nil
}

// Verify returns nil on veritfication succeeded, returns an error instance containing
// the reason why verification failed otherwise
func (check Check) Verify() error {
	var stdoutMatcher, stderrMatcher *regexp.Regexp
	if check.Stdout != nil {
		stdoutMatcher = regexp.MustCompile(*check.Stdout)
	}
	if check.Stderr != nil {
		stderrMatcher = regexp.MustCompile(*check.Stderr)
	}
	switch true {
	case check.observed == nil:
		return fmt.Errorf("Run() has not been called on this Check{} instance (Command: [%s])", strings.Join(check.Command, ","))
	case check.ExitCode != check.observed.ExitCode:
		return fmt.Errorf("expected exit code %v did not match actual exit code %v", check.ExitCode, check.observed.ExitCode)
	case check.Stdout != nil && !stdoutMatcher.Match([]byte(*check.observed.Stdout)):
		return fmt.Errorf("expected stdout '%s' did not match actual stdout '%s'", *check.Stdout, *check.observed.Stdout)
	case check.Stderr != nil && !stderrMatcher.Match([]byte(*check.observed.Stderr)):
		return fmt.Errorf("expected stderr '%s' did not match actual stderr '%s'", *check.Stderr, *check.observed.Stderr)
	}
	return nil
}
