package queue

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"order/entities"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQAdapter struct {
	channel *amqp.Channel
}

type RabbitMQAdapter interface {
	Publish(exchangeName string, eventPayload interface{}) error
	Consume(queueName string, topics []string, orderUsecase entities.OrderUsecase) error
}

func NewRabbitMQAdapter(channel *amqp.Channel) *rabbitMQAdapter {
	return &rabbitMQAdapter{
		channel: channel,
	}
}

func (r *rabbitMQAdapter) Publish(exchangeName string, eventPayload interface{}) error {
	err := r.channel.ExchangeDeclare(
		"ecommerce", // name
		"topic",     // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	eventJson, err := json.Marshal(eventPayload)
	if err != nil {
		return errors.New("error converting struct to json")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err_ := r.channel.PublishWithContext(ctx,
		"ecommerce",  // exchange
		exchangeName, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(eventJson),
		})
	if err_ != nil {
		return err_
	}

	return nil
}

func (r *rabbitMQAdapter) Consume(queueName string, topics []string, orderUsecase entities.OrderUsecase) error {
	err := r.channel.ExchangeDeclare(
		"ecommerce", // name
		"topic",     // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	q, err := r.channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	for _, routingKey := range topics {
		if err := r.channel.QueueBind(
			q.Name,      // queue name
			routingKey,  // routing key
			"ecommerce", // exchange
			false,
			nil); err != nil {
			return err
		}
	}

	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			// log.Printf(" [x] %s", d.Body)
			switch msg.RoutingKey {
			case "topic.order.update":
				orderUsecase.UpdateStatus()
			}
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

	return nil
}
