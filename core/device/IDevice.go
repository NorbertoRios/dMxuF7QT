package device

import (
	"genx-go/message"
	"genx-go/repository/models"
	"time"
)

//IDevice device inteface
type IDevice interface {
	Parameter24() string
	Parameter500() string
	Send(string) error
	NewRequiredParameter(string, string)
	// OnLoadCurrentConfig() *models.ConfigurationModel
	// OnLoadNonSendedConfig() *models.ConfigurationModel
	OnLoadConfig(string, string) *models.ConfigurationModel
	SendFacadeCallback(string)
	OnSynchronizationTaskCompleted()
	MessageArrived(*message.RawMessage)
	LastActivityTimeStamp() time.Time
	Identity() string
	CreateNewTask(string, string, func(string))
	OnDeviceRemoving()
}
