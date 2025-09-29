/*
package handlers

import (

	"encoding/json"
	"net/http"
	"task-manager/internal/models"
	"task-manager/internal/repository"

	"golang.org/x/crypto/bcrypt"

)

// reg fun
func Register(w http.ResponseWriter, r *http.Request) { //handles http req for user registration and w is write response back to client & r is incoming http req

		var user models.User
		json.NewDecoder(r.Body).Decode(&user) //decodes json req body into user struct

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		repository.DB.Create(&user)       //saves user record to database using GORM
		w.WriteHeader(http.StatusCreated) //sets http status to 201 created
		json.NewEncoder(w).Encode(user)
	}

// login fun

	func Login(w http.ResponseWriter, r *http.Request) {
		var input models.User
		json.NewDecoder(r.Body).Decode(&input)

		var user models.User
		repository.DB.Where("username = ?", input.Username).First(&user)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		w.Write([]byte("Login successful"))

}
*/
/*package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/internal/models"
	"task-manager/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key") // Replace with a secure key or load from env

// Register Handler
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save user to DB
	if err := repository.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login Handler
func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := repository.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}



	// Me Handler
func Me(w http.ResponseWriter, r *http.Request) {
    userID, err := ExtractUserIDFromToken(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var user models.User
    if err := repository.DB.First(&user, userID).Error; err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}
func ExtractUserIDFromToken(r *http.Request) (uint, error) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        return 0, fmt.Errorf("missing Authorization header")
    }

    tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return 0, fmt.Errorf("invalid token")
    }

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return 0, fmt.Errorf("user_id not found in token")
    }

    return uint(userIDFloat), nil
}



	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Return token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
*/
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"task-manager/internal/models"
	"task-manager/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key") // Replace with a secure key or load from env

// Register Handler
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := repository.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login Handler
func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := repository.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// Me Handler
func Me(w http.ResponseWriter, r *http.Request) {
	userID, err := ExtractUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := repository.DB.First(&user, userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// ExtractUserIDFromToken Utility
func ExtractUserIDFromToken(r *http.Request) (uint, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("missing Authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in token")
	}

	return uint(userIDFloat), nil
}
