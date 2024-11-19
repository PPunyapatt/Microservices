package main

import (
	"stock/adapters/queue"
	"stock/config"
)

func main() {
	config.DatabaseInit()
	config.RabbitMqInit()

	queueAdapter := queue.NewRabbitMQAdapter(config.Channel())
}
