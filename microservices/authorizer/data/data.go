package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

var DB *gorm.DB

// OpenDataBase should connect to mongoDB to manage users and api-keys.
func OpenDataBase() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("connected to mongodb: %v", client)

	usersCollection := client.Database("testing").Collection("users")

	fmt.Printf("users collection: %v", usersCollection)
}

// BeforeSave is a gorm hook in order to initiate the deadline field.
func (t *Task) BeforeSave(tx *gorm.DB) error {

}

// BeforeSave hook to hash the password before saving.
func (u *User) BeforeSave(tx *gorm.DB) error {

}
