// +build linux

package snet

import (
	"os"

	"golang.org/x/sys/unix"
)

// Accept accepts a new connection from provided lsitener file descriptor.
func Accept(listenerFD int) (connFD int, sockAddr unix.Sockaddr, err error) {
	connFD, sockAddr, err = unix.Accept(listenerFD)
	if err != nil {
		err = os.NewSyscallError("accept", err)
		return -1, nil, err
	}

	return connFD, sockAddr, err
}
