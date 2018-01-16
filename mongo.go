package main

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

func mongoConnect() *mgo.Collection {
	connString := fmt.Sprintf("mongodb://%s:%s",
		os.Getenv("MONGODB_HOST"),
		os.Getenv("MONGODB_PORT"),
	)

	session, err := mgo.Dial(connString)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(os.Getenv("MONGO_DB")).C(os.Getenv("MONGO_COLLECTION"))

	return c
}
