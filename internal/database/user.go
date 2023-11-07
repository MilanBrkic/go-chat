package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDatabase struct {
	collection *mongo.Collection
}
