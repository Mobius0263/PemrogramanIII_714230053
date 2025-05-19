package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Consumer struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	PhoneNumber string             `json:"phoneNumber" bson:"phoneNumber"`
}
