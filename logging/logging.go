// Package logging provides a dead-simple, level-aware logging facility.
package logging

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Logger ist the interface that is implemented by all logging implementations.
type Logger interface {
	// SetLevel sets the log level for this logger instance.
	SetLevel(level Level)
	// Log logs a message in the given log level.
	Log(level Level, format string, v ...interface{})
	// Fatal logs the given arguments and exits the program.
	Fatal(v ...interface{})
}

// Level defines all available log levels.
type Level uint8

// The various log levels.
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var levelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

// LevelFromString parses a string and returns the approriate Level with the
// same name. If no such Level exists, a non-nil error is returned.
func LevelFromString(s string) (*Level, error) {
	s = strings.ToUpper(s)
	for k, v := range levelNames {
		if v == s {
			return &k, nil
		}
	}
	return nil, errors.New("No such level name exists.")
}

func (l Level) String() string {
	return levelNames[l]
}

type stdoutLogger struct {
	level Level
}

func (l *stdoutLogger) Log(level Level, format string, v ...interface{}) {
	if l.level > level {
		return
	}
	prolog := fmt.Sprintf("[%s] ", level.String())
	msg := fmt.Sprintf(format, v...)
	log.Print(prolog, msg)
}

func (l *stdoutLogger) Fatal(v ...interface{}) {
	prolog := fmt.Sprintf("[%s] ", ERROR.String())
	msg := fmt.Sprint(v...)
	log.Print(prolog, msg)
	os.Exit(1)
}

func (l *stdoutLogger) SetLevel(level Level) {
	l.level = level
}

// NewLogger creates and returns a new default Logger.
func NewLogger() Logger {
	return &stdoutLogger{DEBUG}
}

var defaultLogger = NewLogger()

// SetLevel sets the logging level of the default logger.
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// Log logs the given message via the default logger.
func Log(level Level, format string, v ...interface{}) {
	defaultLogger.Log(level, format, v...)
}

// Fatal logs the given message via the default logger.
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}
