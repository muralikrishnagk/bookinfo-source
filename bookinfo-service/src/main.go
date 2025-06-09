package main

import (
	"fmt"
	"net/http"
	"os"
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"service\": \"bookinfo-service\", \"version\": \"1.0.0\"}"))
	})

	fmt.Printf("Starting bookinfo-service on port %s\n", port)
	http.ListenAndServe(":"+port, logRequestHandler(http.DefaultServeMux))
}
