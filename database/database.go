package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DatabaseInstance()

func DatabaseInstance() *mongo.Client {
	url := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(err)

	}
	return client
}

func OpenCollection(client *mongo.Client, collection_name string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("users-go").Collection(collection_name)
	return collection
}
