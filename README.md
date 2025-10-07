# 🚀 Go Web Server Learning Project

A comprehensive web server built in Go to demonstrate core language concepts and web development patterns.

## 🎯 What You'll Learn

This project demonstrates essential Go concepts:

- **Package System**: How Go organizes code into packages
- **Structs & JSON**: Data structures and serialization
- **HTTP Handlers**: Building web endpoints
- **Middleware**: Cross-cutting concerns like logging and CORS
- **Error Handling**: Go's explicit error handling approach
- **Concurrency**: Automatic concurrent request handling with goroutines

## 🚀 Quick Start

1. **Run the server:**
   ```bash
   go run main.go
   ```

2. **Visit the application:**
   - Home page: http://localhost:8080
   - About page: http://localhost:8080/about

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Interactive home page with API testing |
| GET | `/about` | Detailed explanation of Go concepts |
| GET | `/api/users` | Get all users (JSON response) |
| GET | `/api/health` | Health check endpoint |
| POST | `/api/echo` | Echo back JSON data |

## 🧪 Testing the API

### Get Users
```bash
curl http://localhost:8080/api/users
```

### Health Check
```bash
curl http://localhost:8080/api/health
```

### Echo Test
```bash
curl -X POST http://localhost:8080/api/echo \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello Go!", "test": true}'
```

## 🏗️ Project Structure

```
golang-web-server/
├── main.go          # Main server code with all handlers
├── go.mod          # Go module definition
└── README.md       # This file
```

## 🔍 Key Go Concepts Explained

### 1. Package Declaration
```go
package main  // Entry point package
```

### 2. Struct Definition with JSON Tags
```go
type User struct {
    ID   int    `json:"id"`     // JSON tag controls serialization
    Name string `json:"name"`
}
```

### 3. HTTP Handler Function Signature
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // w: write response
    // r: incoming request
}
```

### 4. Middleware Pattern
```go
func middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Do something before
        next.ServeHTTP(w, r)  // Call next handler
        // Do something after
    }
}
```

### 5. Error Handling
```go
if err != nil {
    log.Printf("Error occurred: %v", err)
    return
}
```

### 6. JSON Encoding/Decoding
```go
// Encode to JSON
json.NewEncoder(w).Encode(data)

// Decode from JSON
json.NewDecoder(r.Body).Decode(&data)
```

## 🌟 Features

- **Interactive Web Interface**: Test APIs directly from the browser
- **Comprehensive Logging**: Request/response logging middleware
- **CORS Support**: Cross-origin resource sharing enabled
- **JSON API**: RESTful endpoints with proper JSON responses
- **Error Handling**: Proper HTTP status codes and error responses
- **Concurrent**: Handles multiple requests simultaneously
- **Educational**: Extensive comments explaining Go concepts

## 🎓 Learning Path

1. **Start Simple**: Run the server and explore the home page
2. **Read the Code**: Study `main.go` to understand the structure
3. **Test APIs**: Use the interactive buttons or curl commands
4. **Modify & Experiment**: Try adding new endpoints or changing responses
5. **Learn More**: Visit `/about` for detailed concept explanations

## 🔧 Next Steps

Once you're comfortable with this server, try:

- Adding a database (SQLite or PostgreSQL)
- Implementing authentication with JWT
- Adding more complex routing with gorilla/mux
- Building a REST API with full CRUD operations
- Adding tests with Go's testing package
- Deploying to a cloud platform

## 📚 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)
- [Go Web Examples](https://gowebexamples.com/)

Happy learning! 🎉
