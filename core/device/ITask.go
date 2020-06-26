package device

//ITask task interface
type ITask interface {
	Complete()
	Execute()
	OnReceiveNeededMessage()
}
