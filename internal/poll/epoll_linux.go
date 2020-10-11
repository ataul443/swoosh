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
	ReadEvents = unix.EPOLLIN | unix.EPOLLPRI

	// WriteEvents represents write epoll events.
	WriteEvents = unix.EPOLLOUT

	// ReadWriteEvents represents both read and write epoll events.
	ReadWriteEvents = ReadEvents | WriteEvents
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

// AddRead put the fs in epoll's interest list with read event on.
func (p *Poller) AddRead(fd int) error {
	err := unix.EpollCtl(p.epfd, unix.EPOLL_CTL_ADD,
		fd, &unix.EpollEvent{
			Events: ReadEvents,
			Fd:     int32(fd),
		})

	if err != nil {
		err = os.NewSyscallError("epoll_ctl", err)
		p.logger.WithError(err).WithField("fd", fd).
			Warn("failed to add in epoll's interest list with read event")
	}

	p.logger.WithField("fd", fd).
		Trace("added to epoll's interest list with read event")
	return err
}
