package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/db"
	"github.com/RodrigoGonzalez78/models"
)

// Agregar un producto
func AddProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Guardar el producto en la base de datos
	_, err = db.SaveProduct(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Obtener todos los productos
func GetProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	products, err := db.GetAllProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(products)
}
