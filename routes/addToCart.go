package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud JSON
	var requestData struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Lógica de negocio: Añadir producto al carrito
	if err := db.AddProductToCart(jwtmetods.IDUser, requestData.ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product added to cart successfully"))
}
