package db

import (
	"context"
	"time"

	"github.com/RodrigoGonzalez78/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllProducts() ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCM.Database("amazon_clone")
	col := db.Collection("products")

	// Opcionalmente, puedes usar opciones adicionales de consulta aquí.
	options := options.Find()

	// Crear un slice para almacenar los productos resultantes.
	var products []models.Product

	// Consultar la base de datos para obtener todos los productos.
	cur, err := col.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	// Iterar sobre los resultados y agregarlos al slice de productos.
	for cur.Next(ctx) {
		var product models.Product
		err := cur.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	// Manejar errores en la iteración o en la consulta.
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
