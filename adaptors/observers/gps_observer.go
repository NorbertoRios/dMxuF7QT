package observers

import (
	"fmt"
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
	"genx-go/logger"
)

//NewGPSObserver ...
func NewGPSObserver() *GPSObserver {
	return &GPSObserver{
		Symbols: []string{"Speed", "Latitude", "Longitude", "Heading", "Satellites", "GPSStat"},
	}
}

//GPSObserver ...
type GPSObserver struct {
	Symbols []string
}

//Notify ...
func (o *GPSObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	hash := make(map[string]interface{})
	for _, symbol := range o.Symbols {
		if v, f := _message.GetValue(symbol); !f {
			logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[GPSObserver | Notify] Cant find %v field in last activity. Activity: %v", symbol, _message))
			continue
		} else {
			hash[symbol] = v
		}
	}
	return sensors.BuildGpsSensor(hash)
}
