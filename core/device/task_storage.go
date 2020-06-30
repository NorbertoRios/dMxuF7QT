package device

import (
	"genx-go/logger"
	"sync"
)

//CounstructTaskStorage returns new task storage
func CounstructTaskStorage(device IDevice) *TaskStorage {
	storage := &TaskStorage{
		Tasks:  make(map[string]ITask, 0),
		Device: device,
		mutex:  &sync.Mutex{},
	}
	return storage
}

//TaskStorage task storage
type TaskStorage struct {
	mutex  *sync.Mutex
	Tasks  map[string]ITask
	Device IDevice
}

//NewTask add new task
func (storage *TaskStorage) NewTask(taskType string, newTask ITask) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	if _, f := storage.Tasks[taskType]; !f || taskType == ConfigTask {
		storage.addTaskToStorage(taskType, newTask)
		return
	}
	logger.Error("[TaskStorage] Cant add new task \"", taskType, "\" to storage. The same task already exists. Device : ", storage.Device.Identity())
}

func (storage *TaskStorage) createTask(taskType, callbackID string, onCompleteCallback func(string)) {
	if onCompleteCallback == nil {
		onCompleteCallback = storage.completeTask
	}
	var task ITask
	switch taskType {
	case ConfigTask:
		{
			task = BuildConfigurationTask(callbackID, storage.Device, storage.Device.OnLoadConfig(storage.Device.Identity(), CurrentConfig), onCompleteCallback)
			break
		}
	case SynchronizationTask:
		{
			task = BuildSynchronizarionTask(storage.Device, onCompleteCallback)
			break
		}
	case Diag1Wire:
	case DiagCAN:
	case DiagJBUS:
		{
			task = BuildCommandTask(taskType, callbackID, storage.Device, onCompleteCallback)
			break
		}
	default:
		logger.Error("[TaskStorage | createTask] Cant create task. Unexpected task type ", taskType)
		return
	}
	storage.addTaskToStorage(taskType, task)
}

func (storage *TaskStorage) addTaskToStorage(taskType string, task ITask) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.Tasks[taskType] = task
	logger.Info("[TaskStorage | addTaskToStorage] Task \"", taskType, "\" has been added to storage. Device : ", storage.Device.Identity())
	go storage.Tasks[taskType].Execute()
}

//NewDeviceResponce on new responce from device
func (storage *TaskStorage) NewDeviceResponce(deviceMessage interface{}) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	for _, task := range storage.Tasks {
		go task.DeviceResponce(deviceMessage)
	}
}

func (storage *TaskStorage) completeTask(taskType string) {
	storage.mutex.Lock()
	task, _ := storage.Tasks[taskType]
	storage.mutex.Unlock()
	if task.CallbackID() != "" {
		storage.Device.SendFacadeCallback(task.CallbackID())
	}
	storage.removeTask(taskType)
}

func (storage *TaskStorage) removeTask(taskType string) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	delete(storage.Tasks, taskType)
}
