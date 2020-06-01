package sensors

import (
	"genx-go/core"
	"genx-go/core/genxcolumns"
	"strconv"
)

//PowerSensor power sensor
type PowerSensor struct {
	BaseSensor
	Supply int32
}

//BuildPowerSensor returns new gps sensor
func BuildPowerSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.Supply]; !f {
		return nil
	} else {
		suply := &genxcolumns.TenthColumn{RawValue: v}
		return &PowerSensor{Supply: suply.Value()}
	}
}

//BuildPowerSensorFromString returns new power sensor
func BuildPowerSensorFromString(data string) ISensor {
	supply, _ := strconv.ParseInt(data, 10, 32)
	return &PowerSensor{Supply: int32(supply)}
}
