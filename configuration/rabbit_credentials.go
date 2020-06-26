package configuration

//RabbitCredentials rabbit config
type RabbitCredentials struct {
	Host                     string
	Port                     int
	Username                 string
	Password                 string
	SystemExchange           string
	FacadeCallbackExchange   string
	FacadeCallbackRoutingKey string
	OrleansDebugRoutingKey   string
}
