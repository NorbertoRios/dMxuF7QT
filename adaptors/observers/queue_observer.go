package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
)

//NewQueueObserver ...
func NewQueueObserver() *QueueObserver {
	return &QueueObserver{
		Symbol: "LocId",
	}
}

//QueueObserver ...
type QueueObserver struct {
	Symbol string
}

//Notify ...
func (o *QueueObserver) Notify(_message *dto.DtoMessage) sensors.ISensor {
	if v, f := _message.GetValue(o.Symbol); f {
		return sensors.BuildAdaptedQueueSensor(uint32(v.(float64)))
	}
	return nil
}
