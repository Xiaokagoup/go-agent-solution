package rabbitmqmanager

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func rabbitMAManager_subcribe_get() {
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

	for {
		// Get a message from the queue
		msg, ok, err := ch.Get(q.Name, true)
		if err != nil {
			log.Fatalf("Failed to get a message: %v", err)
		}

		// if a message was received, print its body
		if ok {
			log.Printf("Received a message: %v", string(msg.Body))
			processMessage(msg.Body)
		}

		// pause for one second before getting next message
		time.Sleep(time.Second)
	}

}
