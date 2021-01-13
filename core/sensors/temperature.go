package sensors

//TemperatureSensor temp sensor
type TemperatureSensor struct {
	Base
	Imei  string  `json:"Id"`
	ID    byte    `json:"-"`
	Value float32 `json:"TemperatureValue"`
}

//ToDTO ...
func (s *TemperatureSensor) ToDTO() map[string]interface{} {
	return make(map[string]interface{})
}
