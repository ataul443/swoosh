// +build linux

package poll

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

// Poller is responsible for watching file descriptors for events.
type Poller struct {

	// file descriptor of associated epoll instance.
	epfd int
}

// New returns a newly created poller.
func New() (*Poller, error) {
	fd, err := unix.EpollCreate1(unix.EPOLL_CLOEXEC)
	if err != nil {
		err = os.NewSyscallError("epoll_create1", err)
		log.WithError(err).Error("failed to create an epoll instance")
		return nil, err
	}

	p := &Poller{
		epfd: fd,
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
		log.WithError(err).WithField("fd", fd).
			Warn("failed to add in epoll's interest list with read event")
	}

	log.WithField("fd", fd).
		Trace("added to epoll's interest list with read event")
	return err
}

// Watch monitors the epoll for events.
func (p *Poller) Watch(processEventFn func(fd int, eFlags uint32) error) error {
	if p == nil {
		return errors.New("nil poller")
	}

	events := make([]unix.EpollEvent, 100)

	for {
		eventsGot, err := unix.EpollWait(p.epfd, events, -1)
		if err != nil {
			if err == unix.EINTR {
				// No connections are ready for read or write, try again
				continue
			}
			err = os.NewSyscallError("epoll_wait", err)
			log.WithError(err).Errorf("failed to watch on epoll fd: %d", p.epfd)
			return err
		}

		events = events[:eventsGot]

		for _, e := range events {
			err = processEventFn(int(e.Fd), e.Events)
			if err != nil {
				return err
			}
		}
	}
}
