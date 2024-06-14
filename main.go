package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"receipt-processor-challenge/handlers"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", handlers.CalculatePointsHandler).Methods("GET")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
