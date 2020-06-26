package device

//IDevice device inteface
type IDevice interface {
	OnSynchronizationTaskCompleted(bool)
	Send(string) error
	Config() string
	OnConfigTaskCompleted()
}
