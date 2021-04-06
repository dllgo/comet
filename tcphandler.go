package comet

import (
	"context"
	"log"
	"sync"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

var tcponce sync.Once
var tcpinstance *TCPHandler

//Tcp 服务单例
func TCPHandlerIns(event IEvent) *TCPHandler {
	tcponce.Do(func() {
		tcpinstance = &TCPHandler{
			pool:       goroutine.Default(),
			eventHandler: event,
		}
	})
	return tcpinstance
}

//tcp event
type TCPHandler struct {
	*gnet.EventServer
	codec      gnet.ICodec
	pool       *goroutine.Pool
	gnetServer gnet.Server
	eventHandler IEvent
}

/*
回收资源
*/
func (eh *TCPHandler) Release() {
	log.Println("[TcpHandler] stop")
	eh.pool.Release()
}

/*
gnet 服务启动成功
*/
func (eh *TCPHandler) OnInitComplete(server gnet.Server) (action gnet.Action) { 
	eh.gnetServer = server
	return
}

/*
gnet 新建连接
*/
func (eh *TCPHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

/*
gnet 连接断开
*/
func (eh *TCPHandler) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		log.Println("[TcpHandler OnClosed] error:", err)
		return
	}
	ctx := c.Context().(context.Context)
	cid := ctx.Value("uid").(string)
	if eh.eventHandler!=nil { 
		eh.eventHandler.OnClosed(eh.GetConn(c),err)
	} 
	ConnectHandlerIns().D(cid)
	return
}
//接收数据
func (eh *TCPHandler) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Use ants pool to unblock the event-loop.
	err := eh.pool.Submit(func() { 
		if eh.eventHandler!=nil { 
			eh.eventHandler.OnMessage(frame,eh.GetConn(c))
		} 
	})

	if err != nil {
		log.Println("[React] error:", err)
	}
	return
} 

/**
返回连接
*/
func (eh *TCPHandler) GetConn(c gnet.Conn) IConn {
	ctx := c.Context().(context.Context)
	cid := ctx.Value("uid").(string)
	conn,_ := ConnectHandlerIns().R(cid)
	return conn.(IConn)
}

// Size 在线人数
func (eh *TCPHandler) Size() int64 {
	return ConnectHandlerIns().Size()
}
