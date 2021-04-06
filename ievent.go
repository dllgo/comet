package comet

type IEvent interface {
	// Parameter:out is the return value which is going to be sent back to the client.
	OnMessage(frame []byte, c IConn)
	// OnClosed fires when a connection has been closed.
	// The parameter:err is the last known connection error.
	OnClosed(c IConn, err error)
}