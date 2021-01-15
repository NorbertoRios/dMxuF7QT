package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
	"genx-go/types"
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
		float := types.NewFloat64(v.(float64))
		return sensors.BuildIgnitionSensorFromString(float.String())
	}
	return nil
}
