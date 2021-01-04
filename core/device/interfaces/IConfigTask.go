package interfaces

//IConfigTask ...
type IConfigTask interface {
	ITask
	CurrentStringCommand() string
	CommandComplete()
	IsNextExist() bool
	GoToNextCommand()
}
