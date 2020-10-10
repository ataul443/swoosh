// Copyright (c) 2020 Shekh Ataul
// This code is licensed under MIT license (see LICENSE for details)

package swoosh

import (
	log "github.com/sirupsen/logrus"
)

type Swoosh struct {
	eventHandler EventHandler
	logger       *log.Logger

	eventLoop *reactor
}

func ListenAndServe(network, address string,
	eventHandler EventHandler) (*Swoosh, error) {
	logger := log.New()

	s := &Swoosh{
		eventHandler: eventHandler,
		logger:       logger,
		eventLoop:    newReactor(eventHandler, logger),
	}
	return s, nil
}

// EnableLog sets log level on the swoosh listener with supplied
// valid swoosh log level. If unknown log level passed it will
// set the log level to FATAL_LEVEL.
func (s *Swoosh) EnableLog(level int) {
	logCallerPrettyfier(s.logger)
	s.setLogLevel(level)

	if level == TraceLevel || level == DebugLevel {
		s.logger.SetReportCaller(true)
		logCallerPrettyfier(s.logger)
	}
}

// GetLogLevel returns current log level of swoosh listener.
func (s *Swoosh) GetLogLevel() int {
	return getSwooshLevel(s.logger.GetLevel())
}
