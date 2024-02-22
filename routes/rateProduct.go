package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Función para calificar un producto en la base de datos
func RateProductFromDB(rating models.Rating) error {
	// Obtener colección de productos
	collection := db.MongoCM.Database("ecommerce_back_golang").Collection("products")

	// Filtrar el producto por ID
	filter := bson.M{"_id": rating.ID}

	// Definir la actualización para agregar la calificación del usuario
	update := bson.M{
		"$pull": bson.M{"ratings": bson.M{"userId": jwtmetods.IDUser}},
		"$push": bson.M{"ratings": bson.M{"userId": jwtmetods.IDUser, "rating": rating.Rating}},
	}

	// Actualizar el producto en la base de datos
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// Handler de la ruta para calificar un producto
func RateProduct(w http.ResponseWriter, r *http.Request) {

	// Decodificar la solicitud JSON
	var rating models.Rating

	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Calificar el producto en la base de datos
	if err := RateProductFromDB(rating); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product rating updated successfully"))
}
