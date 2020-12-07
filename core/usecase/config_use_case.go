package usecase

import (
	configRequest "genx-go/core/configuration/request"
	"genx-go/core/device/interfaces"
)

//NewConfigUseCase ...
func NewConfigUseCase(_device interfaces.IDevice, _req *configRequest.ConfigurationRequest) *ConfigUseCase {
	cCase := &ConfigUseCase{
		caseRequest: _req,
	}
	cCase.device = _device
	return cCase
}

//ConfigUseCase ...
type ConfigUseCase struct {
	BaseUseCase
	caseRequest *configRequest.ConfigurationRequest
}

//Launch ...
func (cCase *ConfigUseCase) Launch() {
	config := cCase.device.Configuration()
	commands := config.NewRequest(cCase.caseRequest)
	cCase.execute(commands)
}
