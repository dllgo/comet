package comet

type IConn interface {
	
   // Context returns a user-defined context.
   Context() (ctx interface{})

   // SetContext sets a user-defined context.
   SetContext(ctx interface{})

   // LocalAddr is the connection's local socket address.
   LocalAddr() (addr net.Addr)

   // RemoteAddr is the connection's remote peer address.
   RemoteAddr() (addr net.Addr)

   // Read reads all data from inbound ring-buffer and event-loop-buffer without moving "read" pointer, which means
   // it does not evict the data from buffers actually and those data will present in buffers until the
   // ResetBuffer method is called.
   Read() (buf []byte)
   
   // AsyncWrite writes data to client/connection asynchronously, usually you would call it in individual goroutines
   // instead of the event-loop goroutines.
   AsyncWrite(buf []byte) error

   // Close closes the current connection.
   Close() error
}
