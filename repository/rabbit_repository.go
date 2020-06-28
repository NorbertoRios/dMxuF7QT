package repository

import (
	"fmt"
	"genx-go/configuration"
	"log"

	"github.com/streadway/amqp"
)

//ConstructRabbitRepository returns new rabbit repo
func ConstructRabbitRepository(credentials *configuration.ServiceCredentials) *RabbitRepository {
	connectionStr := fmt.Sprintf("amqp://%v:%v@%v:%v/", credentials.Rabbit.Username, credentials.Rabbit.Password, credentials.Rabbit.Host, credentials.Rabbit.Port)
	conn, err := amqp.Dial(connectionStr)
	if err != nil {
		logger.Fatal("Failed to connect to RabbitMQ; Connection string: %s", err)
	}
	if err != nil {
		return nil
	}
	ch, err := conn.Channel()
	if err != nil {
		logger.Fatal("Failed to create rabbit channel.: %s", err)
	}
	return &RabbitRepository{
		connection: conn,
		channel:    ch,
	}
}

//RabbitRepository repository for rabbit mq
type RabbitRepository struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	//rabbitConfig      *config.RabbitConfig
	//rabbitCredentials *config.RabbitCredentials
}
