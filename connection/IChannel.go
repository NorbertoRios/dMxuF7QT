package connection

//IChannel ...
type IChannel interface {
	Send(interface{}) error
}
