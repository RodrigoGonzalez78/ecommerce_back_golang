package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
	Quantity    int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Ratings     []Rating           `json:"ratings,omitempty" bson:"ratings,omitempty"`
}
