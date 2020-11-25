package rabbit

import (
	"fmt"
	"genx-go/configuration"
	"genx-go/logger"

	"github.com/streadway/amqp"
)

var conn *RabbitConnection

//Connection ..
func Connection() *RabbitConnection {
	if conn == nil {
		logger.Logger().WriteToLog(logger.Fatal, "[rabbit | RabbitConnection] RabbitMQ RabbitConnection not established")
	}
	return conn
}

//InitializeRabbitRabbitConnection ...
func InitializeRabbitRabbitConnection(_config *configuration.RabbitCredentials) {
	conn = connect(_config)
}

func connect(_config *configuration.RabbitCredentials) *RabbitConnection {
	RabbitConnectionStr := fmt.Sprintf("amqp://%v:%v@%v:%v/", _config.Username, _config.Password, _config.Host, _config.Port)
	conn, err := amqp.Dial(RabbitConnectionStr)
	failOnError(err, "Failed to connect to RabbitMQ; RabbitConnection string:"+RabbitConnectionStr)
	if err != nil {
		return nil
	}
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	if err != nil {
		return nil
	}

	return &RabbitConnection{
		RabbitConnection: conn,
		channel:          ch,
	}
}

//RabbitConnection ...
type RabbitConnection struct {
	RabbitConnection *amqp.Connection
	channel          *amqp.Channel
}

//Publish ..
func (c *RabbitConnection) Publish(message, exchange, routingKey string, retry int) {
	if retry == 0 {
		logger.Logger().WriteToLog(logger.Fatal, "[RabbitConnection | Publish] Number of republishing exhausted")
	}
	err := c.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			DeliveryMode:    2,
			Headers:         make(amqp.Table, 0),
			Body:            []byte(message),
		})
	logger.Logger().WriteToLog(logger.Info, "[RabbitConnection | Publish] Publish ", message, " to ", exchange, ":", routingKey, ".")
	if err != nil {
		logger.Logger().WriteToLog(logger.Info, "[RabbitConnection | Publish] Error while publishing ", message, " to ", exchange, ":", routingKey, ". Error: ", err)
		retry = retry - 1
		c.Publish(message, exchange, routingKey, retry)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Logger().WriteToLog(logger.Fatal, msg)
	}
}
