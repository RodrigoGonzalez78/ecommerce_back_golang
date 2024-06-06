package db

import (
	"context"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAddresses() ([]*models.Address, error) {

	var addresses []*models.Address
	collection := MongoCM.Database("ecommerce_back_golang").Collection("addresses")

	// Consultar todas las direcciones
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	// Iterar a través de las direcciones
	for cursor.Next(context.TODO()) {
		var address models.Address
		if err := cursor.Decode(&address); err != nil {
			return nil, err
		}
		addresses = append(addresses, &address)
	}

	return addresses, nil
}
