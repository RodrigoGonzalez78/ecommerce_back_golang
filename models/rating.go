package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rating struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Rating int                `json:"rating,omitempty"  bson:"rating"`
}
