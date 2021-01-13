package sensors

//BuildFirmwareSensor build hw sensor
func BuildFirmwareSensor(_version string) ISensor {
	return &Firmware{version: _version}
}

//Firmware hardware sensor
type Firmware struct {
	Base
	version string
}

//ToDTO ....
func (s *Firmware) ToDTO() map[string]interface{} {
	return make(map[string]interface{})
}
