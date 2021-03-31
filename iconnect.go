package comet

type IConnection interface {
	Send(reqData []byte) error
}
