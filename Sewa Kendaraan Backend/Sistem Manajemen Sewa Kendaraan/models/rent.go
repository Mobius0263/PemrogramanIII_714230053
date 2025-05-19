package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rent struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ConsumerID    primitive.ObjectID `json:"consumerId" bson:"consumerId"`
	ConsumerName  string             `json:"consumerName" bson:"consumerName"`
	ConsumerPhone string             `json:"consumerPhone" bson:"consumerPhone"`
	VehicleID     primitive.ObjectID `json:"vehicleId" bson:"vehicleId"`
	VehicleBrand  string             `json:"vehicleBrand" bson:"vehicleBrand"`
	VehicleModel  string             `json:"vehicleModel" bson:"vehicleModel"`
	RentDate      string             `json:"rentDate" bson:"rentDate"`
	ReturnDate    string             `json:"returnDate" bson:"returnDate"`
	TotalAmount   float64            `json:"totalAmount" bson:"totalAmount"`
}
