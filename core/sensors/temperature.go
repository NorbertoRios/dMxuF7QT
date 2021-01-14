package sensors

//BuildTemperatureSensor ...
func BuildTemperatureSensor(_id byte, _imei string, _value float32) *TemperatureSensor {
	return &TemperatureSensor{
		ID:    _id,
		Imei:  _imei,
		Value: _value,
	}
}

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
