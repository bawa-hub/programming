package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr: ":9090",
		Handler: mux,
	}

	go func() {
		fmt.Println("Server running on :8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server error: ", err)
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()
	fmt.Println("Shutdown signal received")

	// Stop receiving further signals
	stop()
}

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10*time.Second)
		json.NewEncoder(w).Encode("Hello World")
}