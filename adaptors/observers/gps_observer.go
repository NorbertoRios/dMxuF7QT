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
		v, f := _message.GetValue(symbol)
		if !f {
			logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[GPSObserver | Notify] Cant find %v field in last activity. Activity: %v", symbol, _message))
			return nil
		}
		hash[symbol] = v
	}
	if len(hash) == 0 {
		logger.Logger().WriteToLog(logger.Info, fmt.Sprintf("[GPSObserver | Notify] GPS sensor not found. Symbols: %v. Message:%v", o.Symbols, _message))
		return nil
	}
	return sensors.BuildAdaptedGpsSensor(hash)
}
