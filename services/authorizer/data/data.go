package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UsersCollection *mongo.Collection

// OpenDataBase should connect to mongoDB to manage users and api-keys.
func OpenDataBase() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("connected to mongodb: %v", client)

	UsersCollection = client.Database("testing").Collection("users")

	fmt.Printf("users collection: %v", UsersCollection)
}
