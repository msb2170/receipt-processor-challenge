package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"receipt-processor-challenge/models"
)

// CalculatePoints calculates points for a receipt based on the specified rules
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += len(strings.TrimSpace(receipt.Retailer))

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil {
		totalCents := totalFloat * 100
		if math.Mod(totalCents, 100) == 0 {
			points += 50
		}
	} else {
		log.Printf("Error parsing total: %v", err)
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if totalFloat != 0 && math.Mod(totalFloat*100, 25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += len(receipt.Items) / 2 * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

var receipts = map[string]interface{}{} // Map to store receipts in memory

// ProcessReceiptHandler handles the /receipts/process endpoint
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Generate a unique ID for the receipt
	id := uuid.New().String()

	receipts[id] = struct{}{} // Store the receipt in memory

	response := map[string]string{
		"id": id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// CalculatePointsHandler calculates points for a receipt based on the specified rules
func CalculatePointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Check if the receipt with the provided ID exists in memory
	if _, ok := receipts[id]; !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	
	points := CalculatePoints(models.Receipt{})

	response := map[string]int{
		"points": points,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
