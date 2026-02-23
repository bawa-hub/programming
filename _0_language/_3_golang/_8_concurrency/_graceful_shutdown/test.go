package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	log.Println("Starting server...")

	// Root context listens for SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// WaitGroup for background workers
	var wg sync.WaitGroup

	// ---------------------------
	// Background Worker
	// ---------------------------
	wg.Add(1)
	go func() {
		defer wg.Done()
		backgroundWorker(ctx)
	}()

	// ---------------------------
	// HTTP Handler
	// ---------------------------
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request")
		time.Sleep(10 * time.Second) // simulate long request
		fmt.Fprintln(w, "Request completed")
		log.Println("Request finished")
	})

	// ---------------------------
	// HTTP Server
	// ---------------------------
	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Start server in goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("HTTP server listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// ---------------------------
	// Wait for shutdown signal
	// ---------------------------
	<-ctx.Done()
	log.Println("Shutdown signal received")

	// Stop receiving further signals
	stop()

	// ---------------------------
	// Graceful shutdown with timeout
	// ---------------------------
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown error: %v", err)
	}

	// Wait for background tasks
	wg.Wait()

	log.Println("Server exited gracefully")
}

func backgroundWorker(ctx context.Context) {
	log.Println("Background worker started")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Background worker shutting down...")
			time.Sleep(3 * time.Second) // simulate cleanup
			log.Println("Background worker stopped")
			return
		case <-ticker.C:
			log.Println("Background worker running...")
		}
	}
}