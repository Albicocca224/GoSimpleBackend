package middleware

import "net/http"

// Auth protects routes by checking a "token" header
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "secret-minecraft-key" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the token is correct, move to the next handler
		next.ServeHTTP(w, r)
	})
}
