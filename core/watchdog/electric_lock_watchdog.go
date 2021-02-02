package watchdog

import (
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/core/usecase"
	"genx-go/logger"
	"time"
)

//NewElectricLockWatchdog ...
func NewElectricLockWatchdog(_task interfaces.ITask) *ElectricLockWatchdog {
	wd := &ElectricLockWatchdog{
		task:           _task,
		expirationTime: _task.Request().(*request.UnlockRequest).Time(),
	}
	wd.device = _task.Device()
	wd.stopChannel = make(chan interface{})
	return wd
}

//ElectricLockWatchdog ...
type ElectricLockWatchdog struct {
	Watchdog
	task           interfaces.ITask
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
						cmds := wd.task.Invoker().CancelTask(wd.task, "Time is over")
						usecase.NewBaseUseCase(wd.device, cmds).Launch()
						return
					}
				}
			}
		}
	}()
}
