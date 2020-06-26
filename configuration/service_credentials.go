package configuration

//ServiceCredentials represents service credentials
type ServiceCredentials struct {
	MysqDeviceMasterConnectionString string
	WebAPIPort                       int
	DeviceFacadeHost                 string
	Rabbit                           *RabbitCredentials
}
