package main

import (
	"log"
	"net/http"
	"receipt-processor-challenge/handlers"
)

func main() {
    http.HandleFunc("/hello", handlers.HelloHandler)
    http.HandleFunc("/users", handlers.UsersHandler)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
