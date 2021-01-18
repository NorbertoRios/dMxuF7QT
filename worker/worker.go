package worker

import (
	coreUseCase "genx-go/core/usecase"
	"genx-go/logger"
	"genx-go/message"
	"genx-go/unitofwork"
	"genx-go/worker/usecase"
	"sync"
)

//NewWorker ...
func NewWorker(_uow unitofwork.IDeviceUnitOfWork) *Worker {
	return &Worker{
		uow:            _uow,
		messageChannel: make(chan *EntryData, 1000000),
		devices:        make(map[string]bool),
		mutex:          &sync.Mutex{},
	}
}

//Worker ...
type Worker struct {
	mutex          *sync.Mutex
	uow            unitofwork.IDeviceUnitOfWork
	messageChannel chan *EntryData
	devices        map[string]bool
}

//NewDevice ...
func (w *Worker) NewDevice(identity string) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.devices[identity] = true
}

//DeviceExist ...
func (w *Worker) DeviceExist(identity string) bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.devices[identity]
}

//Push ...
func (w *Worker) Push(data *EntryData) {
	w.messageChannel <- data
}

//Run ...
func (w *Worker) Run() {
	for {
		select {
		case entryData := <-w.messageChannel:
			{
				device := w.uow.Device(entryData.RawMessage.Identity())
				if device == nil {
					w.uow.Register(entryData.RawMessage.Identity(), entryData.Channel)
					device = w.uow.Device(entryData.RawMessage.Identity())
				}
				msg := usecase.NewMessageIncomeUseCase(entryData.RawMessage, device).Launch()
				device.NewChannel(entryData.Channel)
				coreUseCase.NewMessageArrivedUseCase(device, msg).Launch()
				if w.uow.Commit() {
					if _, s := msg.(*message.LocationMessage); s {
						device.Ack()
					}
				} else {
					logger.Logger().WriteToLog(logger.Error, "[Worker | Run] Error while commit.")
				}
			}
		}
	}
}
