package device

import (
	"genx-go/logger"
	"time"
)

//DiagTask task for send diag message
type DiagTask struct {
	TaskType          string
	taskCompleted     func(string)
	command           string
	TaskStorage       *TaskStorage
	diagRequestSendAt time.Time
	device            IDevice
	state             int
}

//Command returns command for diag tasks
func Command(taskType string) string {
	switch taskType {
	case Diag1Wire:
		{
			return "DIAG 1WIRE"
		}
	case DiagCAN:
		{
			return "DIAG CAN"
		}
	case DiagJBUS:
		{
			return "DIAG JBUSSTAT VBRIEF"
		}
	default:
		{
			return ""
		}
	}
}

//BuildDiagTask build new diag task
func BuildDiagTask(device IDevice, storage *TaskStorage, onTaskCompleted func(string), taskType string) {
	task := &DiagTask{
		TaskType:      taskType,
		command:       Command(taskType),
		taskCompleted: onTaskCompleted,
		device:        device,
		state:         Opened,
	}
	storage.NewTask(task.TaskType, task)
}

//CallbackID for tasks which need send responce to facade
func (task *DiagTask) CallbackID() string {
	return ""
}

//Execute execute task
func (task *DiagTask) Execute() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[DiagTask] panic:Recovered in Execute:", r)
		}
	}()
	switch task.state {
	case Opened:
		{
			if task.command == "" {
				logger.Error("[DiagTask | Execute] Diag task cant send empty command. Task type ", task.TaskType)
				task.Complete()
				return
			}
			if err := task.device.Send(task.command); err != nil {
				logger.Error("[DiagTask | Execute] Error while command send. ", err)
				return
			}
			task.state = Sended
			task.diagRequestSendAt = time.Now().UTC()
			return

		}
	case Sended:
		{
			if time.Now().UTC().Sub(task.diagRequestSendAt).Seconds() > 60 {
				task.state = Opened
				task.Execute()
			}
		}
	case ParametersReceived:
		{
			task.onParametersReceivedState()
		}
	}
}

func (task *DiagTask) onParametersReceivedState() {

}

func (task *DiagTask) Complete() {

}

func (task *DiagTask) DeviceResponce(responce interface{}) {

}
