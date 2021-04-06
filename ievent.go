package comet

 
type IEvent interface {
  
	// OnOpened fires when a new connection has been opened.
	// The parameter:c has information about the connection such as it's local and remote address.
	// Parameter:out is the return value which is going to be sent back to the client.
	// It is generally not recommended to send large amounts of data back to the client in OnOpened.
	//
	// Note that the bytes returned by OnOpened will be sent back to client without being encoded.
	OnOpened(c IConn) (out []byte)

	// OnClosed fires when a connection has been closed.
	// The parameter:err is the last known connection error.
	OnClosed(c IConn, err error)
 
	// React fires when a connection sends the server data.
	// Call c.Read() or c.ReadN(n) within the parameter:c to read incoming data from client.
	// Parameter:out is the return value which is going to be sent back to the client.
	OnMessage(frame []byte, c IConn) (out []byte)

	// Tick fires immediately after the server starts and will fire again
	// following the duration specified by the delay return value.
	Tick() (delay time.Duration)
}