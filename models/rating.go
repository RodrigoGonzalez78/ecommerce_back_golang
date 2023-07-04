package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rating struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Rating int                `bson:"rating"`
}
