package main

import (
	"genx-go/configuration"
	"genx-go/connection"
	"genx-go/connection/controller"
	"genx-go/connection/interfaces"
	"genx-go/logger"
	"genx-go/repository"
	"genx-go/types"
	"genx-go/unitofwork"
	"genx-go/worker"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

func buildServiceConfiguration() *configuration.ServiceCredentials {
	file := types.NewFile("/config/initialize/credentials.example.json")
	provider := configuration.ConstructCredentialsJSONProvider(file)
	serviceConfig, err := provider.ProvideCredentials()
	if err != nil {
		panic("[NewGenxService] Error while service initializing")
	}
	return serviceConfig
}

func buildDeviceUnitOfWork(config *configuration.ServiceCredentials) unitofwork.IDeviceUnitOfWork {
	serviceConfig := buildServiceConfiguration()
	_connection, err := gorm.Open(mysql.Open(serviceConfig.MysqDeviceMasterConnectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 gLogger.Default.LogMode(gLogger.Info),
	})
	if err != nil {
		logger.Logger().WriteToLog("[GenxService | buildDeviceUnitOfWork] Error connecting to raw database:" + err.Error())
	}
	activityRepository := repository.NewDeviceActivityRepository(_connection)
	historyRepository := repository.NewDeviceStateRepository(_connection)
	return unitofwork.NewDeviceUnitOfWork(historyRepository, activityRepository)
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
		server: connection.ConstructUDPServer("", 10064, serverController),
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
