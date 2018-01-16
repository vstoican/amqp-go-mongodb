package main

import (
	"os"
	"strings"
)

func main() {

	workers := make(map[string]*RMQWorker)

	loadConfiguration()

	for _, queue := range strings.Fields(os.Getenv("QUEUES")) {
		workers[queue] = NewWorker(queue)
		workers[queue].startConsumer()
	}

}

func loadConfiguration() {
	if os.Getenv("RABBITMQ_USER") == "" {
		os.Setenv("RABBITMQ_USER", "guest")
	}
	if os.Getenv("RABBITMQ_PASSWORD") == "" {
		os.Setenv("RABBITMQ_PASSWORD", "guest")
	}
	if os.Getenv("RABBITMQ_HOST") == "" {
		os.Setenv("RABBITMQ_HOST", "localhost")
	}
	if os.Getenv("RABBITMQ_PORT") == "" {
		os.Setenv("RABBITMQ_PORT", "5672")
	}
	if os.Getenv("MONGODB_HOST") == "" {
		os.Setenv("MONGODB_HOST", "localhost")
	}
	if os.Getenv("MONGODB_PORT") == "" {
		os.Setenv("MONGODB_PORT", "27017")
	}
	if os.Getenv("MONGODB_DB") == "" {
		os.Setenv("MONGODB_DB", "amqp")
	}
	if os.Getenv("MONGODB_COLLECTION") == "" {
		os.Setenv("MONGODB_COLLECTION", "messages")
	}
	if os.Getenv("QUEUES") == "" {
		os.Setenv("QUEUES", "messages")
	}
}
