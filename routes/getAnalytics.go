package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
)

// Handler de la ruta para recuperar datos de análisis
func GetAnalytics(w http.ResponseWriter, r *http.Request) {

	analytics, err := db.GetAnalytics()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analytics)
}
