package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Description  string             `json:"description" bson:"description"`
	Price        string             `json:"price" bson:"price"`
	Accommodates int                `json:"accommodates" bson:"accommodates"`
}
