package rabbitmqmanager

import (
	"log"

	"github.com/streadway/amqp"
)

func rabbitMAManager_subcribe_consume() {
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

	// Declare a queue
	queueName := "myqueue"
	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		q.Name,       // queue
		"",           // routing key
		"myexchange", // exchange
		false,        // noWait
		nil,          // args
	)
	if err != nil {
		log.Fatalf("Failed to bind queue to exchange: %s", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // autoAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err)
	}

	// Handle received messages
	for msg := range msgs {
		log.Printf("Received message from queue %s: %s", queueName, msg.Body)

		// Process the message
		processMessage(msg.Body)
	}
}

func processMessage(message []byte) {
	// TODO: Process the message
	log.Printf("Processing message: %s", message)
}
