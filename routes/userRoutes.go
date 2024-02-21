package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// AddToCart añade un producto al carrito
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

	// Obtener colección de usuarios
	collection := client.Database("test").Collection("users")

	// Buscar el usuario por ID
	filter := bson.M{"_id": req.user}
	update := bson.M{
		"$addToSet": bson.M{"cart": bson.M{"product": requestData.ID, "quantity": 1}},
	}
	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product added to cart successfully"))
}

// RemoveFromCart elimina un producto del carrito
func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del producto de la ruta
	id := mux.Vars(r)["id"]

	// Obtener colección de usuarios
	collection := db.MongoCM.Database("test").Collection("users")

	// Buscar el usuario por ID
	filter := bson.M{"_id": req.user}
	update := bson.M{
		"$pull": bson.M{"cart": bson.M{"product": id}},
	}
	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product removed from cart successfully"))
}

// SaveDefaultAddress guarda la dirección predeterminada del usuario
func SaveDefaultAddress(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud JSON
	var requestData struct {
		Address string `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Obtener colección de usuarios
	collection := client.Database("test").Collection("users")

	// Buscar el usuario por ID y actualizar la dirección
	filter := bson.M{"_id": req.user}
	update := bson.M{
		"$set": bson.M{"address": requestData.Address},
	}
	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Default address saved successfully"))
}

// OrderProduct realiza un pedido de productos
func OrderProduct(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud JSON
	var requestData struct {
		Cart       []CartItem `json:"cart"`
		TotalPrice float64    `json:"totalPrice"`
		Address    string     `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Realizar las operaciones necesarias para realizar el pedido
	// ...

	// Enviar la respuesta
	w.Write([]byte("Order placed successfully"))
}

// GetMyOrders obtiene los pedidos del usuario actual
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	// Obtener colección de pedidos
	collection := db.MongoCM.Database("test").Collection("orders")

	// Consultar los pedidos del usuario actual
	cursor, err := collection.Find(context.TODO(), bson.M{"userId": req.user})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer cursor.Close(context.TODO())

	// Convertir a JSON y enviar la respuesta
	var orders []models.Order
	if err := cursor.All(context.TODO(), &orders); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(orders)
}

// GetOrderedProducts obtiene los productos ordenados por el usuario actual
func GetOrderedProducts(w http.ResponseWriter, r *http.Request) {
	// Obtener colección de pedidos
	collection := db.MongoCM.Database("test").Collection("orders")

	// Consultar los pedidos del usuario actual
	cursor, err := collection.Find(context.TODO(), bson.M{"userId": req.user})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer cursor.Close(context.TODO())

	// Iterar a través de los pedidos y recopilar los productos
	var products []models.Product
	for cursor.Next(context.TODO()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		for _, item := range order.Products {
			products = append(products, item.Product)
		}
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(products)
}
