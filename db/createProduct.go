package db

import (
	"context"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(product models.Product) (string, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCM.Database("ecommerce_back_golang")
	col := db.Collection("products")

	result, err := col.InsertOne(cxt, product)

	if err != nil {
		return "", err
	}

	//Obtener el id y pasar a string
	objetId, _ := result.InsertedID.(primitive.ObjectID)
	return objetId.String(), nil

}
