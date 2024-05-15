package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
)

// Handler de la ruta para recuperar todas las órdenes
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := db.GetAllOrders()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(orders)
}
