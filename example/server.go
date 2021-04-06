package main

import ( 
	"log" 

	"github.com/dllgo/comet"
)
//tcp event
type EventHandler struct {
	comet.IEvent
} 
func (e *EventHandler)OnClosed(c comet.IConn, err error){
	log.Println("[EventHandler OnClosed] client: " + c.RemoteAddr().String() )
}
func (e *EventHandler)OnMessage(frame []byte, c comet.IConn){
	log.Println("[React] frame:", frame)
	c.AsyncWrite(frame)
	return 
}
  

func main() {
	srv := comet.NewComet()
	srv.Serve(&EventHandler{},6000)
}
