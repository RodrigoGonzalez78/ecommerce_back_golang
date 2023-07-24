package db

import (
	"context"
	"time"

	"github.com/RodrigoGonzalez78/models"
	"github.com/RodrigoGonzalez78/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveUser(user models.User) (string, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("amazon_clone")
	col := db.Collection("users")

	user.Password, _ = tools.EncriptPassword(user.Password)

	result, err := col.InsertOne(cxt, user)

	if err != nil {
		return "", err
	}

	//Obtener el id y pasar a string
	objetId, _ := result.InsertedID.(primitive.ObjectID)
	return objetId.String(), nil

}
