package tools

import (
	"time"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func EncriptPassword(pass string) (string, error) {

	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	return string(bytes), err
}

func GeneringJwt(t models.User) (string, error) {

	myClave := []byte("Amazon_Clone")

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

	tokenStrin, err := token.SignedString(myClave)

	if err != nil {
		return tokenStrin, err
	}

	return tokenStrin, nil
}
