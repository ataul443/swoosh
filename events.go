package swoosh

// EventHandler represents handlers which are going to be provided
// to a swoosh listener.
type EventHandler interface {

	// OnConnOpen will fire just after a new connection opens.
	OnConnOpen(conn Conn) error

	// OnConnPacket will fire when a connection has data from client.
	OnConnPacket(conn Conn) error

	// OnConnClose will fire just after a connection closes.
	OnConnClose(conn Conn) error
}
