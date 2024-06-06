package routes

import (
	"context"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteAddressDB(id primitive.ObjectID) error {

	// Obtener colección de direcciones
	collection := db.MongoCM.Database("test").Collection("addresses")

	// Eliminar la dirección
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

// DeleteAddress elimina una dirección por su ID
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la URL
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	err := DeleteAddressDB(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Enviar la respuesta
	w.Write([]byte("Address deleted successfully"))
}
