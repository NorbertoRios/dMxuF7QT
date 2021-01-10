package interfaces

//IChannel ...
type IChannel interface {
	Send(interface{}) error
}
