package observers

import (
	"container/list"
	"fmt"
	"genx-go/core/device/interfaces"
	"genx-go/logger"
)

//ConsoleTestObserver ..
type ConsoleTestObserver struct {
}

//Task ..
func (observer *ConsoleTestObserver) Task() interfaces.ITask {
	return nil
}

//Attached .
func (observer *ConsoleTestObserver) Attached() {

}

//Update .
func (observer *ConsoleTestObserver) Update(msg interface{}) *list.List {
	cmd := &WriteToConsoleCommand{
		msg: fmt.Sprintf("[ConsoleTestObserver | Update ] Received new message %v", msg),
	}
	cList := list.New()
	cList.PushBack(cmd)
	return cList
}

//WriteToConsoleCommand ..
type WriteToConsoleCommand struct {
	msg interface{}
}

//Execute ..
func (command *WriteToConsoleCommand) Execute(device interfaces.IDevice) *list.List {
	logger.Logger().WriteToLog(logger.Info, command.msg)
	return list.New()
}
