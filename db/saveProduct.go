package db

import (
	"context"
	"time"

	"github.com/RodrigoGonzalez78/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveProduct(product models.Product) (string, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("amazon_clone")
	col := db.Collection("products")

	result, err := col.InsertOne(cxt, product)

	if err != nil {
		return "", err
	}

	//Obtener el id y pasar a string
	objetId, _ := result.InsertedID.(primitive.ObjectID)
	return objetId.String(), nil

}
