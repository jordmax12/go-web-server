# 🚀 Go Web Server Learning Project

A comprehensive web server built in Go to demonstrate core language concepts and web development patterns.

## 🚀 Quick Start

1. **Run the server:**
   ```bash
   go run main.go
   ```

2. **Visit the application:**
   - Home page: http://localhost:8080
   - About page: http://localhost:8080/about

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

Happy learning! 🎉
