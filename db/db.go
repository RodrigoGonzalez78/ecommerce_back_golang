package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCM = ConectarBD()
var clientOptions = options.Client().ApplyURI("")

//Funcion para conectar a mongodb
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa a la BD!")

	return client
}

func CheckConection() int {
	err := MongoCM.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
