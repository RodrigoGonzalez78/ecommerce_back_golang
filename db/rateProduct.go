package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Función para calificar un producto en la base de datos
func RateProductFromDB(rating models.Rating, idUser primitive.ObjectID) error {
	// Obtener colección de productos
	collection := MongoCM.Database("ecommerce_back_golang").Collection("products")

	// Filtrar el producto por ID
	filter := bson.M{"_id": rating.ID}

	// Definir la actualización para agregar la calificación del usuario
	update := bson.M{
		"$pull": bson.M{"ratings": bson.M{"userId": idUser}},
		"$push": bson.M{"ratings": bson.M{"userId": idUser, "rating": rating.Rating}},
	}

	// Actualizar el producto en la base de datos
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
