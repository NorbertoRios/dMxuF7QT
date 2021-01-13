package observers

import (
	"fmt"
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
	"genx-go/logger"
)

//NewPowerSensorObserver ...
func NewPowerSensorObserver() *PowerSensorObserver {
	return &PowerSensorObserver{
		Symbols: []string{"PowerState", "Supply"},
	}
}

//PowerSensorObserver ...
type PowerSensorObserver struct {
	Symbols []string
}

//Notify ...
func (o *PowerSensorObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	powerState := ""
	supply := int32(0)
	for _, symbol := range o.Symbols {
		if v, f := _message.GetValue(symbol); !f {
			logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[PowerSensorObserver | Notify] Cant find %v field in last activity. Activity: %v", symbol, _message))
			continue
		} else {
			switch symbol {
			case "PowerState":
				{
					powerState = v.(string)
					break
				}
			case "Supply":
				{
					supply = v.(int32)
					break
				}
			}
		}
	}
	return sensors.BuildPowerSensorFromString(supply, powerState)
}
