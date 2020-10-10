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
	TraceLevel = iota

	// Log level for fine grained debug level messages.
	DebugLevel

	// LEVEL Log level for info level messages.
	InfoLevel

	// Log level for error level messages.
	ErrorLevel

	// Log level for warning level messages.
	WarnLevel

	// Log level for fatal level messages.
	FatalLevel
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
	case TraceLevel:
		innerLogLevel = log.TraceLevel

	case DebugLevel:
		innerLogLevel = log.DebugLevel

	case InfoLevel:
		innerLogLevel = log.InfoLevel

	case ErrorLevel:
		innerLogLevel = log.ErrorLevel

	case WarnLevel:
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
		innerSwooshLevel = TraceLevel

	case log.DebugLevel:
		innerSwooshLevel = DebugLevel

	case log.InfoLevel:
		innerSwooshLevel = InfoLevel

	case log.ErrorLevel:
		innerSwooshLevel = ErrorLevel

	case log.WarnLevel:
		innerSwooshLevel = WarnLevel

	default:
		innerSwooshLevel = FatalLevel
	}

	return innerSwooshLevel
}
