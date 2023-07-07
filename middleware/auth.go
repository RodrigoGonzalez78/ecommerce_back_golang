package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {

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

		// Establece el ID de usuario y el token en el contexto de la solicitud
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "token", token)
		r = r.WithContext(ctx)

		// Pasa al siguiente middleware o controlador
		next.ServeHTTP(w, r)
	})
}
