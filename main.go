package main

import "os"

func main() {
	loadConfiguration()

}

func loadConfiguration() {
	os.Setenv("RABBITMQ_USER", "guest")
	os.Setenv("RABBITMQ_PASSWORD", "guest")
	os.Setenv("RABBITMQ_HOST", "localhost")
	os.Setenv("RABBITMQ_PORT", "5672")
	os.Setenv("MONGODB_HOST", "localhost")
	os.Setenv("MONGODB_PORT", "27017")
	os.Setenv("MONGODB_DB", "amqp")
	os.Setenv("MONGODB_COLLECTION", "messages")
	os.Setenv("QUEUES", "messages")

}
