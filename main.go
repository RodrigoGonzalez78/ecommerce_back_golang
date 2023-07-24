package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RodrigoGonzalez78/db"
	"github.com/RodrigoGonzalez78/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	if db.CheckConection() == 0 {
		log.Fatal("Sin conexcion a la BD!")

	}
	router := mux.NewRouter()

	// Rutas
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola mundo"))
	})

	router.HandleFunc("/signup", routes.SignUp).Methods("POST")
	router.HandleFunc("/signin", routes.SignIn).Methods("POST")

	router.HandleFunc("/add-product", routes.AddProduct).Methods("POST")
	router.HandleFunc("/get-products", routes.GetProducts).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	hamdlers := cors.AllowAll().Handler(router)

	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":"+PORT, hamdlers))

}
