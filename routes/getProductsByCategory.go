package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
)

// Handler de la ruta para obtener productos por categoría
func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	// Obtener la categoría de la consulta
	category := r.URL.Query().Get("category")

	// Obtener productos por categoría desde la base de datos
	products, err := db.GetProductsByCategory(category)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(products)
}
