package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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
        <h3>1. Package System & File Organization</h3>
        <p>Go organizes code into packages. Our <code>main</code> package is split across multiple files:</p>
        <pre>main.go      - Server setup and main()
models.go    - Data structures (User, Response)
handlers.go  - HTTP request handlers
middleware.go - Middleware functions</pre>
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
		Data:    GetAllUsers(), // Using function from models.go
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// healthHandler provides a simple health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
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
