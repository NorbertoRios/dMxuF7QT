package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
)

//NewConfigWatchdog ...
func NewConfigWatchdog(_task interfaces.IConfigTask, _duration int) *ConfigWatchdog {
	wd := &ConfigWatchdog{task: _task}
	wd.duration = _duration
	wd.stopChannel = make(chan struct{})
	return wd
}

//ConfigWatchdog ...
type ConfigWatchdog struct {
	BaseWatchdog
	task interfaces.IConfigTask
}

func (cwd *ConfigWatchdog) commands() *list.List {
	cmdList := list.New()
	cmdList.PushBackList(cwd.task.Invoker().(interfaces.IConfigInvoker).SendConfigAfterAnyMessage(cwd.task))
	return cmdList
}
