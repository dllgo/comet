package main

import (
	"context"
	"log"

	"github.com/dllgo/comet"
)

func main() {
	srv:=comet.NewComet()
	srv.SetMsgHandler(React)
	srv.Serve(6000)
}
func React(c context.Context,frame []byte) (out []byte,err error) {
	// Use ants pool to unblock the event-loop.
	log.Println("[React] frame:", frame)
   return nil,nil
}
