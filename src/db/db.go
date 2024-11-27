package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://wennerberg-wedding-mongo-1:27017"))

	if err != nil {
		//Todo: look into format strings to send the error too
		log.Fatal("Connection to database failed")
	}

	return client
}
