package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal("Unable to connect to Mongo Client: ", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	fmt.Println("Successfully connected to Mongo Database")

	return client

}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("quiz").Collection(collectionName)
}
