package routes

import (
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteProduct elimina un producto por su ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	// Obtener el ID de la URL
	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	err := db.DeleteProduct(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Producto eliminado con exito!"))
}
