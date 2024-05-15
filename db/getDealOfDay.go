package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetDealOfDay() (models.Product, error) {

	// Obtener colección de productos
	productsCollection := MongoCM.Database("ecommerce_back_golang").Collection("products")

	cursor, err := productsCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return models.Product{}, err
	}

	defer cursor.Close(context.TODO())

	var deal models.Product

	for cursor.Next(context.TODO()) {

		var product models.Product

		if err := cursor.Decode(&product); err != nil {
			return models.Product{}, err
		}

		if len(product.Ratings) > len(deal.Ratings) {
			deal = product
		}
	}

	return deal, nil
}
