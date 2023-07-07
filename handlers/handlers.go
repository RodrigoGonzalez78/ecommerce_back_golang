package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Manipulators() {
	r := mux.NewRouter()

	// Rutas
	//r.HandleFunc("/admin/add-product", AddProduct).Methods("POST")

	// Inicia el servidor HTTP
	fmt.Println("Servidor en ejecución en http://localhost:8000")
	http.ListenAndServe(":8000", r)
}
