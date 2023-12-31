package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"-"`
}
