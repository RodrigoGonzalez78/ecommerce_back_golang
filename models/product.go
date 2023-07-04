package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Images      []string           `bson:"images,omitempty"`
	Quantity    int                `bson:"quantity,omitempty"`
	Price       int                `bson:"price,omitempty"`
	Category    string             `bson:"category,omitempty"`
	Ratings     []Rating           `bson:"ratings,omitempty"`
}
