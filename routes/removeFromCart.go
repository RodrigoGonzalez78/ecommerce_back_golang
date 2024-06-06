package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
	"github.com/gorilla/mux"
)

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del producto de la ruta
	id := mux.Vars(r)["id"]

	// Lógica de negocio: Eliminar producto del carrito
	if err := db.RemoveProductFromCart(jwtmetods.IDUser, id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product removed from cart successfully"))
}
