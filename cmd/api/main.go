package main

import (
	"fmt"
	"net/http"

	"github.com/Albicocca224/GoSimpleBackend/internal/handlers"
	"github.com/Albicocca224/GoSimpleBackend/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// 1. Define your base handler
	taskFunc := http.HandlerFunc(handlers.TaskHandler)

	// 2. Wrap the handler with the Auth middleware
	// Now, every request to /task must pass through Auth first
	mux.Handle("/tasks/", middleware.Auth(taskFunc))

	fmt.Println("API running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}
