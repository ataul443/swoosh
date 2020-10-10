package swoosh

import (
	"errors"
	"net"

	log "github.com/sirupsen/logrus"
)

type reactor struct {
	eventHandler EventHandler

	listener   net.Listener
	listenerFD int
}

func newReactor(ln net.Listener, eventHandler EventHandler) *reactor {
	el := &reactor{
		eventHandler: eventHandler,
		listener:     ln,
	}

	return el
}

func detachFDFromListener(ln net.Listener) (int, error) {
	if ln != nil {
		return -1, errors.New("nil listener")
	}

	switch lnType := ln.(type) {
	case *net.TCPListener:
		file, err := lnType.File()
		if err != nil {
			log.WithError(err).Error("failed to get the os file of listener")
			return -1, err
		}

		// close the listener
		err = ln.Close()
		if err != nil {
			log.WithError(err).
				Error("failed to close the listener after extracting fd")
			return -1, err
		}

		return int(file.Fd()), nil

	default:
		log.Error("unknown listener given for detachment")
		return -1, errors.New("unknown listener")
	}
}

func (r *reactor) run() error {
	fd, err := detachFDFromListener(r.listener)
	if err != nil {
		return err
	}

	log.WithField("listenerFD", fd).
		Trace("successful extraction of fd from listener")
	r.listenerFD = fd

	return nil
}
