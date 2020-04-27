package log

import (
	"fmt"
	"os"

	"github.com/usvc/go-logger"
)

var (
	instance logger.Logger
	Trace    func(...interface{})
	Tracef   func(string, ...interface{})
	Debug    func(...interface{})
	Debugf   func(string, ...interface{})
	Info     func(...interface{})
	Infof    func(string, ...interface{})
	Warn     func(...interface{})
	Warnf    func(string, ...interface{})
	Error    func(...interface{})
	Errorf   func(string, ...interface{})
	Print    func(...interface{}) (int, error)
	Printf   func(string, ...interface{}) (int, error)
)

func Init() {
	instance = logger.New(logger.Options{
		Output:       logger.OutputStderr,
		OutputStream: os.Stderr,
	})
	Trace = instance.Trace
	Tracef = instance.Tracef
	Debug = instance.Debug
	Debugf = instance.Debugf
	Info = instance.Info
	Infof = instance.Infof
	Warn = instance.Warn
	Warnf = instance.Warnf
	Error = instance.Error
	Errorf = instance.Errorf
	Print = fmt.Print
	Printf = fmt.Printf
}
