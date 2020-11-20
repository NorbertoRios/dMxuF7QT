package interfaces

//IClient ...
type IClient interface {
	Execute(IRequest) interface{}
}
