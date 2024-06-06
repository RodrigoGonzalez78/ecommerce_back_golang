package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveDefaultAddress(userID primitive.ObjectID, address string) error {
	// Obtener colección de usuarios
	collection := MongoCM.Database("test").Collection("users")

	// Buscar el usuario por ID y actualizar la dirección
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{"address": address},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
