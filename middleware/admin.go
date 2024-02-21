package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func AdminMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("x-auth-token")
		if token == "" {
			http.Error(w, "No auth token, access denied", http.StatusUnauthorized)
			return
		}

		verifiedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte("passwordKey"), nil
		})
		if err != nil || !verifiedToken.Valid {
			http.Error(w, "Token verification failed, authorization denied", http.StatusUnauthorized)
			return
		}

		// Obtén el ID de usuario verificado
		claims := verifiedToken.Claims.(jwt.MapClaims)
		userID := claims["id"].(string)

		// Verifica si el usuario es un administrador
		isAdmin := IsUserAdmin(userID)

		if !isAdmin {
			http.Error(w, "You are not an admin!", http.StatusUnauthorized)
			return
		}

		// Establece el ID de usuario y el token en el contexto de la solicitud
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "token", token)
		r = r.WithContext(ctx)

		// Pasa al siguiente middleware o controlador
		next.ServeHTTP(w, r)
	})
}

func IsUserAdmin(userID string) bool {
	// Aquí puedes implementar la lógica para verificar si el usuario es un administrador
	// Por ejemplo, consultando en la base de datos o comprobando algún campo en el modelo de usuario
	return true // Devuelve true si el usuario es un administrador, de lo contrario, false
}
