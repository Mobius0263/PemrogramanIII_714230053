package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vehicle struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Brand       string             `json:"brand" bson:"brand"`
	Model       string             `json:"model" bson:"model"`
	Year        int                `json:"year" bson:"year"`
	IsAvailable bool               `json:"isAvailable" bson:"isAvailable"`
}
