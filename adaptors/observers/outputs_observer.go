package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
)

//NewOutputsObserver ...
func NewOutputsObserver() *OutputsObserver {
	return &OutputsObserver{
		Symbol: "Relay",
	}
}

//OutputsObserver ...
type OutputsObserver struct {
	Symbol string
}

//Notify ...
func (o *OutputsObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildOutputsFromString(v.(string))
	}
	return sensors.BuildOutputsFromString("")
}
