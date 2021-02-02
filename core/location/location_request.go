package location

import (
	"container/list"
	"genx-go/core/device/interfaces"
	"genx-go/core/location/task"
	"genx-go/core/process"
	"genx-go/core/request"
	"genx-go/logger"
	"sync"
)

//New ...
func New() interfaces.IProcess {
	req := &Request{}
	req.ProcessTasks = list.New()
	req.Mutex = &sync.Mutex{}
	return req
}

//Request ..
type Request struct {
	process.BaseProcess
}

//NewRequest ...
func (lRequest *Request) NewRequest(req interface{}, _device interfaces.IDevice) *list.List {
	cList := list.New()
	if lRequest.ProcessCurrentTask != nil {
		cList.PushBackList(lRequest.ProcessCurrentTask.Invoker().CancelTask(lRequest.ProcessCurrentTask, "Deprecated"))
	}
	lRequest.ProcessCurrentTask = task.NewLocationTask(req.(*request.BaseRequest), _device, lRequest)
	cList.PushBackList(lRequest.ProcessCurrentTask.Commands())
	return cList
}

//TaskDone ...
func (Request) TaskDone(_task interfaces.ITask) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | Done] Task done")
}

//TaskCancel ...
func (Request) TaskCancel(_task interfaces.ITask, description string) {
	logger.Logger().WriteToLog(logger.Info, "[LocationRequest | Cancel] Task canceled. Description: ", description)
}
