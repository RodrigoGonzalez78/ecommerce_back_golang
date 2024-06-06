package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
)

// OrderProduct realiza un pedido de productos
func OrderProduct(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud JSON
	var requestData struct {
		Cart       []models.Product `json:"cart"`
		TotalPrice float64          `json:"totalPrice"`
		Address    string           `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Realizar las operaciones necesarias para realizar el pedido
	// ...

	// Enviar la respuesta
	w.Write([]byte("Order placed successfully"))
}
