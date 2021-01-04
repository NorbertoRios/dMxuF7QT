package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/usecase"
	"genx-go/logger"
	"time"
)

//NewElectricLockWatchdog ...
func NewElectricLockWatchdog(_device interfaces.IDevice, _commands *list.List, _expirationTime time.Time) *ElectricLockWatchdog {
	wd := &ElectricLockWatchdog{
		expirationTime: _expirationTime,
	}
	wd.device = _device
	wd.commands = _commands
	return wd
}

//ElectricLockWatchdog ...
type ElectricLockWatchdog struct {
	Watchdog
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
						usecase.NewBaseUseCase(wd.device, wd.commands).Launch()
						return
					}
				}
			}
		}
	}()
}
