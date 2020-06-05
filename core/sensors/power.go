package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"strconv"
)

//PowerSensor power sensor
type PowerSensor struct {
	Base
	Supply int32
}

//PowerState returns power state
func (ps *PowerSensor) PowerState() string {
	switch ps.Trigered {
	case 1:
		{
			return "Powered"
		}
	case 2:
		{
			return "Backup battery"
		}
	case 3:
		{
			return "Power off"
		}
	case 4:
		{
			return "Sleeep Mode"
		}
	}
	if ps.Supply > 0 {
		return "Powered"
	}
	return "Unknown"
}

//BuildPowerSensor returns new gps sensor
func BuildPowerSensor(data map[string]interface{}) ISensor {
	if v, f := data[core.Supply]; !f {
		return nil
	} else {
		posibleReasons := map[byte]byte{
			0:  1, //1 - Powered
			1:  1,
			7:  1,
			48: 1,
			49: 2, //2 - Backup battery
			31: 4, //4 - Sleeep Mode
			5:  3, //3 - Power off
			59: 3,
		}
		suply := &columns.Tenth{RawValue: v}
		sensor := &PowerSensor{Supply: suply.Value()}
		sensor.Trigered = Trigered(data, posibleReasons)
		return sensor
	}
}

//BuildPowerSensorFromString returns new power sensor
func BuildPowerSensorFromString(data string) ISensor {
	supply, _ := strconv.ParseInt(data, 10, 32)
	return &PowerSensor{Supply: int32(supply)}
}
