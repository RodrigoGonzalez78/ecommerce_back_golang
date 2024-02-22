package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/gorilla/mux"
)

// Handler de la ruta para buscar productos por nombre
func SearchProducts(w http.ResponseWriter, r *http.Request) {
	// Obtener el nombre de la ruta
	name := mux.Vars(r)["name"]

	// Buscar productos por nombre en la base de datos
	products, err := db.SearchProductsByName(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(products)
}
