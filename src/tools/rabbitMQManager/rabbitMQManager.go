package rabbitmqmanager

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn := connectToRabbitMQServer()
	defer conn.Close()

	ch := createNewChannel(conn)
	defer ch.Close()

	q := declareQueueInChannel(ch, "ansysAgentQueue")

	publishMessageToQueue(ch, q, "Agent-3 Hello World Yangs  !")
	// consumeMessageFromQueue(ch, q)

}

func connectToRabbitMQServer() *amqp.Connection {
	url := "amqps://jiean:K9FD4zm9EJr7Fsg@b-ecba71c7-95cb-439c-9c13-9e556ef4ddec.mq.eu-west-3.amazonaws.com:5671/"

	// Connect to RabbitMQ server
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	// defer conn.Close()

	return conn
}

func createNewChannel(conn *amqp.Connection) *amqp.Channel {
	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	// defer ch.Close()

	return ch
}

func declareQueueInChannel(ch *amqp.Channel, queueName string) *amqp.Queue {
	// Declare a queue
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	return &q
}

func publishMessageToQueue(ch *amqp.Channel, q *amqp.Queue, message string) {
	// Publish a message
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
	log.Printf("Published message to queue %s: %s", q.Name, message)
}

func consumeMessageFromQueue(ch *amqp.Channel, q *amqp.Queue) {
	// Consume a message
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
