package interfaces

//ITask itask interface
type ITask interface {
	Start()
	Observers() []IObserver
	Device() IDevice
	Cancel(string)
	Done()
	Request() interface{}
}
