package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Name string             `json:"name"`
	City string             `json:"city"`
	Age  int                `json:"age"`
}
