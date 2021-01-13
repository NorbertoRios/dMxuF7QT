package device

import (
	"genx-go/core/sensors"
	"time"
)

//SensorState ...
type SensorState struct {
	deviceSensor sensors.ISensor
	updateTime   time.Time
}
