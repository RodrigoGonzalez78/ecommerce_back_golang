package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Obtener todos los productos
func AdminGetProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	products, err := db.GetAllProducts()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(products)
}

// Agregar un producto
func AdminAddProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Guardar el producto en la base de datos
	_, err = db.CreateProduct(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteProduct elimina un producto por su ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la URL
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// Obtener colección de productos
	collection := db.MongoCM.Database("test").Collection("products")

	// Eliminar el producto
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product deleted successfully"))
}

// GetOrders recupera todas las órdenes
func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order

	// Obtener colección de órdenes
	collection := db.MongoCM.Database("test").Collection("orders")

	// Consultar todas las órdenes
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer cursor.Close(context.TODO())

	// Iterar a través de las órdenes
	for cursor.Next(context.TODO()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		orders = append(orders, order)
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(orders)
}

// ChangeOrderStatus cambia el estado de una orden
func ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud JSON
	var request struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Obtener colección de órdenes
	collection := db.MongoCM.Database("test").Collection("orders")

	// Actualizar el estado de la orden
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": request.ID}, bson.M{"$set": bson.M{"status": request.Status}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Order status updated successfully"))
}

// GetAnalytics recupera datos de análisis
func GetAnalytics(w http.ResponseWriter, r *http.Request) {
	var totalEarnings float64
	var mobileEarnings, essentialEarnings, applianceEarnings, booksEarnings, fashionEarnings float64

	// Obtener colección de órdenes
	collection := db.MongoCM.Database("test").Collection("orders")

	// Consultar todas las órdenes
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer cursor.Close(context.TODO())

	// Iterar a través de las órdenes
	for cursor.Next(context.TODO()) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Calcular ganancias totales
		for _, product := range order.Products {
			totalEarnings += float64(product.Quantity) * product.Price
		}

		// Calcular ganancias por categoría
		for _, product := range order.Products {
			switch product.Category {
			case "Mobiles":
				mobileEarnings += float64(product.Quantity) * product.Price
			case "Essentials":
				essentialEarnings += float64(product.Quantity) * product.Price
			case "Appliances":
				applianceEarnings += float64(product.Quantity) * product.Price
			case "Books":
				booksEarnings += float64(product.Quantity) * product.Price
			case "Fashion":
				fashionEarnings += float64(product.Quantity) * product.Price
			}
		}
	}

	// Crear la respuesta JSON
	earnings := map[string]float64{
		"totalEarnings":     totalEarnings,
		"mobileEarnings":    mobileEarnings,
		"essentialEarnings": essentialEarnings,
		"applianceEarnings": applianceEarnings,
		"booksEarnings":     booksEarnings,
		"fashionEarnings":   fashionEarnings,
	}

	// Enviar la respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(earnings)
}
