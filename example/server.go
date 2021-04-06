package main

import (
	"context"
	"log"

	"github.com/dllgo/comet"
)
//tcp event
type EventHandler struct {
	*comet.IEvent
} 
func (e *EventHandler)OnOpened(c comet.IConn) (out []byte){
	log.Println(fmt.Sprintf("[EventHandler OnOpened] client: RemoteAddr:%v", c.RemoteAddr().String()))
	
} 
func (e *EventHandler)OnClosed(c comet.IConn, err error){
	log.Println("[EventHandler OnClosed] client: " + c.LocalAddr().String() )
	
}
func (e *EventHandler)OnMessage(frame []byte, c comet.IConn) (out []byte){
	log.Println("[React] frame:", frame)
	return nil
}
 
func (e *EventHandler)Tick() (delay time.Duration){
	var interval time.Duration
	interval = 20 * time.Second
	delay = interval 
	return
} 

func main() {
	srv := comet.NewComet()
	srv.SetMsgHandler(React)
	srv.Serve(&EventHandler{},6000)
}
