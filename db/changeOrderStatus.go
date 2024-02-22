package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChangeOrderStatus(newOrderStatus models.Order) error {
	// Obtener colección de órdenes
	collection := MongoCM.Database("ecommerce_back_golang").Collection("orders")

	// Actualizar el estado de la orden
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": newOrderStatus.ID}, bson.M{"$set": bson.M{"status": newOrderStatus.Status}})

	return err
}
