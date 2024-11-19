package config

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var channel *amqp.Channel

func RabbitMqInit() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Connect Error", err.Error())
	} else {
		fmt.Println("Connect to RabbitMQ Success")
	}

	channel, err = conn.Channel()
	if err != nil {
		fmt.Println("Channel Error", err.Error())
	}
}

func Channel() *amqp.Channel {
	return channel
}
