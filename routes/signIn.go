package routes

import (
	"encoding/json"
	"net/http"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	jwtmetods "github.com/RodrigoGonzalez78/ecommerce_back_golang/jwtMetods"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
)

// Ruta para loguear un usuario
func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidas"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	document, exist := db.CheckLogin(user.Email, user.Password)

	if !exist {
		http.Error(w, "Usuario y/o contraseña invalidas", 400)
		return
	}

	response, err := jwtmetods.GeneringJwt(document)

	if err != nil {
		http.Error(w, "Ocurrio un error "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
