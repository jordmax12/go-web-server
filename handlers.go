package main

import (
	"encoding/json"
	"golang-web-server/html"
	"golang-web-server/models"
	"net/http"
	"time"
)

// homeHandler serves the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html.HomeHTML))
}

// aboutHandler explains Go concepts used in this server
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html.AboutHTML))
}

// getUsersHandler returns all users as JSON
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    GetAllUsers(), // Using function from models.go
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// healthHandler provides a simple health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Success: true,
		Message: "Server is healthy",
		Data:    CreateHealthData(), // Using function from models.go
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// echoHandler echoes back the JSON data sent to it
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var requestData map[string]interface{}
	
	// Decode JSON from request body
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		response := models.Response{
			Success: false,
			Message: "Invalid JSON data",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	response := models.Response{
		Success: true,
		Message: "Data echoed successfully",
		Data: map[string]interface{}{
			"received_at": time.Now().Format(time.RFC3339),
			"your_data":   requestData,
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
