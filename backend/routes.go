package main

import (
    "net/http"
)

// Define your additional routes here if needed
func setupRoutes() {
    http.HandleFunc("/api", apiHandler)
}

// Simple API handler example
func apiHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"message": "Hello from the API!"}`))
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}
