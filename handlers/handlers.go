package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge/models"

	"github.com/google/uuid"
)

// ReceiptHandler handles the /receipt endpoint
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var receipt models.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    // Generate a unique ID for the receipt
    receipt.ID = uuid.New().String()

    response := map[string]string{
        "id": receipt.ID,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}



