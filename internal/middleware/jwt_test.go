// a valid jwt token allows access through middleware
// request reaches the protected handler and returns 200 ok
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestJWTMiddleware(t *testing.T) { //defines test fun

	// Create a valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,                                //custom claim @ authorization
		"exp":     time.Now().Add(time.Hour).Unix(), //expiration time
	})
	tokenString, _ := token.SignedString(jwtKey)

	// Create a mock request with the token
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	// Create a dummy handler to test if middleware passes the request
	handler := JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Record the response
	w := httptest.NewRecorder() //creates a response recorder to capture the output
	handler.ServeHTTP(w, req)

	//result
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}
