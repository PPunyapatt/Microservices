package usecases

import (
	"order/adapters/queue"
	"order/entities"
)

type OrderUsecase struct {
	orderRepo entities.OrderRepository
	rabbitMQ  queue.RabbitMQAdapter
}

func NewOrderUsecase(orderRepo entities.OrderRepository, rabbitMQ queue.RabbitMQAdapter) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: orderRepo,
		rabbitMQ:  rabbitMQ,
	}
}

func (o *OrderUsecase) CreateOrder(order entities.Order, orderItem entities.OrderItem) error {
	if err := o.orderRepo.CreateOrder(order); err != nil {
		return err
	}

	err := o.rabbitMQ.Publish("topic.stock.reserved", order)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUsecase) UpdateStatus() error {
	return nil
}
