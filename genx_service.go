package main

import (
	"genx-go/configuration"
	"genx-go/connection/interfaces"
	"genx-go/repository"
	"genx-go/types"
	"genx-go/unitofwork"
)

func NewGenxService() {
	file := types.NewFile("/config/initializer/ReportConfiguration.xml")
	provider := configuration.ConstructCredentialsJSONProvider(file)
	serviceConfig, err := provider.ProvideCredentials()
	if err != nil {
		panic("[NewGenxService] Error while service initializing")
	}
	activityRepository := repository.NewDeviceActivityRepository(serviceConfig.MysqDeviceMasterConnectionString)
	historyRepository := repository.NewDeviceStateRepository(serviceConfig.MysqDeviceMasterConnectionString)
	uow := unitofwork.NewDeviceUnitOfWork(activityRepository, historyRepository)
}

//GenxService ...
type GenxService struct {
	server interfaces.IServer
}

//Run ...
func (service *GenxService) Run() {
	go service.server.Listen()
	for {

	}
}
