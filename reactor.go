package swoosh

import (
	"net"
)

type reactor struct {
	eventHandler EventHandler

	listener net.Listener
}

func newReactor(ln net.Listener, eventHandler EventHandler) *reactor {
	el := &reactor{
		eventHandler: eventHandler,
		listener:     ln,
	}

	return el
}
