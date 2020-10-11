package snet

import (
	"errors"
	"net"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

// SockAddrToAddr translates a unix.Sockaddr to appropriate net.Addr.
// Returns an error when got an unknown network.
func SockAddrToAddr(sockAddr unix.Sockaddr) (net.Addr, error) {
	var addr net.Addr

	switch sa := sockAddr.(type) {
	case *unix.SockaddrInet4:
		addr = &net.TCPAddr{
			IP:   append([]byte{}, sa.Addr[:]...),
			Port: sa.Port,
		}
		log.WithFields(log.Fields{
			"addr":    addr.String(),
			"network": "tcp-inet4",
		}).Trace("translated unix.Sockaddr to net.Addr")

	case *unix.SockaddrInet6:
		var zone string
		if sa.ZoneId != 0 {
			interfaceIdx, err := net.InterfaceByIndex(int(sa.ZoneId))
			if err == nil {
				zone = interfaceIdx.Name
			}
		}

		if !(zone == "" && sa.ZoneId != 0) {
			addr = &net.TCPAddr{
				IP:   append([]byte{}, sa.Addr[:]...),
				Port: sa.Port,
				Zone: zone,
			}
		}

		log.WithFields(log.Fields{
			"addr":    addr.String(),
			"network": "tcp-inet6",
			"zone":    "zone",
		}).Trace("translated unix.Sockaddr to net.Addr")

	default:
		log.Error("unknown network type for translation")
		return nil, errors.New("unknown network type")
	}

	return addr, nil
}
