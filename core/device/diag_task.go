package device

import "time"

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
		TaskStorage:   storage,
		command:       Command(taskType),
		taskCompleted: onTaskCompleted,
		device:        device,
		state:         Opened,
	}
	task.TaskStorage.NewTask(task.TaskType, task)
}
