package main

import (
	"golang-web-server/models"
	"time"
)

// Sample data - in a real app, this would come from a database
var users = []models.User{
	{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", JoinDate: "2023-01-15"},
	{ID: 2, Name: "Bob Smith", Email: "bob@example.com", JoinDate: "2023-02-20"},
	{ID: 3, Name: "Carol Davis", Email: "carol@example.com", JoinDate: "2023-03-10"},
}

// GetAllUsers returns all users (simulates database query)
func GetAllUsers() []models.User {
	return users
}

// CreateHealthData creates health check data
func CreateHealthData() models.HealthData {
	return models.HealthData{
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    "Server is running",
		Version:   "1.0.0",
	}
}
