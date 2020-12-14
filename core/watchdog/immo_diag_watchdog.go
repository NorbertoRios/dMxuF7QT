package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewDiagImmoWatchdog ..
func NewDiagImmoWatchdog(_task interfaces.ITask, _duration int) *DiagImmoWatchdog {
	wd := &DiagImmoWatchdog{}
	wd.task = _task
	wd.duration = _duration
	wd.stopChannel = make(chan struct{})
	return wd
}

//DiagImmoWatchdog ...
type DiagImmoWatchdog struct {
	BaseWatchdog
}

func (dwd *DiagImmoWatchdog) commands() *list.List {
	cmdList := list.New()
	cmdList.PushBackList(dwd.task.Invoker().(interfaces.IImmoInvoker).DiagWatchdogsCommands(dwd.task))
	return cmdList
}
