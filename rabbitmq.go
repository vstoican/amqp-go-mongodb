package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/streadway/amqp"
)

// RMQWorker - RabbitMQ Worker
type RMQWorker struct {
	connection *amqp.Connection
	queue      string
	mongo      *mgo.Session
}

// NewWorker - Create connection to RabbitMQ
func NewWorker(queue string) *RMQWorker {

	var err error

	worker := new(RMQWorker)
	worker.queue = queue
	worker.mongo = mongoConnect()

	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	worker.connection, err = amqp.Dial(connString)
	failOnError(err, "Failed to connect to RabbitMQ")

	return worker
}

// startConsumer - Start consumer for a queue
func (worker *RMQWorker) startConsumer() {

	ch, err := worker.connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		worker.queue,      // queue
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
			log.Printf("Worker[%s]: Received a message: %s", worker.queue, d.Body)
			insertMessage(worker.mongo, worker.queue, d)
		}
	}()

	log.Printf(" [*] Worker[%s] Waiting for messages...", worker.queue)
	<-forever

}
