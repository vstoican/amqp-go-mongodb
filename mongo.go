package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
	mgo "gopkg.in/mgo.v2"
)

// Message inserted in MongoDB
type Message struct {
	Date   time.Time
	Queue  string
	Fields struct {
		ConsumerTag string
		DeliveryTag uint64
		Redelivered bool
		Exchange    string
		RoutingKey  string
	}
	Properties struct {
		DeliveryMode    uint8
		ContentType     string
		ContentEncoding string
	}
	Content struct {
		Data   string
		Action string
	}
}

// mongoConnect - Create connection to MongoDB DB and collection
func mongoConnect() *mgo.Session {
	connString := fmt.Sprintf("mongodb://%s:%s",
		os.Getenv("MONGODB_HOST"),
		os.Getenv("MONGODB_PORT"),
	)

	session, err := mgo.Dial(connString)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func insertMessage(s *mgo.Session, queue string, msg amqp.Delivery) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(os.Getenv("MONGODB_DB")).C(os.Getenv("MONGODB_COLLECTION"))

	message := new(Message)
	message.Date = time.Now()
	message.Queue = queue
	message.Fields.ConsumerTag = msg.ConsumerTag
	message.Fields.DeliveryTag = msg.DeliveryTag
	message.Fields.Redelivered = msg.Redelivered
	message.Fields.Exchange = msg.Exchange
	message.Fields.RoutingKey = msg.RoutingKey
	message.Properties.ContentType = msg.ContentType
	message.Properties.ContentEncoding = msg.ContentEncoding
	message.Properties.DeliveryMode = msg.DeliveryMode
	message.Content.Data = string(msg.Body)
	message.Content.Action = msg.AppId

	err := c.Insert(*message)
	if err != nil {
		log.Fatal(err)
	}

}
