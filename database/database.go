package database

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connection *mongo.Database

func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(""))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(nil)
	if err != nil {
		log.Fatal(err)
	}

	Connection = client.Database("linkapi-golang")
	log.Println("mongodb connected")
}
