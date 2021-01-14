package observers

import (
	"genx-go/adaptors/dto"
	"genx-go/core/sensors"
)

//IAdaptorObserver ...
type IAdaptorObserver interface {
	Notify(*dto.DtoMessage) sensors.ISensor
}
