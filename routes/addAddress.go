package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
)

// AddAddress agrega una nueva dirección
func AddAddress(w http.ResponseWriter, r *http.Request) {

	var address models.Address

	// Decodificar la solicitud JSON
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err := db.AddAddress(address)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	json.NewEncoder(w).Encode(address)
}
