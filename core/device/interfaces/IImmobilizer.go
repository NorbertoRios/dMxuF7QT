package interfaces

//IImmobilizer ...
type IImmobilizer interface {
	IProcess
	Trigger() string
	State(IDevice) string
}
