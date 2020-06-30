package device

import (
	"fmt"
	"genx-go/logger"
	"genx-go/message"
	"time"
)

//BuildSynchronizarionTask build new SynchronizarionTask
func BuildSynchronizarionTask(device IDevice, onTaskCompleted func(string)) *SynchronizarionTask {
	return &SynchronizarionTask{
		TaskType:      SynchronizationTask,
		taskCompleted: onTaskCompleted,
		device:        device,
		state:         Opened,
	}
}

//SynchronizarionTask task for synchronization
type SynchronizarionTask struct {
	TaskType                     string
	taskCompleted                func(string)
	currentParameterMessage      *message.ParametersMessage
	diagRequiredParametersSendAt time.Time
	device                       IDevice
	state                        int
}

//CallbackID for tasks which need send responce to facade
func (task *SynchronizarionTask) CallbackID() string {
	return ""
}

//Execute execute task
func (task *SynchronizarionTask) Execute() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[SynchronizarionTask] panic:Recovered in Execute:", r)
		}
	}()
	switch task.state {
	case Opened:
		{
			if err := task.device.Send("DIAG PARAMS=500,24;"); err != nil {
				logger.Error("[SynchronizarionTask] Error while command send. ", err)
				return
			}
			task.state = DiagRequiredParametersSended
			return
		}
	case DiagRequiredParametersSended:
		{
			if time.Now().UTC().Sub(task.diagRequiredParametersSendAt).Seconds() > 60 {
				task.state = Opened
				task.Execute()
				return
			}
			return
		}
	case ParametersReceived:
		{
			task.onParametersReceivedState()
		}
	}
}

func (task *SynchronizarionTask) onSubtaskCompleted(taskName string) {
	task.taskCompleted(taskName)
	task.state = Opened
	task.Execute()
}

func (task *SynchronizarionTask) onParametersReceivedState() {
	if task.state != DiagRequiredParametersSended {
		task.Execute()
		return
	}
	for _, rp := range []string{"500", "24"} {
		if _, ok := task.currentParameterMessage.Parameters[rp]; !ok {
			task.Execute()
			return
		}
	}
	if task.device.Parameter24() == task.currentParameterMessage.Parameters["24"] &&
		task.device.Parameter500() == task.currentParameterMessage.Parameters["500"] {
		task.Complete()
		return
	}
	task.device.CreateNewTask(ConfigTask, "", task.onSubtaskCompleted)
	logger.Info(fmt.Sprint("[SynchronizarionTask] New subtask for push current config to device is  created"))
	task.state = SubtaskIsActive
	task.Execute()
}

//DeviceResponce on device responce
func (task *SynchronizarionTask) DeviceResponce(responce interface{}) {
	if task.state != DiagRequiredParametersSended {
		return
	}
	switch responce.(type) {
	case *message.ParametersMessage:
		{
			if v, f := responce.(*message.ParametersMessage); f {
				task.currentParameterMessage = v
				task.state = ParametersReceived
				task.Execute()
			}
		}
	}
}

//Complete calls on task complete
func (task *SynchronizarionTask) Complete() {
	defer func() {
		task.device = nil
	}()
	task.device.OnSynchronizationTaskCompleted()
	task.taskCompleted(task.TaskType)
}
