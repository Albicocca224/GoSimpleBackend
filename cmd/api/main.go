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

	// 2. Chain Middleware: Logging -> Auth -> TaskHandler
	// Requests go through Logging first, then Auth, then the Handler.
	protectedHandler := middleware.LoggingMiddleware(middleware.APIKeyMiddleware(taskHandler))

	// 3. Register Route
	// We use "/tasks" because parameters like ?id=1 are handled inside the function
	mux.Handle("/tasks", protectedHandler)

	fmt.Println("Server starting on port :8080...") // [cite: 16]
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
