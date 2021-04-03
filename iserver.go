package comet

import "context"

//
type MsgHandler func(c *context.Context, frame []byte) (out []byte,err error)

type IServer interface {
	// 启动zim服务
	Serve(port int)
	// 停止zim服务
	Stop() 
	//设置消息处理器
	SetMsgHandler(msgHandler MsgHandler)
}
