package routes

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/RodrigoGonzalez78/ecommerce_back_golang/db"
	"github.com/RodrigoGonzalez78/ecommerce_back_golang/models"
)

// Ruta para registrar un usuario nuevo
func SignUp(w http.ResponseWriter, r *http.Request) {

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

	// Verificar el formato del correo electrónico
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(newUser.Email) {
		http.Error(w, "Formato de correo electrónico inválido", http.StatusBadRequest)
		return
	}

	if len(newUser.Password) < 8 {
		http.Error(w, "La contaseña debe tener almenos 8 caracteres!", 400)
		return
	}

	_, encotrado, _ := db.FindUserByEmail(newUser.Email)

	if encotrado {
		http.Error(w, "Ya esta registrado el email!", 400)
		return
	}

	id, err := db.CreateUser(newUser)

	if err != nil {
		http.Error(w, "No se pudo registrar el usuario: "+err.Error(), 400)
		return
	}

	w.Write([]byte(id))
	w.WriteHeader(http.StatusCreated)
}
