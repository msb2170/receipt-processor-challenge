package main

import (
	"log"
	"net/http"
	"receipt-processor-challenge/handlers"
)

func main() {
    http.HandleFunc("/receipt/process", handlers.ProcessReceiptHandler)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
