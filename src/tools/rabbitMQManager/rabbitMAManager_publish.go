package rabbitmqmanager

import (
	"log"

	"github.com/streadway/amqp"
)

func rabbitMAManager_publish() {
	// Define the AMQP connection URL
	url := "amqp://guest:guest@localhost:5672/"

	// Connect to the AMQP server
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to AMQP: %s", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declare an exchange
	exchangeName := "myexchange"
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // autoDelete
		false,        // internal
		false,        // noWait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	// Publish messages to the exchange
	message1 := []byte("Hello, world!")
	err = ch.Publish(
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message1,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %s", err)
	}

	message2 := []byte("Goodbye, world!")
	err = ch.Publish(
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message2,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %s", err)
	}

	log.Printf("Published messages to exchange %s", exchangeName)
}
