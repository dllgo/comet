package comet

//
type IServer interface {
	// 启动zim服务
	Serve(port int)
	// 停止zim服务
	Stop()
}
