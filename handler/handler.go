package handler

import (
	"fmt"
	"net/http"
	"ratelimiter/internal/rate_limiter" // Import the rate_limiter package
)

type RequestHandler struct {
	limiter *rate_limiter.TokenBucket
}

func NewRequestHandler(limiter *rate_limiter.TokenBucket) *RequestHandler {
	return &RequestHandler{
		limiter: limiter,
	}
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if the request is allowed by the rate limiter
	if h.limiter.Allow() {
		// If allowed, process the request
		fmt.Fprintln(w, "Request allowed!")
	} else {
		// If not allowed, reject the request with a 429 Too Many Requests
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
	}
}
