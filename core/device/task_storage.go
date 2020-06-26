package device

import (
	"log"
	"sync"
	"time"
)

//CounstructTaskStorage returns new task storage
func CounstructTaskStorage(device IDevice) *TaskStorage {
	storage := &TaskStorage{
		Tasks:  make(map[string]ITask, 0),
		Device: device,
		mutex:  &sync.Mutex{},
	}
	defer storage.synchronize()
	return storage
}

//TaskStorage task storage
type TaskStorage struct {
	mutex  *sync.Mutex
	Tasks  map[string]ITask
	Device IDevice
}

func (storage *TaskStorage) synchronize() {
	ticker := time.NewTicker(60 * time.Minute)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("[BuildBaseDevice] Recovered in recync cron:", r)
			}
		}()
		for {
			select {
			case <-ticker.C:
				BuildSynchronizarionTask(storage.Device, storage, storage.removeTask)
			}
		}
	}()
}

//NewTask add new task
func (storage *TaskStorage) NewTask(taskType string, newTask ITask) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	if _, f := storage.Tasks[taskType]; !f || taskType == ConfigTask {
		storage.addTaskToStorage(taskType, newTask)
		return
	}
	log.Println("[TaskStorage] Cant add new task ", taskType, " to storage. The same task already exists")
}

func (storage *TaskStorage) addTaskToStorage(taskType string, task ITask) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.Tasks[taskType] = task
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

func (storage *TaskStorage) completeSilentTask(taskType string) {
	storage.removeTask(taskType)
}

func (storage *TaskStorage) completeTask(taskType string) {
	storage.mutex.Lock()
	task, _ := storage.Tasks[taskType]
	storage.mutex.Unlock()
	storage.Device.SendFacadeCallback(task.CallbackID())
	storage.removeTask(taskType)
}

func (storage *TaskStorage) removeTask(taskType string) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	delete(storage.Tasks, taskType)
}
