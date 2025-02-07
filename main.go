package main

import (
	"log"
	"net/http"
	"receipt-processor/internal/handlers"
	"receipt-processor/internal/services"
	"receipt-processor/internal/storage"
)

func main() {
	store := storage.NewInMemoryStore()
	service := services.NewPointsService()
	handler := handlers.NewReceiptHandler(store, service)
	mux := http.NewServeMux()
	mux.HandleFunc("/receipts/process", handler.Process)
	mux.HandleFunc("/receipts/", handler.Points)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

