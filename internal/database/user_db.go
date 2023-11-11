package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	collection *mongo.Collection
}

func (db *UserDatabase) GetByUsername(username string) (any, bool) {
	var result bson.M

	err := db.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Println("No matching document found")
		return nil, false
	} else if err != nil {
		panic(fmt.Sprintf("Error finding document: %s", err))
	}

	return result, true
}
