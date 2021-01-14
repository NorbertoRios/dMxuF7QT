package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
)

//NewIgnitionObserver ...
func NewIgnitionObserver() *IgnitionObserver {
	return &IgnitionObserver{
		Symbol: "IgnitionState",
	}
}

//IgnitionObserver ...
type IgnitionObserver struct {
	Symbol string
}

//Notify ...
func (o *IgnitionObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildIgnitionSensorFromString(string(v.(float64)))
	}
	return nil
}
