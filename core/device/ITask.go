package device

//ITask task interface
type ITask interface {
	Complete()
	Execute()
	DeviceResponce(interface{})
	CallbackID() string
}
