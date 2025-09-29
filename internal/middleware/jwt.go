// verifies jwt tokens incoming req, extract the userid from token and allows only authenticated reqs
package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type contextKey string //to avoid key collsion in context

const userIDKey contextKey = "userID" //storing userid

var jwtKey = []byte("your_secret_key") // stored securely

func JWTMiddleware(next http.Handler) http.Handler { //checks dor valid jwt before allow access
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, claims, err := parseJWT(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Add userID to context
		//ctx := context.WithValue(r.Context(), "userID", claims["user_id"])
		ctx := context.WithValue(r.Context(), userIDKey, claims["user_id"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func parseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
