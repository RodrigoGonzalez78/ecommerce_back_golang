package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
)

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	// Lógica de negocio: Obtener pedidos del usuario actual
	orders, err := db.GetOrdersByUserID(jwtmetods.IDUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Convertir a JSON y enviar la respuesta
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
