package database

import "go.mongodb.org/mongo-driver/mongo"

type MessageDatabase struct {
	collection *mongo.Collection
}
