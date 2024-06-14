package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/models"
)

// HelloHandler handles the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}

// UsersHandler handles the /users endpoint
func UsersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    users := []models.User{
        {ID: 1, Name: "John Doe"},
        {ID: 2, Name: "Jane Doe"},
    }

    json.NewEncoder(w).Encode(users)
}
