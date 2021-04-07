package main

import (
	"encoding/json"
	"example/pb"
	"fmt"
	"github.com/dllgo/comet"
	"log"
	"net"
	"strconv"
	"time"
	"unsafe"

	util2 "github.com/alberliu/gn/test/util"
	"github.com/golang/protobuf/proto"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	client := TcpClient{}
	fmt.Println("input UserId,DeviceId,SyncSequence")
	// fmt.Scanf("%d %d %d", &client.UserId, &client.DeviceId, &client.Seq)
	client.Start()
	select {}
}

func Json(i interface{}) string {
	bytes, _ := jsoniter.Marshal(i)
	return string(bytes)
}

type TcpClient struct {
	UserId   int64
	DeviceId int64
	Seq      int64
	codec    *util2.Codec
}

func (c *TcpClient) Output(s string ,p string, requestId int64, message proto.Message) {

	var input = pb.Request{
		I:       strconv.FormatInt(requestId,10),
		ID:      "100",
		Service: s,
		Path:    p,
		Version: 100,
		ST:      requestId,
	}

	if message != nil {
		bytes, err := proto.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}
		input.Data = bytes
	}

	inputByf, err := proto.Marshal(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = c.codec.Conn.Write(inputByf)
	// _, err = c.codec.Conn.Write(util2.Encode(inputByf))
	if err != nil {
		fmt.Println(err)
	}
}

func (c *TcpClient) Start() {
	connect, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		fmt.Println(err)
		return
	}

	c.codec = util2.NewCodec(connect)

	c.SignIn()
	c.SyncTrigger()
	go c.Heartbeat()
	go c.Receive()
}

func (c *TcpClient) SignIn() {
	signIn := comet.SignInInput{
		UserId:   1000,
		DeviceId: 5000000,
		Token:    "0",
	}
	c.Output("usersvc","login", time.Now().UnixNano(), &signIn)
}

func (c *TcpClient) SyncTrigger() {
	c.Output("imsvc","sync",  time.Now().UnixNano(), &comet.SyncInput{Seq: c.Seq})
}

func (c *TcpClient) Heartbeat() {
	ticker := time.NewTicker(time.Second * 20)
	for range ticker.C {
		c.Output("imsvc","heartbeat", time.Now().UnixNano(), nil)
	}
}

func (c *TcpClient) Receive() {
	for {
		_, err := c.codec.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			bytes, ok, err := c.codec.Decode()
			if err != nil {
				fmt.Println(err)
				return
			}

			if ok {
				c.HandlePackage(bytes)
				continue
			}
			break
		}
	}
}

func (c *TcpClient) HandlePackage(bytes []byte) {
	var output comet.Output
	err := proto.Unmarshal(bytes, &output)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch output.Type {
	case comet.PackageType_PT_HANDSHAKE:
		fmt.Println(Json(output))
	case comet.PackageType_PT_HEARTBEAT:
		fmt.Println("心跳响应")
	case comet.PackageType_PT_SYNC:
		fmt.Println("离线消息同步开始------")
		syncResp := comet.SyncOutput{}
		err := proto.Unmarshal(output.Data, &syncResp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("离线消息同步响应:code", output.Code, "message:", output.Message)
		for _, msg := range syncResp.Messages {
			fmt.Printf("消息：发送者类型：%d 发送者id：%d  接收者类型：%d 接收者id：%d  消息内容：%+v seq：%d \n",
				msg.Sender.SenderType, msg.Sender.SenderId, msg.ReceiverType, msg.ReceiverId, FormatMessage(msg.MessageType, msg.MessageContent), msg.Seq)
			c.Seq = msg.Seq
		}

		ack := comet.MessageACK{
			DeviceAck:   c.Seq,
			ReceiveTime: time.Now().UnixNano() / 1000000,
		}
		c.Output("imsvc","sync",output.RequestId, &ack)
		fmt.Println("离线消息同步结束------")
	case comet.PackageType_PT_MESSAGE:
		messageSend := comet.MessageSend{}
		err := proto.Unmarshal(output.Data, &messageSend)
		if err != nil {
			fmt.Println(err)
			return
		}

		msg := messageSend.Message
		fmt.Printf("消息：发送者类型：%d 发送者id：%d  接收者类型：%d 接收者id：%d  消息内容：%+v seq：%d \n",
			msg.Sender.SenderType, msg.Sender.SenderId, msg.ReceiverType, msg.ReceiverId, FormatMessage(msg.MessageType, msg.MessageContent), msg.Seq)

		c.Seq = msg.Seq
		ack := comet.MessageACK{
			DeviceAck:   msg.Seq,
			ReceiveTime: time.Now().UnixNano() / 1000000,
		}
		c.Output("imsvc","send", output.RequestId, &ack)
	default:
		fmt.Println("switch other")
	}
}

func JsonMarshal(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Println("json序列化：",err)
	}
	return Bytes2str(bytes)
}

func FormatMessage(messageType comet.MessageType, messageContent []byte) string {
	if messageType == comet.MessageType_MT_UNKNOWN {
		log.Println("error message type")
		return "error message type"
	}
	var (
		msg proto.Message
		err error
	)
	switch messageType {
	case comet.MessageType_MT_TEXT:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	case comet.MessageType_MT_FACE:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	case comet.MessageType_MT_VOICE:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	case comet.MessageType_MT_IMAGE:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	case comet.MessageType_MT_FILE:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	case comet.MessageType_MT_LOCATION:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	case comet.MessageType_MT_COMMAND:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	case comet.MessageType_MT_CUSTOM:
		msg = &comet.Text{}
		err = proto.Unmarshal(messageContent, msg)
	}

	bytes, err := jsoniter.Marshal(msg)
	if err != nil {
		log.Println("[server] start run Server",err)
		return ""
	}
	return Bytes2str(bytes)
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
