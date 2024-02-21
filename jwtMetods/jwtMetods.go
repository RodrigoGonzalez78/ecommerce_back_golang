package jwtmetods

import (
	"errors"
	"strings"
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Clave para generar el token
var SecretKey []byte = []byte("eccomerce_back")

func GeneringJwt(t models.User) (string, error) {

	payload := jwt.MapClaims{
		"email":   t.Email,
		"name":    t.Name,
		"address": t.Address,
		"type":    t.Type,
		"cart":    t.Cart,
		"_id":     t.ID.Hex(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStrin, err := token.SignedString(SecretKey)

	if err != nil {
		return tokenStrin, err
	}

	return tokenStrin, nil
}

// valores para todos los endpoints
var Email string
var IDUser primitive.ObjectID

// Proceso para extraer los datos del token
func ProcessToken(tk string) (*models.Claim, bool, primitive.ObjectID, error) {

	claim := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claim, false, IDUser, errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claim, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err == nil {

		_, found, _ := db.FindUserByEmail(claim.Email)

		if found {
			Email = claim.Email
			IDUser = claim.ID
		}

		return claim, found, IDUser, nil
	}

	if tkn.Valid {
		return claim, false, IDUser, errors.New("token invalido")
	}

	return claim, false, IDUser, err
}
