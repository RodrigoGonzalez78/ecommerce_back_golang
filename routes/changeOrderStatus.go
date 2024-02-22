package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
)

// ChangeOrderStatus cambia el estado de una orden
func ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {

	// Decodificar la solicitud JSON
	var changeOrder models.Order

	if err := json.NewDecoder(r.Body).Decode(&changeOrder); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err := db.ChangeOrderStatus(changeOrder)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Order status updated successfully"))
}
