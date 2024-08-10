package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Create a new http.Server struct
	server := &http.Server{
		Addr:    ":8080", // Define the address and port to listen on
		Handler: mux,     // Use the ServeMux as the server's handler
	}

	// Start the server using ListenAndServe
	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
