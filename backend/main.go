package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"orderguard/backend/database" // Import our new package
)

func main() {
	// 1. Connect to Database
	dbPool := database.ConnectDB()
	defer dbPool.Close() // Ensure it closes when app stops

	// 2. Set up HTTP Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "ok", "message": "OrderGuard API is running", "database": "connected"}`)
	})

	log.Printf("🚀 OrderGuard Backend starting on http://localhost:%s", port)
	
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}