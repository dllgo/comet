package comet

import (
	"context"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/panjf2000/gnet"
)

var workonce sync.Once
var workinstance *WorkHandler

//工作线程单例
func WorkHandlerIns() *WorkHandler {
	workonce.Do(func() {
		workinstance = &WorkHandler{}
	})
	return workinstance
}

/*
工作线程
*/
type WorkHandler struct {
}

/**
处理接收到的消息
*/
func (wh *WorkHandler) handleFrame(frame []byte, c gnet.Conn) {
	ctx := c.Context().(context.Context)
	cid := ctx.Value("cid").(string)
	log.Println("[TcpHandler] handle 接收到", cid, "的消息")

	var input Input
	if err := proto.Unmarshal(frame, &input); err != nil {
		log.Println("[TcpHandler] handle 解码错误", err)
		return
	}

	switch input.Type {
	case PackageType_PT_HANDSHAKE:
		// 握手
		wh.handshake(c, input)
	case PackageType_PT_SYNC:
		// 同步
		wh.sync(c, input)
	case PackageType_PT_HEARTBEAT:
		//心跳
		wh.heartbeat(c, input)
	case PackageType_PT_MESSAGE:
		// 消息
		wh.message(c, input)
	case PackageType_PT_ACK:
		// 回执
		wh.ack(c, input)
	default:
		log.Println("[TcpHandler] handle 接收到", "handler switch other")
	}

}

// handshake 握手
func (wh *WorkHandler) handshake(c gnet.Conn, input Input) {
	log.Println("[TcpHandler] handle", "handshake 握手")
	var output = Output{
		Type:      PackageType_PT_HANDSHAKE,
		RequestId: input.RequestId,
	}
	frame, err := proto.Marshal(&output)
	if err != nil {
		return
	}
	c.AsyncWrite(frame)
}

// ack 回执
func (wh *WorkHandler) ack(c gnet.Conn, input Input) {
	log.Println("[TcpHandler] handle", "ack")
	var output = Output{
		Type:      PackageType_PT_ACK,
		RequestId: input.RequestId,
	}
	frame, err := proto.Marshal(&output)
	if err != nil {
		return
	}
	c.AsyncWrite(frame)
}

// sync 同步
func (wh *WorkHandler) sync(c gnet.Conn, input Input) {
	log.Println("[TcpHandler] handle", "sync")
	var output = Output{
		Type:      PackageType_PT_SYNC,
		RequestId: input.RequestId,
	}
	frame, err := proto.Marshal(&output)
	if err != nil {
		return
	}
	c.AsyncWrite(frame)
}

// Heartbeat 心跳
func (wh *WorkHandler) heartbeat(c gnet.Conn, input Input) {
	log.Println("[TcpHandler] handle", "Heartbeat 心跳")
	var output = Output{
		Type:      PackageType_PT_HEARTBEAT,
		RequestId: input.RequestId,
	}
	frame, err := proto.Marshal(&output)
	if err != nil {
		return
	}
	c.AsyncWrite(frame)
}

// Message 消息处理
func (wh *WorkHandler) message(c gnet.Conn, input Input) {
	log.Println("[TcpHandler] handle", "Heartbeat 心跳")
	var output = Output{
		Type:      PackageType_PT_HEARTBEAT,
		RequestId: input.RequestId,
	}
	frame, err := proto.Marshal(&output)
	if err != nil {
		return
	}
	c.AsyncWrite(frame)
}
