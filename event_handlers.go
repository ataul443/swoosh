package swoosh

// Action is a reaction taken by the server after an event.
type Action uint

const (
	CloseConnection Action = iota

	None
)

// EventHandler represents handlers which are going to be provided
// to a swoosh listener.
type EventHandler interface {

	// OnConnOpen will fire just after a new connection opens.
	OnConnOpen(conn Conn) (out []byte, action Action)

	// OnConnPacket will fire when a connection has data from client.
	OnConnPacket(conn Conn) (out []byte, action Action)

	// OnConnClose will fire just after a connection closes.
	OnConnClose(conn Conn) (action Action)
}
