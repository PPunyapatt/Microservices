package main

import (
	"log"
	"order/adapters/handler"
	"order/adapters/queue"
	"order/adapters/repository"
	"order/config"
	"order/routes"
	"order/usecases"
)

func main() {
	config.DatabaseInit()
	config.RabbitMqInit()

	queueAdapter := queue.NewRabbitMQAdapter(config.Channel())

	orderRepo := repository.NewOrderRepository(config.DB())
	orderUsecase := usecases.NewOrderUsecase(orderRepo, queueAdapter)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	routes.Setup(orderHandler)
	//Consumer
	err := queueAdapter.Consume(
		"orderQueue",
		[]string{"topic.order.*"},
		orderUsecase,
	)
	failOnError(err, "Consume error")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
