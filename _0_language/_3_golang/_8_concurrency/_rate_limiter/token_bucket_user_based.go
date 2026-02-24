package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

/*
   -----------------------------
   Token Bucket Implementation
   -----------------------------
*/

type RateLimiter struct {
	rate       float64
	capacity   float64
	tokens     float64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewRateLimiter(rate float64, capacity float64) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()

	// Refill tokens
	rl.tokens += elapsed * rl.rate
	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}
	rl.lastRefill = now

	if rl.tokens >= 1 {
		rl.tokens -= 1
		return true
	}

	return false
}

/*
   ------------------------------------
   Per-User (Per-IP) Limiter Manager
   ------------------------------------
*/

type client struct {
	limiter  *RateLimiter
	lastSeen time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

// Get limiter for IP
func getLimiter(ip string) *RateLimiter {
	mu.Lock()
	defer mu.Unlock()

	c, exists := clients[ip]
	if !exists {
		limiter := NewRateLimiter(2, 5) // 2 req/sec, burst 5
		clients[ip] = &client{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}

	c.lastSeen = time.Now()
	return c.limiter
}

/*
   -----------------------------
   Cleanup Inactive Users
   -----------------------------
*/

func cleanupClients() {
	for {
		time.Sleep(1 * time.Minute)

		mu.Lock()
		for ip, c := range clients {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

/*
   -----------------------------
   Middleware
   -----------------------------
*/

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Unable to determine IP", http.StatusInternalServerError)
			return
		}

		limiter := getLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

/*
   -----------------------------
   Main
   -----------------------------
*/

func main() {

	go cleanupClients()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request allowed")
	})

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", rateLimitMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}