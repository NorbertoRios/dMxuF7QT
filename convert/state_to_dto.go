package convert

import (
	"genx-go/core/sensors"
	"genx-go/dto"
)

//NewStateToDTO ...
func NewStateToDTO(_state []sensors.ISensor) *StateToDTO {
	return &StateToDTO{
		state: _state,
	}
}

//StateToDTO ...
type StateToDTO struct {
	state []sensors.ISensor
}

//Convert ...
func (std *StateToDTO) Convert() dto.IMessage {
	dtoMessage := dto.NewMessage()
	for _, sensor := range std.state {
		dtoMessage.AppendRange(sensor.ToDTO())
	}
	return dtoMessage
}
