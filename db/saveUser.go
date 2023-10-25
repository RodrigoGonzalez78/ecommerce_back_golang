package db

import (
	"context"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/tools"
	"github.com/RodrigoGonzalez78/models"
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
