package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProduct(id primitive.ObjectID) error {

	// Obtener colección de productos
	collection := MongoCM.Database("ecommerce_back_golang").Collection("products")

	// Eliminar el producto
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	return err
}
