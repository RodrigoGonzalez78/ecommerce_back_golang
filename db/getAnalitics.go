package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Función para recuperar datos de análisis desde la base de datos
func GetAnalytics() (map[string]float64, error) {
	analytics := make(map[string]float64)

	// Obtener colección de órdenes
	collection := MongoCM.Database("ecommerce_back_golang").Collection("orders")

	// Consultar todas las órdenes
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Variables para almacenar ganancias totales y por categoría
	var totalEarnings, mobileEarnings, essentialEarnings, applianceEarnings, booksEarnings, fashionEarnings float64

	// Iterar a través de las órdenes
	for cursor.Next(context.TODO()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}

		// Calcular ganancias totales
		for _, product := range order.Products {
			totalEarnings += float64(product.Quantity) * product.Price
		}

		// Calcular ganancias por categoría
		for _, product := range order.Products {
			switch product.Category {
			case "Mobiles":
				mobileEarnings += float64(product.Quantity) * product.Price
			case "Essentials":
				essentialEarnings += float64(product.Quantity) * product.Price
			case "Appliances":
				applianceEarnings += float64(product.Quantity) * product.Price
			case "Books":
				booksEarnings += float64(product.Quantity) * product.Price
			case "Fashion":
				fashionEarnings += float64(product.Quantity) * product.Price
			}
		}
	}

	// Almacenar los resultados en el mapa de análisis
	analytics["totalEarnings"] = totalEarnings
	analytics["mobileEarnings"] = mobileEarnings
	analytics["essentialEarnings"] = essentialEarnings
	analytics["applianceEarnings"] = applianceEarnings
	analytics["booksEarnings"] = booksEarnings
	analytics["fashionEarnings"] = fashionEarnings

	return analytics, nil
}
