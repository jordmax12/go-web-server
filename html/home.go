package html

const HomeHTML = `
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
