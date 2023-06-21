package controllers

import (
	"encoding/json"
	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/dgrijalva/jwt-go"

	"net/http"
	"time"
)

var secretKey = "your-secret-key"

// LoginController handles user login and generates a JWT token
func LoginController(w http.ResponseWriter, r *http.Request) {
	var user models.User // Update the struct reference
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if !isValidUserCredentials(user.Email, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	tokenString, err := generateToken(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// isValidUserCredentials checks if the provided email and password are valid
func isValidUserCredentials(email, password string) bool {
	// Create a database connection
	newDb := models.Db

	// Perform the database query
	var user models.User
	result := newDb.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		// Handle the error
		return false
	}

	// Check if a user record was found
	if result.RowsAffected == 0 {
		return false
	}

	return true
}

// generateToken generates a JWT token for the provided email
func generateToken(email string) (string, error) {
	// Create the claims for the token
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Create a new token object with the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key to generate the final token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JwtAuthentication is a middleware for JWT authentication
func JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
