package database

import (
	"context"
	"fmt"
	"go-chat/internal/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	User         *UserDatabase
	Conversation *ConversationDatabase
	Message      *MessageDatabase
}

func Connect() *Database {
	connectionString := config.MONGO_URL

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("test")

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the mongodb...")

	db := client.Database("go-chat")

	userCollection := db.Collection("user")
	conversationCollection := db.Collection(("conversation"))
	messageCollection := db.Collection(("message"))

	userDb := &UserDatabase{collection: userCollection}
	conversationDb := &ConversationDatabase{collection: conversationCollection}
	messageDb := &MessageDatabase{collection: messageCollection}

	return &Database{User: userDb, Conversation: conversationDb, Message: messageDb}
}
