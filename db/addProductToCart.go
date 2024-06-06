package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProductToCart(userID primitive.ObjectID, productID string) error {
	// Obtener colección de usuarios
	collection := MongoCM.Database("test").Collection("users")

	// Buscar el usuario por ID y actualizar el carrito
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$addToSet": bson.M{"cart": bson.M{"product": productID, "quantity": 1}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
