package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	users = make(map[int]User)
	mu    sync.Mutex
	idSeq = 1
)

func main() {
	// Root context listens for SIGINT & SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/", userHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Start server in goroutine
	go func() {
		log.Println("Server running on :8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	log.Println("Shutdown signal received")

	// Stop receiving further signals
	stop()

	// Graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown failed: %v", err)
	} else {
		log.Println("Server exited gracefully")
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)
	case http.MethodGet:
		getAllUsers(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUser(w, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	u.ID = idSeq
	idSeq++
	users[u.ID] = u
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var list []User
	for _, u := range users {
		list = append(list, u)
	}

	json.NewEncoder(w).Encode(list)
}

func getUser(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	u, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var updated User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	updated.ID = id
	users[id] = updated

	json.NewEncoder(w).Encode(updated)
}

func deleteUser(w http.ResponseWriter, id int) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}

// 🧪 4️⃣ How to Test

// Create user:
// curl -X POST localhost:8080/users \
// -H "Content-Type: application/json" \
// -d '{"name":"Vikram","age":30}'

// Get all:
// curl localhost:8080/users

// Update:
// curl -X PUT localhost:8080/users/1 \
// -H "Content-Type: application/json" \
// -d '{"name":"Updated","age":35}'

// Delete:
// curl -X DELETE localhost:8080/users/1