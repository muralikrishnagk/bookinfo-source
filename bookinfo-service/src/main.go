package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/api/service/info", func(w http.ResponseWriter, r *http.Request) {
		info := map[string]string{
			"service":   "bookinfo-service",
			"version":   version,
			"commit":    commit,
			"buildDate": buildDate,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(info)
	})

	fmt.Printf("Starting bookinfo-service version: %s, commit: %s, buildDate: %s on port %s\n", version, commit, buildDate, port)
	http.ListenAndServe(":"+port, logRequestHandler(http.DefaultServeMux))
}
