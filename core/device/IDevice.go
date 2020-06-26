package device

import "genx-go/repository/models"

//IDevice device inteface
type IDevice interface {
	Send(string) error
	OnLoadCurrentConfig() *models.ConfigurationModel
	OnLoadNonSendedConfig() *models.ConfigurationModel
	SendFacadeCallback(string)
	OnSynchronizationTaskCompleted()
}
