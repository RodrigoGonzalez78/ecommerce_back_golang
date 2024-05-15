package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
)

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
	if err := db.RateProductFromDB(rating, jwtmetods.IDUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product rating updated successfully"))
}
