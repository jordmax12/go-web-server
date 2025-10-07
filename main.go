package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a new HTTP multiplexer (router)
	mux := http.NewServeMux()
	
	// Register routes with middleware
	mux.HandleFunc("/", corsMiddleware(loggingMiddleware(homeHandler)))
	mux.HandleFunc("/about", corsMiddleware(loggingMiddleware(aboutHandler)))
	mux.HandleFunc("/api/users", corsMiddleware(loggingMiddleware(getUsersHandler)))
	mux.HandleFunc("/api/health", corsMiddleware(loggingMiddleware(healthHandler)))
	mux.HandleFunc("/api/echo", corsMiddleware(loggingMiddleware(echoHandler)))
	
	// Server configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	log.Println("🚀 Starting Go web server on http://localhost:8080")
	log.Println("📚 This server demonstrates key Go concepts:")
	log.Println("   • HTTP handling and routing")
	log.Println("   • Structs and JSON serialization")
	log.Println("   • Middleware patterns")
	log.Println("   • Error handling")
	log.Println("   • Concurrent request handling")
	log.Println("   • Proper code organization across files")
	
	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
