package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Función para buscar productos por nombre desde la base de datos
func SearchProductsByName(name string) ([]models.Product, error) {

	// Obtener colección de productos
	collection := MongoCM.Database("ecommerce_back_golang").Collection("products")

	// Consultar los productos por nombre
	cursor, err := collection.Find(context.TODO(), bson.M{"name": bson.M{"$regex": name, "$options": "i"}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	// Convertir a slice de productos
	var products []models.Product

	if err := cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}

	return products, nil
}
