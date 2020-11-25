package rabbitLogger

import (
	"fmt"
	"genx-go/configuration"
	"genx-go/logger"
	"genx-go/rabbit"
)

var rLogger *rabbitLogger

//Logger ...
func Logger() *rabbitLogger {
	if rLogger == nil {
		rLogger = buildRabbitLogger()
	}
	return rLogger
}

type rabbitLogger struct {
	exchange   string
	routingKey string
	retry      int
	connection *rabbit.RabbitConnection
}

func buildRabbitLogger() *rabbitLogger {
	var config *configuration.ServiceCredentials
	return &rabbitLogger{
		exchange:   config.FacadeCallbackExchange,
		routingKey: config.FacadeCallbackRoutingKey,
		connection: rabbit.Connection(),
		retry:      10,
	}
}

//WriteToLog write content to log
func (l *rabbitLogger) Publish(content ...interface{}) {
	strContent := fmt.Sprintln(content...)
	logger.Logger().WriteToLog(logger.Info, "[rabbitLogger | Publish] Publish message to BO socket: ", strContent)
	l.connection.Publish(strContent, l.exchange, l.routingKey, l.retry)
}
