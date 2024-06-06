package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderedProductsByUserID(userID primitive.ObjectID) ([]models.Product, error) {
	// Obtener colección de pedidos
	collection := MongoCM.Database("test").Collection("orders")

	// Consultar los pedidos del usuario actual
	cursor, err := collection.Find(context.TODO(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Iterar a través de los pedidos y recopilar los productos
	var products []models.Product
	for cursor.Next(context.TODO()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		for _, item := range order.Products {
			products = append(products, item)
		}
	}

	return products, nil
}
