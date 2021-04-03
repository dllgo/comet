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
设置消息处理器
*/
func (c *comet) SetMsgHandler(msgHandler MsgHandler) {
	c.server.SetMsgHandler(msgHandler)
}
/*
启动comet
*/
func (c *comet)Serve(port int) {
	// 启动comet服务
	c.server.Serve(port)
	fmt.Printf("Starting comet server at :%d...\n", port)
}
