package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Status string             `json:"status" bson:"status"`
}
