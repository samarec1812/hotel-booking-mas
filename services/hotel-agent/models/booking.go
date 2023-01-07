package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Booking struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	RoomId        primitive.ObjectID `json:"room_id"`
	UserName      string             `json:"user_name"`
	ArrivalDate   time.Time          `json:"arrival_date"`
	DepartureDate time.Time          `json:"departure_date"`
	Status        string             `json:"status"`
}
