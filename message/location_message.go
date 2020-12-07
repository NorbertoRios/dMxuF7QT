package message

//NewEmptyLocationMessage ..
func NewEmptyLocationMessage() *LocationMessage {
	return &LocationMessage{
		Messages: make([]*Message, 0),
		Ack:      "",
	}
}

//NewLocationMessage ...
func NewLocationMessage(messages []*Message, ack string) *LocationMessage {
	return &LocationMessage{
		Messages: messages,
		Ack:      ack,
	}
}

//LocationMessage ...
type LocationMessage struct {
	Messages []*Message
	Ack      string
}
