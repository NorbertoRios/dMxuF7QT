package sensors

//BuildFirmwareSensor build hw sensor
func BuildFirmwareSensor(version string) ISensor {
	return &Firmware{Version: version}
}

//Firmware hardware sensor
type Firmware struct {
	Version string
}
