package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RodrigoGonzalez78/db"
	"github.com/RodrigoGonzalez78/models"
	"github.com/RodrigoGonzalez78/tools"
)

// Ruta para registrar un usuario nuevo
func SignUp(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Call me")
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}

	if len(newUser.Email) == 0 {
		http.Error(w, "Imail es requerido!", 400)
		return
	}

	if len(newUser.Email) < 6 {
		http.Error(w, "La contaseña debe tener almenos 6 caracteres!", 400)
		return
	}

	_, encotrado, _ := db.FindUser(newUser.Email)

	if encotrado {
		http.Error(w, "Ya esta registrado el email!", 400)
		return
	}

	id, err := db.SaveUser(newUser)

	if err != nil {
		http.Error(w, "No se pudo registrar el usuario: "+err.Error(), 400)
		return
	}

	w.Write([]byte(id))
	w.WriteHeader(http.StatusCreated)
}

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

	document, exist := db.Login(user.Email, user.Password)

	if !exist {
		http.Error(w, "Usuario y/o contraseña invalidas", 400)
		return
	}

	response, err := tools.GeneringJwt(document)

	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
