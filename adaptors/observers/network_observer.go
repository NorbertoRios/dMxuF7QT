package observers

import (
	"fmt"
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
	"genx-go/logger"
)

//NewNetworkObserver ...
func NewNetworkObserver() *NetworkObserver {
	return &NetworkObserver{
		Symbols: []string{"RSSI", "CSID"},
	}
}

//NetworkObserver ...
type NetworkObserver struct {
	Symbols []string
}

//Notify ...
func (o *NetworkObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	hash := make(map[string]interface{})
	for _, symbol := range o.Symbols {
		if v, f := _message.GetValue(symbol); !f {
			logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[NetworkObserver | Notify] Cant find %v field in last activity. Activity: %v", symbol, _message))
			continue
		} else {
			hash[symbol] = v
		}
	}
	if len(hash) == 0 {
		return nil
	}
	return sensors.BuildAdaptedNetworkSensor(hash)
}
