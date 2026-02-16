package patterns

import (
	"sync"
	"time"
)

// Advanced Pattern 5: Goroutine with Backpressure
type BackpressureGoroutine struct {
	input    chan interface{}
	output   chan interface{}
	pressure int
	maxPressure int
	mu       sync.RWMutex
}

func NewBackpressureGoroutine(maxPressure int) *BackpressureGoroutine {
	return &BackpressureGoroutine{
		input:       make(chan interface{}),
		output:      make(chan interface{}),
		pressure:    0,
		maxPressure: maxPressure,
	}
}

func (bp *BackpressureGoroutine) Start() {
	go bp.process()
}

func (bp *BackpressureGoroutine) process() {
	for item := range bp.input {
		bp.mu.Lock()
		if bp.pressure >= bp.maxPressure {
			bp.mu.Unlock()
			// Drop item or block
			continue
		}
		bp.pressure++
		bp.mu.Unlock()
		
		// Process item
		time.Sleep(100 * time.Millisecond) // Simulate work
		
		select {
		case bp.output <- item:
		default:
			// Output channel is full
		}
		
		bp.mu.Lock()
		bp.pressure--
		bp.mu.Unlock()
	}
}

func (bp *BackpressureGoroutine) Send(item interface{}) bool {
	select {
	case bp.input <- item:
		return true
	default:
		return false // Backpressure applied
	}
}

func (bp *BackpressureGoroutine) Receive() <-chan interface{} {
	return bp.output
}