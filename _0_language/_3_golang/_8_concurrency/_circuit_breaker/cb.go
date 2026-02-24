package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mu sync.Mutex

	state State

	failureCount     int
	failureThreshold int

	resetTimeout time.Duration
	lastFailure  time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failureThreshold: threshold,
		resetTimeout:     timeout,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()

	switch cb.state {
	case Open:
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			cb.state = HalfOpen
		} else {
			cb.mu.Unlock()
			return errors.New("circuit is open")
		}
	}

	cb.mu.Unlock()

	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()

		if cb.failureCount >= cb.failureThreshold {
			cb.state = Open
			fmt.Println("Circuit moved to OPEN")
		}

		return err
	}

	// Success
	cb.failureCount = 0
	cb.state = Closed
	return nil
}

func main() {
	cb := NewCircuitBreaker(3, 5*time.Second)

	for i := 0; i < 10; i++ {
		err := cb.Execute(func() error {
			fmt.Println("Calling external service")
			return errors.New("service failed")
		})

		fmt.Println("Result:", err)
		time.Sleep(1 * time.Second)
	}
}