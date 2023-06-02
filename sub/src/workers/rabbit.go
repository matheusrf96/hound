package workers

import (
	"log"
	"sub/src/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type fn func()

func New(routine fn, interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	c := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				routine()
			case <-c:
				ticker.Stop()
				return
			}
		}
	}()
}

func GetRabbitMQData() {
	conn, err := amqp.Dial(config.RabbitMQConnectionString)
	if err != nil {
		log.Printf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("hound", true, false, false, false, nil)
	if err != nil {
		log.Printf("%s: %s", "Failed to declare a queue", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("%s: %s", "Failed to register a consumer", err)
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
