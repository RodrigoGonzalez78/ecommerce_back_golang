package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
)

// GetAddresses recupera todas las direcciones
func GetAddresses(w http.ResponseWriter, r *http.Request) {

	addresses, err := db.GetAddresses()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(addresses)
}
