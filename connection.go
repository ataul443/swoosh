// Copyright (c) 2020 Shekh Ataul
// This code is licensed under MIT license (see LICENSE for details)

package swoosh

import (
	"bytes"
	"net"
)

// Conn represents a thread safe client connection.
type Conn interface {
	// LocalAddr returns the local network address
	LocalAddr() net.Addr

	// RemoteAddr returns the remote network address
	RemoteAddr() net.Addr
}

type conn struct {
	fd int

	readStream  *bytes.Buffer
	writeStream *bytes.Buffer

	localAddr  net.Addr
	remoteAddr net.Addr
}

func newTCPConn(connFD int, remoteAddr, localAddr net.Addr) *conn {
	c := &conn{
		fd:         connFD,
		localAddr:  localAddr,
		remoteAddr: remoteAddr,

		// We can use a pool of buffers here
		readStream:  bytes.NewBuffer(make([]byte, 4096)),
		writeStream: bytes.NewBuffer(make([]byte, 4096)),
	}

	return c
}

func (c *conn) LocalAddr() net.Addr {
	return c.localAddr
}

func (c *conn) RemoteAddr() net.Addr {
	return c.remoteAddr
}
