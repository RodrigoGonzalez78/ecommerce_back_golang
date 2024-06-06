package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrdersByUserID(userID primitive.ObjectID) ([]models.Order, error) {
	// Obtener colección de pedidos
	collection := MongoCM.Database("test").Collection("orders")

	// Consultar los pedidos del usuario actual
	cursor, err := collection.Find(context.TODO(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Convertir los resultados en una lista de pedidos
	var orders []models.Order
	if err := cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}

	return orders, nil
}
