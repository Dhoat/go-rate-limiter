package main

import (
	"log"
	"net/http" // Correct import path
	"ratelimiter/handler"
	"ratelimiter/internal/rate_limiter" // Correct import path
	"time"
)

func main() {
	// Set up a rate limiter, using the Token Bucket algorithm
	limiter := rate_limiter.NewTokenBucket(5, 1, time.Second) // 5 tokens per second
	handler := handler.NewRequestHandler(limiter)

	// Set up the HTTP server
	http.Handle("/", handler)

	// Start the server
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
