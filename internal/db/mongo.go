package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo() (collection *mongo.Collection, client *mongo.Client, err error) {

	var ctx = context.TODO()

	var mongoURL = fmt.Sprintf("mongodb://%v:%v@%v:%v/?retryWrites=false",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	)
	var database = os.Getenv("MONGO_DATABASE")
	var envCollection = os.Getenv("MONGO_COLLECTION")

	// Create a new client and connect to the server
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return
	}

	collection = client.Database(database).Collection(envCollection)

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}
	log.Println("Successfully connected and pinged. - MONGODB")

	return collection, client, nil

}
