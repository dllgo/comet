package comet

import (
	"errors"
	"net"

	"github.com/panjf2000/gnet"
)

//
type Conn struct {
	uid  string
	conn gnet.Conn
}

//
func NewConn(uid string, conn gnet.Conn) IConn {
	return &Conn{
		ui:  uid,
		conn: conn,
	}
}
 
// Context returns a user-defined context.
func (c *Conn)Context() (ctx interface{}){
	if c.conn == nil {
		return nil
	}
	return c.conn.Context()
}

// SetContext sets a user-defined context.
func (c *Conn)SetContext(ctx interface{}){
	if c.conn != nil {
		c.conn.SetContext(ctx)
	}
}

// LocalAddr is the connection's local socket address.
func (c *Conn)LocalAddr() (addr net.Addr){
	if c.conn == nil {
		return nil
	}
	return c.conn.LocalAddr()
}

// RemoteAddr is the connection's remote peer address.
func (c *Conn)RemoteAddr() (addr net.Addr){
	if c.conn == nil {
		return nil
	}
	return c.conn.RemoteAddr()
}

// Read reads all data from inbound ring-buffer and event-loop-buffer without moving "read" pointer, which means
// it does not evict the data from buffers actually and those data will present in buffers until the
// ResetBuffer method is called.
func (c *Conn)Read() (buf []byte){
	if c.conn == nil {
		return nil
	}
	return c.conn.Read()
	
}
   
// AsyncWrite writes data to client/connection asynchronously, usually you would call it in individual goroutines
// instead of the event-loop goroutines.
func (c *Conn)AsyncWrite(buf []byte) error{ 
	if c.conn == nil {
		return errors.New("[Connection.AsyncWrite] conn not exist")
	}
	c.conn.AsyncWrite(buf)
	return nil
}

// Close closes the current connection.
func (c *Conn)Close() error{
	if c.conn == nil {
		return errors.New("[Connection.Close] conn not exist")
	}
	c.conn.Close()
	return nil
}