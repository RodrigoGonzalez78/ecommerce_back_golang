package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RemoveProductFromCart(userID primitive.ObjectID, productID string) error {
	// Obtener colección de usuarios
	collection := MongoCM.Database("test").Collection("users")

	// Buscar el usuario por ID y actualizar el carrito
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$pull": bson.M{"cart": bson.M{"product": productID}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
