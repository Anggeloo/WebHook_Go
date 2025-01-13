# WebHook in Go 🚀
This repository contains a simple implementation of a WebHook server built using the Go programming language. The server listens for incoming HTTP POST requests, processes JSON payloads, and responds with a simple message.

## Project Structure 🗂️
```
.
├── WebHook.go   # Main application file
└── README.md    # Project documentation
```

## Features ✨
- Handles HTTP POST requests to the `/webhook` endpoint.
- Parses and validates JSON payloads.
- Logs incoming messages to the console.
- Responds with a JSON message.

## Prerequisites ✅

Before running the application, ensure you have the following installed:

- Go (version 1.16 or later)

## How to Run 🏃‍♂️

1. Clone the repository:

   ```bash
   git clone https://github.com/Anggeloo/WebHook_Go.git
   cd WebHook_Go
   ```

2. Build and run the application:

   ```bash
   go run WebHook.go
   ```

3. The server will start on `http://localhost:8080`.

## Testing the WebHook 🧪

You can test the WebHook using tools like [Postman](https://www.postman.com/) or `curl`.

### Example Request 📤

Send a POST request to `http://localhost:8080/webhook` with a JSON payload:

```bash
curl -X POST http://localhost:8080/webhook \
-H "Content-Type: application/json" \
-d '{"message": "Hello from client"}'
```

### Example Request with PowerShell 🖥️
If you're using PowerShell, use the following command:

```powershell
Invoke-WebRequest -Uri http://localhost:8080/webhook `
    -Method POST `
    -ContentType "application/json" `
    -Body '{"message": "Hola desde PowerShell"}'
```

### Example Response 📥
```json
{
  "response": "Hello World"
}
```

### Console Output 🖥️
The server logs the received message:

```plaintext
Message received: Hello from client
```