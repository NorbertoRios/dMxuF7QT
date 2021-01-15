package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
	"genx-go/logger"
)

//NewTimeObserver ...
func NewTimeObserver() *TimeObserver {
	return &TimeObserver{
		Symbols: []string{"TimeStamp", "EventTime"},
	}
}

//TimeObserver ...
type TimeObserver struct {
	Symbols []string
}

//Notify ...
func (o *TimeObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	hash := make(map[string]interface{})
	for _, symbol := range o.Symbols {
		v, f := _message.GetValue(symbol)
		if !f {
			logger.Logger().WriteToLog(logger.Info, "[TimeObserver | Notify] Cant find ", symbol, " in DTO message")
			continue
		}
		hash[symbol] = v.(string)
	}
	if len(hash) == 0 {
		return nil
	}
	return sensors.BuildAdaptedTimeSensor(hash)
}
