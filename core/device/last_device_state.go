package device

import "genx-go/core/sensors"

//LastKnownDeviceState last known device state
type LastKnownDeviceState struct {
	Param24  string
	Param500 string
	Sensors  []sensors.ISensor
}
