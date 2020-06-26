package message

import (
	"genx-go/core"
	"genx-go/utils"
)

//BuildMessage build new message
func BuildMessage(rData map[string]interface{}, messageType string) *Message {
	sensorsBuilder := BuildBinaryMessageSensors()
	sUtils := &utils.StringUtils{Data: rData[core.UnitName].(string)}
	message := &Message{
		Sensors:     sensorsBuilder.Build(rData),
		Identity:    sUtils.Identity(),
		MessageType: messageType,
	}
	return message
}
