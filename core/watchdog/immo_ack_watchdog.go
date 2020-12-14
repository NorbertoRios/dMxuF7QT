package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewAckImmoWatchdog ...
func NewAckImmoWatchdog(_task interfaces.ITask, _duration int) *AckImmoWatchdog {
	wd := &AckImmoWatchdog{}
	wd.duration = _duration
	wd.stopChannel = make(chan struct{})
	wd.task = _task
	return wd
}

//AckImmoWatchdog ...
type AckImmoWatchdog struct {
	BaseWatchdog
}

func (awd *AckImmoWatchdog) commands() *list.List {
	cmdList := list.New()
	cmdList.PushBackList(awd.task.Invoker().(interfaces.IImmoInvoker).AckWatchdogsCommands(awd.task))
	return cmdList
}
