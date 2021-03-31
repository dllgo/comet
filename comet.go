package comet

import "fmt"

type comet struct {
	server IServer
}

var defaultComet = newComet()

/*
初始化comet
*/
func newComet() *comet {
	c := comet{
		server: NewServer(),
	}
	return &c
}

/*
启动comet
*/
func Serve(port int) {
	// 启动comet服务
	defaultComet.server.Serve(port)
	fmt.Printf("Starting comet server at :%d...\n", port)
}
