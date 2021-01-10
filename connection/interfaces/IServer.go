package interfaces

//IServer ...
type IServer interface {
	Listen()
	SendBytes(interface{}, []byte) (int64, error)
}
