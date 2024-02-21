package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Estructura para procesar el jwt
type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `json:"id"`
	jwt.StandardClaims
}
