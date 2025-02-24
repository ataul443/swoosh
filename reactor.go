package swoosh

import (
	"errors"
	"net"

	"github.com/ataul443/swoosh/internal/poll"
	"github.com/ataul443/swoosh/internal/snet"
	log "github.com/sirupsen/logrus"
)

type reactor struct {
	eventHandler EventHandler

	listener   net.Listener
	listenerFD int

	connections map[int]Conn
	poller      *poll.Poller
}

func newReactor(ln net.Listener, eventHandler EventHandler) (*reactor, error) {
	poller, err := poll.New()
	if err != nil {
		return nil, err
	}

	el := &reactor{
		eventHandler: eventHandler,
		listener:     ln,
		connections:  make(map[int]Conn),
		poller:       poller,
	}

	return el, nil
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

	// put listener to be polled
	err = r.poller.AddRead(r.listenerFD)
	if err != nil {
		return err
	}

	// Start polling
	err = r.poller.Watch(r.processEvent)
	if err != nil {
		return err
	}

	return nil
}

func (r *reactor) processEvent(fd int, eventFlags uint32) error {
	_, connFound := r.connections[fd]
	if !connFound {
		return r.accept()
	}

	return nil
}

func (r *reactor) accept() error {
	connFD, sockAddr, err := snet.Accept(r.listenerFD)
	if err != nil {
		log.WithError(err).WithField("listenerFD", r.listenerFD).
			Error("failed to accept a new connection")
		return err
	}

	// get a remote net.Addr from sockAddr
	remoteAddr, err := snet.SockAddrToAddr(sockAddr)
	if err != nil {
		// Only error we can get is of unknown network scheme
		return err
	}
	localAddr := r.listener.Addr()

	c := newTCPConn(connFD, remoteAddr, localAddr)
	r.connections[connFD] = c

	log.WithFields(log.Fields{
		"localAddr":  c.localAddr.String(),
		"remoteAddr": c.remoteAddr.String(),
		"network":    c.remoteAddr.Network(),
	}).Trace("accepted a new connection")

	log.WithFields(log.Fields{
		"localAddr":  c.localAddr.String(),
		"remoteAddr": c.remoteAddr.String(),
		"network":    c.remoteAddr.Network(),
	}).Trace("firing OnConnOpen eventHandler on the connection")
	// Fire the OnConnOpen handler
	err = r.eventHandler.OnConnOpen(c)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"localAddr":  c.localAddr.String(),
			"remoteAddr": c.remoteAddr.String(),
			"network":    c.remoteAddr.Network(),
		}).Warn("OnConnOpen eventHandler on the connection  failed")
		return err
	}

	return nil
}
