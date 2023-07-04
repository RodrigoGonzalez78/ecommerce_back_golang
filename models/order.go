package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Products   []Product          `bson:"products,omitempty"`
	TotalPrice int                `bson:"totalPrice,omitempty"`
	Address    string             `bson:"address,omitempty"`
	UserID     string             `bson:"userId,omitempty"`
	OrderedAt  int64              `bson:"orderedAt,omitempty"`
	Status     int                `bson:"status,omitempty"`
}
