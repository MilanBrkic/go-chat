package database

import (
	"context"
	"fmt"
	"go-chat/internal/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	collection *mongo.Collection
}

func (db *UserDatabase) GetByUsername(username string) (model.User, bool) {
	var result model.User

	err := db.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Println("No matching document found")
		return model.User{}, false
	} else if err != nil {
		panic(fmt.Sprintf("Error finding document: %s", err))
	}

	return result, true
}

func (db *UserDatabase) GetById(id string) (model.User, bool) {
	var result model.User

	idPrimitive, err := primitive.ObjectIDFromHex("654ff8a9d0f3ecff004bfef0")

	if err != nil {
		log.Fatal(err)
		return model.User{}, false
	}

	err = db.collection.FindOne(context.Background(), bson.M{"_id": idPrimitive}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Println("No matching document found")
		return model.User{}, false
	} else if err != nil {
		panic(fmt.Sprintf("Error finding document: %s", err))
	}

	return result, true
}

func (db *UserDatabase) Insert(user model.User) {
	_, err := db.collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(fmt.Sprintf("Error inserting document: %s", err))
	}

}
