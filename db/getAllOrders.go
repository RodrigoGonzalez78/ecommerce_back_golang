package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Función para recuperar todas las órdenes desde la base de datos
func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order

	// Obtener colección de órdenes
	collection := MongoCM.Database("ecommerce_back_golang").Collection("orders")

	// Consultar todas las órdenes
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	// Iterar a través de las órdenes
	for cursor.Next(context.TODO()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
