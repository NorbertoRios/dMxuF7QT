package mock

import (
	"genx-go/connection"
	"genx-go/logger"
)

//UDPChannel ..
type UDPChannel struct {
	connection.UDPChannel
	lastMessage string
}

//Send message to device by UDP
func (c *UDPChannel) Send(message interface{}) error {
	logger.Logger().WriteToLog(logger.Info, "Sended message ", message.(string))
	createdDevice.(*Device).LastSentMessage = message.(string)
	return nil
}
