package swoosh

import log "github.com/sirupsen/logrus"

type reactor struct {
	eventHandler EventHandler
	logger       *log.Logger
}

func newReactor(eventHandler EventHandler, l *log.Logger) *reactor {
	el := &reactor{
		eventHandler: eventHandler,
		logger:       l,
	}

	return el
}
