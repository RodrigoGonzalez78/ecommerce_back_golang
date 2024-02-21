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

// GetAddresses recupera todas las direcciones
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	var addresses []*models.Address

	// Obtener colección de direcciones
	collection := db.MongoCM.Database("test").Collection("addresses")

	// Consultar todas las direcciones
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer cursor.Close(context.TODO())

	// Iterar a través de las direcciones
	for cursor.Next(context.TODO()) {
		var address models.Address
		if err := cursor.Decode(&address); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		addresses = append(addresses, &address)
	}

	// Convertir a JSON y enviar la respuesta
	json.NewEncoder(w).Encode(addresses)
}

// AddAddress agrega una nueva dirección
func AddAddress(w http.ResponseWriter, r *http.Request) {
	var address models.Address

	// Decodificar la solicitud JSON
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Obtener colección de direcciones
	collection := db.MongoCM.Database("test").Collection("addresses")

	// Insertar la nueva dirección
	_, err := collection.InsertOne(context.TODO(), address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	json.NewEncoder(w).Encode(address)
}

// DeleteAddress elimina una dirección por su ID
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la URL
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// Obtener colección de direcciones
	collection := db.MongoCM.Database("test").Collection("addresses")

	// Eliminar la dirección
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Address deleted successfully"))
}
