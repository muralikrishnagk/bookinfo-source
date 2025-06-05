package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/api/processor/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"service\": \"bookinfo-processor\", \"version\": \"1.0.0\"}"))
	})

	http.HandleFunc("/api/processor/process", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\": \"processed\", \"item_id\": \"example123\"}"))
	})

	fmt.Printf("Starting Bookinfo Processor service on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
