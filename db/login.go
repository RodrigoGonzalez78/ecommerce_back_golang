package db

import (
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, pass string) (models.User, bool) {
	user, found, _ := FindUser(email)

	if !found {
		return user, false
	}

	passwordBytes := []byte(pass)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
