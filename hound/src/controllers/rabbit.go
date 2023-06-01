package controllers

import (
	"context"
	"fmt"
	"hound/src/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func HandleData(data []byte, originIp string) <-chan string {
	c := make(chan string)

	go func() {
		c <- handleData(data, originIp)
	}()

	return c
}

func handleData(data []byte, originIp string) string {
	conn, err := amqp.Dial(config.RabbitMQConnectionString)
	if err != nil {
		return fmt.Sprintf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Sprintf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hound",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Sprintf("%s: %s", "Failed to declare a queue", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	if err != nil {
		return fmt.Sprintf("%s: %s", "Failed to publish a message", err)
	}

	return fmt.Sprintf("[x] Sent %s", data)
}
