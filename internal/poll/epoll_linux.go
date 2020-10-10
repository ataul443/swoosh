// +build linux

package poll

import (
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

const (

	// ExceptionEvents represents problematic epoll events.
	ExceptionEvents = unix.EPOLLPRI | unix.EPOLLRDHUP | unix.EPOLLHUP |
		unix.EPOLLERR

	// ReadEvents represents read epoll events.
	ReadEvents = unix.EPOLLIN | ExceptionEvents

	// WriteEvents represents write epoll events.
	WriteEvents = unix.EPOLLOUT | ExceptionEvents
)

// Poller is responsible for watching file descriptors for events.
type Poller struct {

	// file descriptor of associated epoll instance.
	epfd int

	logger *log.Logger
}

// New returns a newly created poller.
func New(logger *log.Logger) (*Poller, error) {
	fd, err := unix.EpollCreate1(unix.EPOLL_CLOEXEC)
	if err != nil {
		err = os.NewSyscallError("epoll_create1", err)
		logger.WithError(err).Error("failed to create an epoll instance")
		return nil, err
	}

	p := &Poller{
		epfd:   fd,
		logger: logger,
	}

	return p, nil
}
