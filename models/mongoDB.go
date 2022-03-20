package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func MongoDb() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		println(err)
	}
	ctx, cencel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cencel()
	err = client.Connect(ctx)
	if err != nil {
		println(err)
	}
	database = client.Database("express")
}

func GetDatabase() *mongo.Database {
	return database
}
