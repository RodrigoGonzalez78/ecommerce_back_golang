package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
)

func GetDealOfDay(w http.ResponseWriter, r *http.Request) {
	deal, err := db.GetDealOfDay()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(deal)
}
