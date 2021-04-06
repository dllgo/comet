package comet

import "fmt"

type comet struct {
	server IServer
}

var defaultComet = NewComet()

/*
初始化comet
*/
func NewComet() *comet {
	c := comet{
		server: NewServer(),
	}
	return &c
}

/*
启动comet
*/
func (c *comet)Serve(event IEvent,port int) {
	// 启动comet服务
	c.server.Serve(event,port)
	fmt.Printf("Starting comet server at :%d...\n", port)
}
