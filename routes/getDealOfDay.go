package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetDealOfDay obtiene el producto del día
func GetDealOfDay(w http.ResponseWriter, r *http.Request) {

	// Obtener colección de productos
	collection := db.MongoCM.Database("test").Collection("products")

	// Consultar todos los productos
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}

	defer cursor.Close(context.TODO())

	// Obtener el producto con la mayor calificación
	var deal models.Product

	for cursor.Next(context.TODO()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if len(product.Ratings) > len(deal.Ratings) {
			deal = product
		}
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(deal)
}
