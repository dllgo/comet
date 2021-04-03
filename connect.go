package comet

import (
	"errors"
	"log"

	"github.com/panjf2000/gnet"
)

//
type Connection struct {
	cid  string
	
	conn gnet.Conn
}

//
func NewConnection(cid string, conn gnet.Conn) IConnection {
	return &Connection{
		cid:  cid,
		conn: conn,
	}
}

//
func (c *Connection) Send(reqData []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[Connection.Send]recover send data error:", err)
		}
	}()
	if c.conn == nil {
		return errors.New("[Connection.Send] conn not exist")
	}
	return c.conn.AsyncWrite(reqData)
}
