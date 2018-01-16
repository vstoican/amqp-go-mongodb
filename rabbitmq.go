package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// RMQWorker - RabbitMQ Worker
type RMQWorker struct {
	connection *amqp.Connection
}

// RMQconnect - Create connection to RabbitMQ
func RMQconnect() {
	worker := new(RMQWorker)
	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	conn, err := amqp.Dial(connString)
	failOnError(err, "Failed to connect to RabbitMQ")

	worker.connection = conn
}

// startConsumer - Start consumer for a queue
func (worker *RMQWorker) startConsumer(queue string) {
	ch, err := worker.connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		queue,             // queue
		"amqp-go-mongodb", // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
