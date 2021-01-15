package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
	"genx-go/logger"
)

//NewTripObserver ...
func NewTripObserver() *TripObserver {
	return &TripObserver{
		Symbol: "Odometer",
	}
}

//TripObserver ...
type TripObserver struct {
	Symbol string
}

//Notify ...
func (o *TripObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildAdaptedTripSensor(int32(v.(float64)))
	}
	logger.Logger().WriteToLog(logger.Info, "[TripObserver | Notify] Cant find ", o.Symbol, " in dto message.")
	return nil
}
