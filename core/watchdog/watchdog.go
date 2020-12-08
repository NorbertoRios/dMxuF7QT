package watchdog

import (
	"container/list"
	"genx-go/core/device/interfaces"
	baseInterface "genx-go/core/interfaces"
	"genx-go/core/usecase"
	"time"
)

//NewWatchdog ...
func NewWatchdog(_commands *list.List, _device interfaces.IDevice, _duration int) *Watchdog {
	return &Watchdog{
		duration:    _duration,
		useCase:     usecase.NewBaseUseCase(_device, _commands),
		stopChannel: make(chan struct{}),
	}
}

//Watchdog ...
type Watchdog struct {
	duration    int
	useCase     baseInterface.IUseCase
	stopChannel chan struct{}
}

//Stop ...
func (wd *Watchdog) Stop() {
	wd.stopChannel <- struct{}{}
}

//Start ...
func (wd *Watchdog) Start() {
	go func() {
		ticker := time.NewTicker(time.Duration(wd.duration) * time.Second)
		for {
			select {
			case <-ticker.C:
				{
					ticker.Stop()
					wd.useCase.Launch()
					return
				}
			case <-wd.stopChannel:
				{
					ticker.Stop()
					return
				}
			}
		}
	}()
}
