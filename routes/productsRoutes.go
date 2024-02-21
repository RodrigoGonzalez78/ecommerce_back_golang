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

// GetProductsByCategory obtiene todos los productos de una categoría
func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {

	// Obtener la categoría de la consulta
	category := r.URL.Query().Get("category")

	// Obtener colección de productos
	collection := db.MongoCM.Database("test").Collection("products")

	// Consultar los productos por categoría
	cursor, err := collection.Find(context.TODO(), bson.M{"category": category})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer cursor.Close(context.TODO())

	// Convertir a JSON y enviar la respuesta
	var products []models.Product
	if err := cursor.All(context.TODO(), &products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(products)
}

// SearchProducts busca productos por nombre
func SearchProducts(w http.ResponseWriter, r *http.Request) {
	// Obtener el nombre de la ruta
	name := mux.Vars(r)["name"]

	// Obtener colección de productos
	collection := db.MongoCM.Database("test").Collection("products")

	// Consultar los productos por nombre
	cursor, err := collection.Find(context.TODO(), bson.M{"name": bson.M{"$regex": name, "$options": "i"}})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	defer cursor.Close(context.TODO())

	// Convertir a JSON y enviar la respuesta
	var products []models.Product

	if err := cursor.All(context.TODO(), &products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}

	json.NewEncoder(w).Encode(products)
}

// RateProduct permite calificar un producto
func RateProduct(w http.ResponseWriter, r *http.Request) {

	// Decodificar la solicitud JSON
	var rating models.Rating

	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Obtener colección de productos
	collection := db.MongoCM.Database("test").Collection("products")

	// Buscar el producto por ID
	filter := bson.M{"_id": rating.ID}

	update := bson.M{
		"$pull": bson.M{"ratings": bson.M{"userId": req.user}},
		"$push": bson.M{"ratings": bson.M{"userId": req.user, "rating": rating.Rating}},
	}

	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Product rating updated successfully"))
}

// GetDealOfDay obtiene el producto del día
func GetDealOfDay(w http.ResponseWriter, r *http.Request) {

	// Obtener colección de productos
	collection := db.MongoCM.Database("test").Collection("products")

	// Consultar todos los productos
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return

	}

	defer cursor.Close(context.TODO())

	// Obtener el producto con la mayor calificación
	var deal models.Product

	for cursor.Next(context.TODO()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if len(product.Ratings) > len(deal.Ratings) {
			deal = product
		}
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(deal)
}
