package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
)

func GetOrderedProducts(w http.ResponseWriter, r *http.Request) {
	// Lógica de negocio: Obtener productos ordenados por el usuario actual
	products, err := db.GetOrderedProductsByUserID(jwtmetods.IDUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Convertir a JSON y enviar la respuesta
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
