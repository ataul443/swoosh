// Copyright (c) 2020 Shekh Ataul
// This code is licensed under MIT license (see LICENSE for details)

package swoosh

import (
	"net"

	log "github.com/sirupsen/logrus"
)

type Swoosh struct {
	eventHandler EventHandler
	eventLoop    *reactor

	// do not use it after calling eventLoop.run().
	// eventLoop will close it as needed.
	stdListener net.Listener
}

func Listen(network, address string, eh EventHandler) (*Swoosh, error) {
	ln, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	el, err := newReactor(ln, eh)
	if err != nil {
		return nil, err
	}

	s := &Swoosh{
		eventHandler: eh,
		stdListener:  ln,
		eventLoop:    el,
	}

	return s, nil
}

// Serve will start serving new connections.
func (s *Swoosh) Serve() error {
	return s.eventLoop.run()
}

// EnableLog sets log level on the swoosh listener with supplied
// valid swoosh log level. If unknown log level passed it will
// set the log level to FATAL_LEVEL.
func (s *Swoosh) EnableLog(level int) {
	log.SetLevel(getLogrusLevel(level))
	if level == TraceLevel || level == DebugLevel {
		log.SetReportCaller(true)
		log.SetFormatter(logCallerPrettyfier())
	}
}

// GetLogLevel returns current log level of swoosh listener.
func (s *Swoosh) GetLogLevel() int {
	return getSwooshLevel(log.GetLevel())
}
