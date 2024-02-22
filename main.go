package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	if db.CheckConection() == 0 {
		log.Fatal("Sin conexcion a la BD!")

	}
	router := mux.NewRouter()

	// Home
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Api de Ecommerce"))
	})

	//Auth
	router.HandleFunc("/signup", routes.SignUp).Methods("POST")
	router.HandleFunc("/signin", routes.SignIn).Methods("POST")

	//Admin
	router.HandleFunc("/admin/add-product", routes.AddProduct).Methods("POST")
	router.HandleFunc("/admin/get-products", routes.AdminGetProducts).Methods("GET")
	router.HandleFunc("/admin/delete-product/{id}", routes.DeleteProduct).Methods("POST")
	router.HandleFunc("/admin/get-orders", routes.GetOrders).Methods("GET")
	router.HandleFunc("/admin/change-order-status", routes.ChangeOrderStatus).Methods("POST")
	router.HandleFunc("/admin/analytics", routes.GetAnalytics).Methods("GET")

	//Address
	router.HandleFunc("/get-address", routes.GetAddresses).Methods("GET")
	router.HandleFunc("/add-address", routes.AddAddress).Methods("POST")
	router.HandleFunc("/delete-address", routes.DeleteAddress).Methods("POST")

	//Products
	router.HandleFunc("/products/", routes.GetProductsByCategory).Methods("GET")
	router.HandleFunc("/products/search/{name}", routes.SearchProducts).Methods("GET")
	router.HandleFunc("/rate-product", routes.RateProduct).Methods("POST")
	router.HandleFunc("/deal-of-day", routes.GetDealOfDay).Methods("GET")

	//Users
	router.HandleFunc("/add-to-cart", routes.AddToCart).Methods("POST")
	router.HandleFunc("/remove-from-cart/{id}", routes.RemoveFromCart).Methods("DELETE")
	router.HandleFunc("/save-default-address", routes.SaveDefaultAddress).Methods("POST")
	router.HandleFunc("/order", routes.OrderProduct).Methods("POST")
	router.HandleFunc("/orders/me", routes.GetMyOrders).Methods("GET")
	router.HandleFunc("/orders/products", routes.GetOrderedProducts).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	hamdlers := cors.AllowAll().Handler(router)

	fmt.Println("Servidor en ejecución en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":"+PORT, hamdlers))

}
