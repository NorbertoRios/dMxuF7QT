package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/lock/request"
	"genx-go/logger"
	"time"
)

//NewElectricLockWatchdog ...
func NewElectricLockWatchdog(_task interfaces.ITask, _commands *list.List) *ElectricLockWatchdog {
	return &ElectricLockWatchdog{
		task:           _task,
		stopChannel:    make(chan struct{}),
		expirationTime: _task.Request().(*request.UnlockRequest).Time(),
	}
}

//ElectricLockWatchdog ...
type ElectricLockWatchdog struct {
	expirationTime time.Time
	task           interfaces.ITask
	stopChannel    chan struct{}
}

//Stop ...
func (wd *ElectricLockWatchdog) Stop() {
	wd.stopChannel <- struct{}{}
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
						wd.task.Cancel("Time is over")
						return
					}
				}
			}
		}
	}()
}
