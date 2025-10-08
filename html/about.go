package html

const AboutHTML = `
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
