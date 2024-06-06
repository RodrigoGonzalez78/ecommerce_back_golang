package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
)

func AddAddress(address models.Address) error {

	// Obtener colección de direcciones
	collection := MongoCM.Database("test").Collection("addresses")

	// Insertar la nueva dirección
	_, err := collection.InsertOne(context.TODO(), address)

	if err != nil {
		return err
	}

	return nil
}
