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
	OnLoadCurrentConfig() *models.ConfigurationModel
	OnLoadNonSendedConfig() *models.ConfigurationModel
	SendFacadeCallback(string)
	OnSynchronizationTaskCompleted(string, string)
	MessageArrived(*message.RawMessage)
	LastActivityTS() time.Time
	Identity() string
}
