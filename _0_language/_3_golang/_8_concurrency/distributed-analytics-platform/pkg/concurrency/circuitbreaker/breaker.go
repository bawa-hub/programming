package circuitbreaker

import (
	"sync"
	"time"
)

// CircuitBreaker represents a circuit breaker
type CircuitBreaker struct {
	name            string
	maxRequests     uint32
	interval        time.Duration
	timeout         time.Duration
	readyToTrip     func(counts Counts) bool
	onStateChange   func(name string, from State, to State)
	
	mutex      sync.Mutex
	state      State
	generation uint64
	counts     Counts
	expiry     time.Time
}

// State represents the state of the circuit breaker
type State int

const (
	StateClosed State = iota
	StateHalfOpen
	StateOpen
)

// Counts represents the counts for the circuit breaker
type Counts struct {
	Requests             uint32
	TotalSuccesses       uint32
	TotalFailures        uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
}

// Config represents circuit breaker configuration
type Config struct {
	Name            string
	MaxRequests     uint32
	Interval        time.Duration
	Timeout         time.Duration
	ReadyToTrip     func(counts Counts) bool
	OnStateChange   func(name string, from State, to State)
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(name string, maxRequests uint32, interval, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		name:        name,
		maxRequests: maxRequests,
		interval:    interval,
		timeout:     timeout,
		readyToTrip: defaultReadyToTrip,
		state:       StateClosed,
	}
}

// NewCircuitBreakerWithConfig creates a new circuit breaker with custom configuration
func NewCircuitBreakerWithConfig(config Config) *CircuitBreaker {
	cb := &CircuitBreaker{
		name:          config.Name,
		maxRequests:   config.MaxRequests,
		interval:      config.Interval,
		timeout:       config.Timeout,
		readyToTrip:   config.ReadyToTrip,
		onStateChange: config.OnStateChange,
		state:         StateClosed,
	}
	
	if cb.readyToTrip == nil {
		cb.readyToTrip = defaultReadyToTrip
	}
	
	return cb
}

// Execute executes a function with circuit breaker protection
func (cb *CircuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	generation, err := cb.beforeRequest()
	if err != nil {
		return nil, err
	}
	
	defer func() {
		e := recover()
		if e != nil {
			cb.afterRequest(generation, false)
			panic(e)
		}
	}()
	
	result, err := req()
	cb.afterRequest(generation, err == nil)
	return result, err
}

// ExecuteWithFallback executes a function with fallback
func (cb *CircuitBreaker) ExecuteWithFallback(req func() (interface{}, error), fallback func(error) (interface{}, error)) (interface{}, error) {
	result, err := cb.Execute(req)
	if err != nil {
		if fallback != nil {
			return fallback(err)
		}
	}
	return result, err
}

// beforeRequest checks if request should be allowed
func (cb *CircuitBreaker) beforeRequest() (uint64, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	now := time.Now()
	state, generation := cb.currentState(now)
	
	if state == StateOpen {
		return generation, ErrOpenState
	} else if state == StateHalfOpen && cb.counts.Requests >= cb.maxRequests {
		return generation, ErrTooManyRequests
	}
	
	cb.counts.onRequest()
	return generation, nil
}

// afterRequest records the result of a request
func (cb *CircuitBreaker) afterRequest(before uint64, success bool) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	now := time.Now()
	state, generation := cb.currentState(now)
	if generation != before {
		return
	}
	
	if success {
		cb.onSuccess(state, now)
	} else {
		cb.onFailure(state, now)
	}
}

// currentState returns the current state and generation
func (cb *CircuitBreaker) currentState(now time.Time) (State, uint64) {
	switch cb.state {
	case StateClosed:
		if !cb.expiry.IsZero() && cb.expiry.Before(now) {
			cb.toNewGeneration(now)
		}
	case StateOpen:
		if cb.expiry.Before(now) {
			cb.setState(StateHalfOpen, now)
		}
	}
	return cb.state, cb.generation
}

// onSuccess handles successful requests
func (cb *CircuitBreaker) onSuccess(state State, now time.Time) {
	switch state {
	case StateClosed:
		cb.counts.onSuccess()
	case StateHalfOpen:
		cb.counts.onSuccess()
		if cb.counts.ConsecutiveSuccesses >= cb.maxRequests {
			cb.setState(StateClosed, now)
		}
	}
}

// onFailure handles failed requests
func (cb *CircuitBreaker) onFailure(state State, now time.Time) {
	switch state {
	case StateClosed:
		cb.counts.onFailure()
		if cb.readyToTrip(cb.counts) {
			cb.setState(StateOpen, now)
		}
	case StateHalfOpen:
		cb.setState(StateOpen, now)
	}
}

// setState sets the state of the circuit breaker
func (cb *CircuitBreaker) setState(state State, now time.Time) {
	if cb.state == state {
		return
	}
	
	prev := cb.state
	cb.state = state
	
	cb.toNewGeneration(now)
	
	if cb.onStateChange != nil {
		cb.onStateChange(cb.name, prev, state)
	}
}

// toNewGeneration creates a new generation
func (cb *CircuitBreaker) toNewGeneration(now time.Time) {
	cb.generation++
	cb.counts = Counts{}
	
	var zero time.Time
	switch cb.state {
	case StateClosed:
		if cb.interval == 0 {
			cb.expiry = zero
		} else {
			cb.expiry = now.Add(cb.interval)
		}
	case StateOpen:
		cb.expiry = now.Add(cb.timeout)
	default: // StateHalfOpen
		cb.expiry = zero
	}
}

// State returns the current state
func (cb *CircuitBreaker) State() State {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	now := time.Now()
	state, _ := cb.currentState(now)
	return state
}

// Counts returns the current counts
func (cb *CircuitBreaker) Counts() Counts {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	now := time.Now()
	_, _ = cb.currentState(now)
	return cb.counts
}

// onRequest increments the request count
func (c *Counts) onRequest() {
	c.Requests++
}

// onSuccess increments the success count
func (c *Counts) onSuccess() {
	c.TotalSuccesses++
	c.ConsecutiveSuccesses++
	c.ConsecutiveFailures = 0
}

// onFailure increments the failure count
func (c *Counts) onFailure() {
	c.TotalFailures++
	c.ConsecutiveFailures++
	c.ConsecutiveSuccesses = 0
}

// defaultReadyToTrip is the default ready to trip function
func defaultReadyToTrip(counts Counts) bool {
	return counts.ConsecutiveFailures > 5
}

// String returns the string representation of the state
func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateHalfOpen:
		return "half-open"
	case StateOpen:
		return "open"
	default:
		return "unknown"
	}
}

// Circuit breaker errors
var (
	ErrOpenState        = &CircuitBreakerError{msg: "circuit breaker is open"}
	ErrTooManyRequests  = &CircuitBreakerError{msg: "too many requests"}
)

// CircuitBreakerError represents a circuit breaker error
type CircuitBreakerError struct {
	msg string
}

func (e *CircuitBreakerError) Error() string {
	return e.msg
}

// IsOpen returns true if the circuit breaker is open
func (cb *CircuitBreaker) IsOpen() bool {
	return cb.State() == StateOpen
}

// IsClosed returns true if the circuit breaker is closed
func (cb *CircuitBreaker) IsClosed() bool {
	return cb.State() == StateClosed
}

// IsHalfOpen returns true if the circuit breaker is half-open
func (cb *CircuitBreaker) IsHalfOpen() bool {
	return cb.State() == StateHalfOpen
}

// Reset resets the circuit breaker to closed state
func (cb *CircuitBreaker) Reset() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	cb.setState(StateClosed, time.Now())
}

// GetStats returns circuit breaker statistics
func (cb *CircuitBreaker) GetStats() CircuitBreakerStats {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	
	now := time.Now()
	state, _ := cb.currentState(now)
	
	return CircuitBreakerStats{
		Name:                cb.name,
		State:               state,
		Generation:          cb.generation,
		Requests:            cb.counts.Requests,
		TotalSuccesses:      cb.counts.TotalSuccesses,
		TotalFailures:       cb.counts.TotalFailures,
		ConsecutiveSuccesses: cb.counts.ConsecutiveSuccesses,
		ConsecutiveFailures:  cb.counts.ConsecutiveFailures,
		Expiry:              cb.expiry,
	}
}

// CircuitBreakerStats represents circuit breaker statistics
type CircuitBreakerStats struct {
	Name                string
	State               State
	Generation          uint64
	Requests            uint32
	TotalSuccesses      uint32
	TotalFailures       uint32
	ConsecutiveSuccesses uint32
	ConsecutiveFailures  uint32
	Expiry              time.Time
}
