package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
)

func SaveDefaultAddress(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud JSON
	var requestData struct {
		Address string `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Lógica de negocio: Guardar dirección predeterminada
	if err := db.SaveDefaultAddress(jwtmetods.IDUser, requestData.Address); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Default address saved successfully"))
}
