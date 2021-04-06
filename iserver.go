package comet

type IServer interface {
	// 启动zim服务
	Serve(event IEvent,port int)
	// 停止zim服务
	Stop()

}
