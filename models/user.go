package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"  bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" validate:"email"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Address  string             `json:"address,omitempty" bson:"address,omitempty"`
	Type     string             `json:"type,omitempty" bson:"type,omitempty"`
	Cart     []Product          `json:"cart,omitempty" bson:"cart,omitempty"`
}
