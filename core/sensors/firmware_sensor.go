package sensors

//BuildFirmwareSensor build hw sensor
func BuildFirmwareSensor(version string) ISensor {
	return &FirmwareSensor{Version: version}
}

//FirmwareSensor hardware sensor
type FirmwareSensor struct {
	Version string
}
