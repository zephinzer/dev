package software

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TypesTests struct {
	suite.Suite
}

func TestTypes(t *testing.T) {
	suite.Run(t, &TypesTests{})
}

func (s *TypesTests) TestCheck_BeforeRun() {
	check := Check{
		Command: []string{"test"},
	}
	s.Nil(check.observed)
}

func (s *TypesTests) TestCheck_Run() {
	check := Check{
		Command: []string{"test"},
	}
	checkErr := check.Run()
	s.Nil(checkErr)
	observed := check.GetObserved()
	s.Equal(1, observed.ExitCode)
}

func (s *TypesTests) TestCheck_Run_withArguments_ok() {
	check := Check{
		Command: []string{"test", "0", "=", "0"},
	}
	checkErr := check.Run()
	s.Nil(checkErr)
	observed := check.GetObserved()
	s.Equal(0, observed.ExitCode)
}

func (s *TypesTests) TestCheck_Run_withArguments_exitCodeNonZero() {
	check := Check{
		Command: []string{"test", "1", "=", "0"},
	}
	checkErr := check.Run()
	s.Nil(checkErr)
	observed := check.GetObserved()
	s.Equal(1, observed.ExitCode)
}

func (s *TypesTests) TestCheck_Run_givesStdoutLogs() {
	check := Check{
		Command: []string{"sh", "-c", "echo 'hello world'"},
	}
	checkErr := check.Run()
	s.Nil(checkErr)
	observed := check.GetObserved()
	s.Equal(0, observed.ExitCode)
	s.Equal("hello world\n", *observed.Stdout)
}

func (s *TypesTests) TestCheck_Run_givesStderrLogs() {
	check := Check{
		Command: []string{"sh", "-c", "echo 'hello world' >&2"},
	}
	checkErr := check.Run()
	s.Nil(checkErr)
	observed := check.GetObserved()
	s.Equal(0, observed.ExitCode)
	s.Equal("hello world\n", *observed.Stderr)
}

func (s *TypesTests) TestCheck_Verify() {
	expectedStdout := "stdout"
	expectedStderr := "stderr"
	check := Check{
		Command:  []string{"test"},
		ExitCode: 0,
		Stderr:   &expectedStderr,
		Stdout:   &expectedStdout,
	}
	check.observed = &check
	s.Nil(check.Verify())
}

func (s *TypesTests) TestCheck_Verify_ensuresRunHasBeenCalled() {
	expectedStdout := "stdout"
	expectedStderr := "stderr"
	check := Check{
		Command:  []string{"test"},
		ExitCode: 0,
		Stderr:   &expectedStderr,
		Stdout:   &expectedStdout,
	}
	s.NotNil(check.Verify())
}

func (s *TypesTests) TestCheck_Verify_ensuresExitCodeMatches() {
	check := Check{
		ExitCode: 0,
		observed: &Check{
			Command:  []string{},
			ExitCode: 1,
		},
	}
	s.NotNil(check.Verify())
}

func (s *TypesTests) TestCheck_Verify_ensuresStdoutMatches() {
	expectedStdout := "stdout"
	unexpectedStdout := "__stdout"
	check := Check{
		Command:  []string{"test"},
		ExitCode: 0,
		Stdout:   &expectedStdout,
		observed: &Check{
			Stdout: &unexpectedStdout,
		},
	}
	s.NotNil(check.Verify())
}

func (s *TypesTests) TestCheck_Verify_ensuresStderrMatches() {
	expectedStderr := "stderr"
	unexpectedStderr := "__stderr"
	check := Check{
		Command:  []string{"test"},
		ExitCode: 0,
		Stderr:   &expectedStderr,
		observed: &Check{
			Stderr: &unexpectedStderr,
		},
	}
	s.NotNil(check.Verify())
}
