package main

import "time"

// User represents a user in our system
// This demonstrates Go structs and JSON tags
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JoinDate string `json:"join_date"`
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// HealthData represents health check response data
type HealthData struct {
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
	Version   string `json:"version"`
}

// Sample data - in a real app, this would come from a database
var users = []User{
	{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", JoinDate: "2023-01-15"},
	{ID: 2, Name: "Bob Smith", Email: "bob@example.com", JoinDate: "2023-02-20"},
	{ID: 3, Name: "Carol Davis", Email: "carol@example.com", JoinDate: "2023-03-10"},
}

// GetAllUsers returns all users (simulates database query)
func GetAllUsers() []User {
	return users
}

// CreateHealthData creates health check data
func CreateHealthData() HealthData {
	return HealthData{
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    "Server is running",
		Version:   "1.0.0",
	}
}
