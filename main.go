package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

// Sample data - in a real app, this would come from a database
var users = []User{
	{ID: 1, Name: "Alice Johnson", Email: "alice@example.com", JoinDate: "2023-01-15"},
	{ID: 2, Name: "Bob Smith", Email: "bob@example.com", JoinDate: "2023-02-20"},
	{ID: 3, Name: "Carol Davis", Email: "carol@example.com", JoinDate: "2023-03-10"},
}

// Middleware function for logging requests
// This demonstrates Go's function types and closures
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)
		
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// CORS middleware to handle cross-origin requests
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}

// homeHandler serves the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Web Server Learning Project</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .endpoint { background: #f5f5f5; padding: 10px; margin: 10px 0; border-radius: 5px; }
        .method { color: #007acc; font-weight: bold; }
        code { background: #e8e8e8; padding: 2px 4px; border-radius: 3px; }
    </style>
</head>
<body>
    <h1>🚀 Go Web Server Learning Project</h1>
    <p>Welcome to your Go web server! This project demonstrates key Go concepts:</p>
    
    <h2>Available Endpoints:</h2>
    <div class="endpoint">
        <span class="method">GET</span> <code>/</code> - This home page
    </div>
    <div class="endpoint">
        <span class="method">GET</span> <code>/about</code> - About page with Go concepts
    </div>
    <div class="endpoint">
        <span class="method">GET</span> <code>/api/users</code> - Get all users (JSON)
    </div>
    <div class="endpoint">
        <span class="method">GET</span> <code>/api/health</code> - Health check endpoint
    </div>
    <div class="endpoint">
        <span class="method">POST</span> <code>/api/echo</code> - Echo back JSON data
    </div>
    
    <h2>Try the API:</h2>
    <button onclick="fetchUsers()">Fetch Users</button>
    <button onclick="testEcho()">Test Echo</button>
    <div id="result" style="margin-top: 20px; padding: 10px; background: #f0f0f0; border-radius: 5px;"></div>
    
    <script>
        async function fetchUsers() {
            try {
                const response = await fetch('/api/users');
                const data = await response.json();
                document.getElementById('result').innerHTML = '<pre>' + JSON.stringify(data, null, 2) + '</pre>';
            } catch (error) {
                document.getElementById('result').innerHTML = 'Error: ' + error.message;
            }
        }
        
        async function testEcho() {
            try {
                const response = await fetch('/api/echo', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ message: 'Hello from the frontend!', timestamp: new Date().toISOString() })
                });
                const data = await response.json();
                document.getElementById('result').innerHTML = '<pre>' + JSON.stringify(data, null, 2) + '</pre>';
            } catch (error) {
                document.getElementById('result').innerHTML = 'Error: ' + error.message;
            }
        }
    </script>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// aboutHandler explains Go concepts used in this server
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>About - Go Concepts</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .concept { background: #f9f9f9; padding: 15px; margin: 15px 0; border-left: 4px solid #007acc; }
        code { background: #e8e8e8; padding: 2px 4px; border-radius: 3px; }
        pre { background: #f5f5f5; padding: 10px; border-radius: 5px; overflow-x: auto; }
    </style>
</head>
<body>
    <h1>🧠 Go Concepts Demonstrated</h1>
    <a href="/">← Back to Home</a>
    
    <div class="concept">
        <h3>1. Package System</h3>
        <p>Go organizes code into packages. Our <code>main</code> package contains the entry point.</p>
        <pre>package main</pre>
    </div>
    
    <div class="concept">
        <h3>2. Structs and JSON Tags</h3>
        <p>Go structs define data structures. JSON tags control serialization:</p>
        <pre>type User struct {
    ID   int    ` + "`json:\"id\"`" + `
    Name string ` + "`json:\"name\"`" + `
}</pre>
    </div>
    
    <div class="concept">
        <h3>3. HTTP Handlers</h3>
        <p>Functions that handle HTTP requests have this signature:</p>
        <pre>func handler(w http.ResponseWriter, r *http.Request)</pre>
    </div>
    
    <div class="concept">
        <h3>4. Middleware Pattern</h3>
        <p>Functions that wrap other handlers for cross-cutting concerns:</p>
        <pre>func middleware(next http.HandlerFunc) http.HandlerFunc</pre>
    </div>
    
    <div class="concept">
        <h3>5. Error Handling</h3>
        <p>Go uses explicit error returns instead of exceptions:</p>
        <pre>if err != nil {
    log.Printf("Error: %v", err)
    return
}</pre>
    </div>
    
    <div class="concept">
        <h3>6. Goroutines & Concurrency</h3>
        <p>The HTTP server automatically handles requests concurrently using goroutines!</p>
    </div>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// getUsersHandler returns all users as JSON
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// healthHandler provides a simple health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "Server is healthy",
		Data: map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
			"uptime":    "Server is running",
			"version":   "1.0.0",
		},
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
		response := Response{
			Success: false,
			Message: "Invalid JSON data",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	response := Response{
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
	
	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
