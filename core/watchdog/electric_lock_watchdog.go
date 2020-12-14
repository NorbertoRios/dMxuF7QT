package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/core/usecase"
	"genx-go/logger"
	"time"
)

//NewElectricLockWatchdog ...
func NewElectricLockWatchdog(_task interfaces.ITask) *ElectricLockWatchdog {
	wd := &ElectricLockWatchdog{
		expirationTime: _task.Request().(*request.UnlockRequest).Time(),
	}
	wd.task = _task
	return wd
}

//ElectricLockWatchdog ...
type ElectricLockWatchdog struct {
	BaseWatchdog
	expirationTime time.Time
}

//Start ...
func (wd *ElectricLockWatchdog) Start() {
	go func() {
		for {
			select {
			case <-wd.stopChannel:
				{
					return
				}
			default:
				{
					if wd.expirationTime.Before(time.Now().UTC()) {
						logger.Logger().WriteToLog(logger.Info, "[ElectricLockWatchdog] Electric lock change task canceled by time.")
						usecase.NewBaseUseCase(wd.task.Device(), wd.commands()).Launch()
						return
					}
				}
			}
		}
	}()
}

func (wd *ElectricLockWatchdog) commands() *list.List {
	cmdList := list.New()
	cmdList.PushBackList(wd.task.Invoker().CanselTask(wd.task, "Time is over"))
	return cmdList
}
