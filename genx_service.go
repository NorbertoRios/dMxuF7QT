package main

import (
	"genx-go/configuration"
	"genx-go/connection"
	"genx-go/connection/controller"
	"genx-go/connection/interfaces"
	"genx-go/repository"
	"genx-go/types"
	"genx-go/unitofwork"
	"genx-go/worker"
)

func buildServiceConfiguration() *configuration.ServiceCredentials {
	file := types.NewFile("/config/initializer/credentials.example.json")
	provider := configuration.ConstructCredentialsJSONProvider(file)
	serviceConfig, err := provider.ProvideCredentials()
	if err != nil {
		panic("[NewGenxService] Error while service initializing")
	}
	return serviceConfig
}

func buildDeviceUnitOfWork(config *configuration.ServiceCredentials) unitofwork.IDeviceUnitOfWork {
	serviceConfig := buildServiceConfiguration()
	activityRepository := repository.NewDeviceActivityRepository(serviceConfig.MysqDeviceMasterConnectionString)
	historyRepository := repository.NewDeviceStateRepository(serviceConfig.MysqDeviceMasterConnectionString)
	return unitofwork.NewDeviceUnitOfWork(activityRepository, historyRepository)
}

func buildWorkersPool(config *configuration.ServiceCredentials, uow unitofwork.IDeviceUnitOfWork) *worker.WorkersPool {
	return worker.NewWorkerPool(config.WorkersCount, uow)
}

//NewGenxService ...
func NewGenxService() *GenxService {
	serviceConfiguration := buildServiceConfiguration()
	deviceUOW := buildDeviceUnitOfWork(serviceConfiguration)
	workersPool := buildWorkersPool(serviceConfiguration, deviceUOW)
	workersPool.Run()
	serverController := controller.NewRawDataController(workersPool)
	return &GenxService{
		server: connection.ConstructUDPServer("127.0.0.1", 10080, serverController),
	}
}

//GenxService ...
type GenxService struct {
	server interfaces.IServer
}

//Run ...
func (service *GenxService) Run() {
	service.server.Listen()
}
