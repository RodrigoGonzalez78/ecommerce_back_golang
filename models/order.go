package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Products   []Product          `json:"products,omitempty" bson:"products,omitempty"`
	TotalPrice int                `json:"totalPrice,omitempty" bson:"totalPrice,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	UserID     string             `json:"userId,omitempty" bson:"userId,omitempty"`
	OrderedAt  int64              `json:"orderedAt,omitempty" bson:"orderedAt,omitempty"`
	Status     int                `json:"status,omitempty" bson:"status,omitempty"`
}
