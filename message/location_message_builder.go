package message

//BuildMessage build new message
func BuildMessage(rData map[string]interface{}, messageType string, identity string) *Message {
	sensorsBuilder := BuildBinaryMessageSensors()
	message := &Message{
		Sensors:     sensorsBuilder.Build(rData),
		Identity:    identity,
		MessageType: messageType,
	}
	return message
}
