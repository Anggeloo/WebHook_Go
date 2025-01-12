package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structure to handle WebHook data
type WebHookRequest struct {
	Message string `json:"message"`
}

func main() {
	// Define a handler for the WebHook
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Decode the JSON request body
		var req WebHookRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Error processing JSON", http.StatusBadRequest)
			return
		}

		// Print the received message
		fmt.Printf("Message received: %s\n", req.Message)

		// Respond with "Hello World"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"response": "Hello World",
		})
	})

	// Start the HTTP server
	fmt.Println("🚀 Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
