package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
)

//NewInputsObserver ...
func NewInputsObserver() *InputsObserver {
	return &InputsObserver{
		Symbol: "GPIO",
	}
}

//InputsObserver ...
type InputsObserver struct {
	Symbol string
}

//Notify ...
func (o *InputsObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildInputsFromString(v.(string))
	}
	return nil
}
