package sensors

import (
	"genx-go/core"
	"genx-go/core/columns"
	"time"
)

//PowerSensor power sensor
type PowerSensor struct {
	Base
	Supply int32
	State  string
}

//ToDTO ..
func (ps *PowerSensor) ToDTO() map[string]interface{} {
	hash := make(map[string]interface{})
	hash["Supply"] = ps.Supply
	hash["PowerState"] = ps.PowerState()
	return hash
}

//PowerState returns power state
func (ps *PowerSensor) PowerState() string {
	switch ps.TrigeredBy {
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
	default:
		{
			if ps.Supply > 0 {
				return "Powered"
			}
			return "Unknown"
		}
	}
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
		sensor.Trigered(data, posibleReasons)
		sensor.symbol = "Power"
		sensor.createdAt = time.Now().UTC()
		return sensor
	}
}

//BuildPowerSensorFromString returns new power sensor
func BuildPowerSensorFromString(data int32, state string) ISensor {
	sensor := &PowerSensor{Supply: data}
	sensor.symbol = "Power"
	sensor.State = state
	sensor.createdAt = time.Now().UTC()
	return sensor
}
