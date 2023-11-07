package database

import "go.mongodb.org/mongo-driver/mongo"

type ConversationDatabase struct {
	collection *mongo.Collection
}
