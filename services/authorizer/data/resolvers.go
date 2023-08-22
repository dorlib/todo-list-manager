package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func addUserKey(username string, key string) {
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}

	result, err := UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)
}

func checkKeyExist(key string) {
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}

	result, err := UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)
}

func getUserKey(username string, key string) {
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}

	result, err := UsersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.InsertedID)
}
