package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"strconv"
)

//PowerSensor power sensor
type PowerSensor struct {
	Base
	PowerState string
	Supply     int32
}

//BuildPowerSensor returns new gps sensor
func BuildPowerSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.Supply]; !f {
		return nil
	} else {
		suply := &columns.Tenth{RawValue: v}
		return &PowerSensor{Supply: suply.Value()}
	}
}

//BuildPowerSensorFromString returns new power sensor
func BuildPowerSensorFromString(data string) ISensor {
	supply, _ := strconv.ParseInt(data, 10, 32)
	return &PowerSensor{Supply: int32(supply)}
}
