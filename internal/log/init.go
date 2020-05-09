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

type Level logger.Level

var (
	DefaultLevel Level = Level(logger.LevelInfo)
	LevelDebug   Level = Level(logger.LevelDebug)
	LevelTrace   Level = Level(logger.LevelTrace)
)

func Init(level ...Level) {
	logLevel := DefaultLevel
	if len(level) > 0 {
		logLevel = level[0]
	}
	instance = logger.New(logger.Options{
		Level:        logger.Level(logLevel),
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
