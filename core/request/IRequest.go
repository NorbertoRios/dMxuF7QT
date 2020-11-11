package request

//IRequest ...
type IRequest interface {
	CallbackID() string
	Equal(IRequest) bool
	Serial() string
}
