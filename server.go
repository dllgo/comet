package comet

import (
	"fmt"
	"log"
	"time"

	"github.com/panjf2000/gnet"
)

//
type server struct {
	msgHandler MsgHandler
	event IEvent
}

/*
初始化zim服务
*/
func NewServer() IServer {
	s := server{}
	return &s
}

/*
设置消息处理器
*/
func (s *server) SetMsgHandler(msgHandler MsgHandler) {
	s.msgHandler = msgHandler
}

/*
准备启动服务的资源
*/
func (s *server) StartTcpServe(port int) {
	log.Println("[CometServer] StartTcpServe")
	// 启动tcp
	if port < 1 {
		port = 9000
	}
	log.Fatal(gnet.Serve(
		TCPHandlerIns(s.msgHandler),
		fmt.Sprintf("tcp://:%v", port),
		gnet.WithMulticore(true),
		// gnet.WithTCPKeepAlive(time.Minute*5), // todo 需要确定是否对长连接有影响
		gnet.WithTicker(true),
	))
}

/*
准备启动服务的资源
*/
func (s *server) StartWSSServe(port int) {
	log.Println("[CometServer] StartWSSServe")
	// 启动websocket
	if port < 1 {
		port = 9001
	}
	log.Fatal(gnet.Serve(
		WSSHandlerIns(),
		fmt.Sprintf("tcp://:%v", port),
		gnet.WithMulticore(true),
		gnet.WithTCPKeepAlive(time.Minute*5), // todo 需要确定是否对长连接有影响
		gnet.WithTicker(true),
	))
}

/*
启动服务
*/
func (s *server) Serve(event IEvent,port int) {
	log.Println("[CometServer] start run Server")
	// go s.StartWSSServe(port + 1)
	// 准备启动服务的资源
	s.event = event
	s.StartTcpServe(port)

}

/*
停止服务 回收资源
*/
func (s *server) Stop() {
	log.Println("[CometServer] stop")
}
