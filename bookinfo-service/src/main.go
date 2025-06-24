package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Build-time injected variables
var (
	version   = "dev"    // Default value, will be overridden by ldflags
	commit    = "none"   // Default value, will be overridden by ldflags
	buildDate = "unknown" // Default value, will be overridden by ldflags
)

func logRequestHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("bookinfo-service: Received request: %s %s\n", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	// Standard health endpoint for basic health checks
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Kubernetes readiness probe endpoint
	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ready"))
	})

	// Kubernetes liveness probe endpoint
	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Alive"))
	})

	// Standardized API info endpoint
	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		// Create a map to hold basic service info
		info := map[string]interface{}{
			"service":   "bookinfo-service",
			"version":   version,
			"commit":    commit,
			"buildDate": buildDate,
		}

		// Create a map to hold all environment variables
		env := make(map[string]string)
		
		// Iterate over all environment variables
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if len(pair) == 2 {
				env[pair[0]] = pair[1]
			}
		}
		
		// Add env vars to the info map
		info["ENV"] = env

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(info)
	})

	// Keep the old endpoint temporarily for backward compatibility
	http.HandleFunc("/api/service/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("WARNING: /api/service/info is deprecated, use /api/info instead")
		http.Redirect(w, r, "/api/info", http.StatusTemporaryRedirect)
	})

	fmt.Printf("Starting bookinfo-service version: %s, commit: %s, buildDate: %s on port %s\n", version, commit, buildDate, port)
	http.ListenAndServe(":"+port, logRequestHandler(http.DefaultServeMux))
}
