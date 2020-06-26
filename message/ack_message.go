package message

//AckMessage represents ack message
type AckMessage struct {
	Identity    string
	Value       string
	MessageType string
}
