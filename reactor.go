package swoosh

import (
	"net"

	log "github.com/sirupsen/logrus"
)

type reactor struct {
	eventHandler EventHandler
	logger       *log.Logger

	listener net.Listener
}

func newReactor(ln net.Listener, eventHandler EventHandler,
	l *log.Logger) *reactor {
	el := &reactor{
		eventHandler: eventHandler,
		logger:       l,
		listener:     ln,
	}

	return el
}
