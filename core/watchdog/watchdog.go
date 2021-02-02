package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/usecase"
	"genx-go/logger"
	"time"
)

//NewWatchdog ...
func NewWatchdog(_device interfaces.IDevice, _commands *list.List, _duration int) *Watchdog {
	return &Watchdog{
		device:      _device,
		commands:    _commands,
		duration:    _duration,
		stopChannel: make(chan interface{}),
	}
}

//Watchdog ...
type Watchdog struct {
	device      interfaces.IDevice
	commands    *list.List
	duration    int
	stopChannel chan interface{}
}

//Start ...
func (w *Watchdog) Start() {
	go func() {
		ticker := time.NewTicker(time.Duration(w.duration) * time.Second)
		for {
			select {
			case <-ticker.C:
				{
					ticker.Stop()
					usecase.NewBaseUseCase(w.device, w.commands).Launch()
					logger.Logger().WriteToLog(logger.Info, "[Watchdog | Stop] Watchdog time is elapsed.")
				}
			case <-w.stopChannel:
				{
					ticker.Stop()
					logger.Logger().WriteToLog(logger.Info, "[Watchdog | Stop] Watchdog is forced to stop")
					return
				}
			}
		}
	}()
}

//Stop ...
func (w *Watchdog) Stop() {
	w.stopChannel <- struct{}{}
}
