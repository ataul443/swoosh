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

	var innerLogLevel log.Level
	switch level {
	case TRACE_LEVEL:
		innerLogLevel = log.TraceLevel
		s.logger.SetReportCaller(true)

	case DEBUG_LEVEL:
		innerLogLevel = log.DebugLevel
		s.logger.SetReportCaller(true)

	case INFO_LEVEL:
		innerLogLevel = log.InfoLevel

	case ERROR_LEVEL:
		innerLogLevel = log.ErrorLevel

	case WARNING_LEVEL:
		innerLogLevel = log.WarnLevel

	default:
		innerLogLevel = log.FatalLevel
	}

	s.logger.SetLevel(innerLogLevel)
	s.logger.Tracef("log level set to %s", innerLogLevel)
}
