package worker

import (
	"genx-go/connection/interfaces"
	"genx-go/message"
)

//EntryData ...
type EntryData struct {
	RawMessage *message.RawMessage
	Channel    interfaces.IChannel
}
