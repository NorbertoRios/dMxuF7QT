package sensors

//BuildFirmwareSensor build hw sensor
func BuildFirmwareSensor(_version string) ISensor {
	sensor := &Firmware{Version: _version}
	sensor.symbol = "Firmware"
	return sensor
}

//Firmware hardware sensor
type Firmware struct {
	Base
	Version string
}

//ToDTO ....
func (s *Firmware) ToDTO() map[string]interface{} {
	return make(map[string]interface{})
}
