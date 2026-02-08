package main

import (
	"fmt"
	"net/http"

	"github.com/Albicocca224/GoSimpleBackend/internal/handlers"
	"github.com/Albicocca224/GoSimpleBackend/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	taskHandler := http.HandlerFunc(handlers.TaskHandler)

	protectedHandler := middleware.LoggingMiddleware(middleware.APIKeyMiddleware(taskHandler))

	mux.Handle("/tasks", protectedHandler)

	fmt.Println("Server starting on port :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
