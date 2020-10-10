// Copyright (c) 2020 Shekh Ataul
// This code is licensed under MIT license (see LICENSE for details)

package swoosh

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	// Log level for fine grained trace level messages.
	TRACE_LEVEL = iota

	// Log level for fine grained debug level messages.
	DEBUG_LEVEL

	//_LEVEL Log level for info level messages.
	INFO_LEVEL

	// Log level for error level messages.
	ERROR_LEVEL

	// Log level for warning level messages.
	WARN_LEVEL

	// Log level for fatal level messages.
	FATAL_LEVEL
)

func logCallerPrettyfier(logger *log.Logger) {
	logger.Formatter = &log.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (funcName string, file string) {
			filename := path.Base(f.File)
			funcName = f.Function[(strings.LastIndex(f.Function, "/") + 1):]

			return fmt.Sprintf("%s()", funcName), fmt.Sprintf("%s:%d",
				filename, f.Line)
		},
	}
}

func (s *Swoosh) setLogLevel(level int) {
	if s == nil {
		return
	}

	innerLogLevel := getLogrusLevel(level)
	s.logger.SetLevel(innerLogLevel)
}

func getLogrusLevel(level int) log.Level {
	var innerLogLevel log.Level
	switch level {
	case TRACE_LEVEL:
		innerLogLevel = log.TraceLevel

	case DEBUG_LEVEL:
		innerLogLevel = log.DebugLevel

	case INFO_LEVEL:
		innerLogLevel = log.InfoLevel

	case ERROR_LEVEL:
		innerLogLevel = log.ErrorLevel

	case WARN_LEVEL:
		innerLogLevel = log.WarnLevel

	default:
		innerLogLevel = log.FatalLevel
	}

	return innerLogLevel
}

func getSwooshLevel(level log.Level) int {
	var innerSwooshLevel int
	switch level {
	case log.TraceLevel:
		innerSwooshLevel = TRACE_LEVEL

	case log.DebugLevel:
		innerSwooshLevel = DEBUG_LEVEL

	case log.InfoLevel:
		innerSwooshLevel = INFO_LEVEL

	case log.ErrorLevel:
		innerSwooshLevel = ERROR_LEVEL

	case log.WarnLevel:
		innerSwooshLevel = WARN_LEVEL

	default:
		innerSwooshLevel = FATAL_LEVEL
	}

	return innerSwooshLevel
}
