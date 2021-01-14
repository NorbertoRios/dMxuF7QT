package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
)

//NewIBIDObserver ...
func NewIBIDObserver() *IBIDObserver {
	return &IBIDObserver{
		Symbol: "IBID",
	}
}

//IBIDObserver ...
type IBIDObserver struct {
	Symbol string
}

//Notify ...
func (o *IBIDObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildIButtonSensorFromString(v.(string))
	}
	return nil
}
