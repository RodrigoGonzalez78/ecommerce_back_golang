package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Mobile    string             `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Flat      string             `json:"flat,omitempty" bson:"flat,omitempty"`
	Area      string             `json:"area,omitempty" bson:"area,omitempty"`
	Pincode   string             `json:"pincode,omitempty" bson:"pincode,omitempty"`
	City      string             `json:"city,omitempty" bson:"city,omitempty"`
	IsDefault bool               `json:"isDefault,omitempty" bson:"isDefault,omitempty"`
}
