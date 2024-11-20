package rate_limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate       int        // Number of tokens allowed per second
	bucketSize int        // Size of the bucket
	tokens     int        // Current number of tokens in the bucket
	lastRefill time.Time  // Time of the last token refill
	mu         sync.Mutex // Mutex to ensure thread safety
}

func NewTokenBucket(rate, bucketSize int, refillInterval time.Duration) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		bucketSize: bucketSize,
		tokens:     bucketSize, // Start with a full bucket
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Calculate the number of tokens to add based on the time elapsed since the last refill
	elapsed := time.Since(tb.lastRefill)
	newTokens := int(elapsed/time.Second) * tb.rate

	// Add new tokens but not more than the bucket size
	tb.tokens = min(tb.bucketSize, tb.tokens+newTokens)
	tb.lastRefill = time.Now()

	// Allow the request if there is at least 1 token in the bucket
	if tb.tokens > 0 {
		tb.tokens-- // Use 1 token
		return true
	}
	return false
}

// Helper function to get the minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
