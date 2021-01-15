package worker

import (
	"genx-go/connection/interfaces"
	"genx-go/message"
	"genx-go/unitofwork"
)

//NewWorkerPool ...
func NewWorkerPool(workersCount int, _uow unitofwork.IDeviceUnitOfWork) *WorkersPool {
	return &WorkersPool{
		pool: newPool(workersCount, _uow),
	}
}

//WorkersPool ...
type WorkersPool struct {
	pool *pool
}

//Flush ...
func (wp *WorkersPool) Flush(rMessage *message.RawMessage, channel interfaces.IChannel) {
	data := &EntryData{RawMessage: rMessage, Channel: channel}
	for _, worker := range wp.pool.all() {
		if worker.DeviceExist(rMessage.Identity()) {
			worker.Push(data)
			return
		}
	}
	worker := wp.pool.next()
	worker.NewDevice(rMessage.Identity())
	worker.Push(data)
}

//Run ...
func (wp *WorkersPool) Run() {
	for _, worker := range wp.pool.workers {
		go worker.Run()
	}
}
