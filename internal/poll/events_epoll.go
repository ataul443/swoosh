package poll

import "golang.org/x/sys/unix"

const (
	// ExceptionEvents represents problematic epoll events.
	ExceptionEvents = unix.EPOLLHUP | // write end closed

		unix.EPOLLRDHUP | // peer closed connection

		unix.EPOLLERR // read end closed

	// ReadEvents represents read epoll events.
	ReadEvents = ExceptionEvents |

		unix.EPOLLIN | // associated file is available for read opearation

		unix.EPOLLPRI //exceptional condition, visit https://man7.org/linux/man-pages/man2/poll.2.html

	// WriteEvents represents write epoll events.
	WriteEvents = ExceptionEvents |

		unix.EPOLLOUT // associated file is available for write operation

	// ReadWriteEvents represents both read and write epoll events.
	ReadWriteEvents = ReadEvents | WriteEvents
)
