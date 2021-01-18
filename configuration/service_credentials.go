package configuration

//ServiceCredentials represents service credentials
type ServiceCredentials struct {
	MysqDeviceMasterConnectionString string
	WebAPIPort                       int
	UDPPort                          int
	WorkersCount                     int
	DeviceFacadeHost                 string
	Rabbit                           *RabbitCredentials
	SystemExchange                   string
	FacadeCallbackExchange           string
	FacadeCallbackRoutingKey         string
	OrleansDebugRoutingKey           string
}
