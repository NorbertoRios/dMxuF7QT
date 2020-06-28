package serviceinstance

import (
	"genx-go/configuration"
	"genx-go/connection"
	"genx-go/core/device"
	"genx-go/message"
	"genx-go/repository"
)

//CounstructServiceInstance returns new service instance
func CounstructServiceInstance(credentials *configuration.ServiceCredentials) *ServiceInstance {
	serviceInstance := &ServiceInstance{}
	//serviceInstance.Storage = device.ConstructStorage(serviceInstance.onDeviceStateUpdated, serviceInstance.onDeviceCreated)
	//serviceInstance.MySQLRepository = repository.ConstructMySQLRepository(credentials)
	//serviceInstance.RabbitRepository = repository.ConstructRabbitRepository(credentials)
	return serviceInstance
}

//ServiceInstance its the oracle of genx service.
type ServiceInstance struct {
	Storage           *device.Storage
	MySQLRepository   repository.IRepository
	RabbitRepository  repository.IRepository
	RawMessageFactory *message.RawMessageFactory
}

//ReceivedNewMessage on new message received
func (si *ServiceInstance) ReceivedNewMessage(channel *connection.UDPChannel, packet []byte) {
	rawMessage := si.RawMessageFactory.BuildRawMessage(packet)
	if rawMessage == nil {
		return
	}
	//device := si.Storage.Device(rawMessage.SerialNumber)

}
