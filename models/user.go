package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty" validate:"email"`
	Password string             `bson:"password,omitempty"`
	Address  string             `bson:"address,omitempty"`
	Type     string             `bson:"type,omitempty"`
	Cart     []Product          `bson:"cart,omitempty"`
}
