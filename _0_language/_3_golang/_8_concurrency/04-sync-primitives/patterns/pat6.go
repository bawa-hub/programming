package patterns

import "sync/atomic"

// Advanced Pattern 6: Atomic Counter with Statistics
type AtomicCounter struct {
	value      int64
	increments int64
	decrements int64
	resets     int64
}

func (ac *AtomicCounter) Increment() int64 {
	atomic.AddInt64(&ac.increments, 1)
	return atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Decrement() int64 {
	atomic.AddInt64(&ac.decrements, 1)
	return atomic.AddInt64(&ac.value, -1)
}

func (ac *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&ac.value)
}

func (ac *AtomicCounter) Reset() int64 {
	atomic.AddInt64(&ac.resets, 1)
	return atomic.SwapInt64(&ac.value, 0)
}

func (ac *AtomicCounter) GetStats() (value, increments, decrements, resets int64) {
	return atomic.LoadInt64(&ac.value),
		atomic.LoadInt64(&ac.increments),
		atomic.LoadInt64(&ac.decrements),
		atomic.LoadInt64(&ac.resets)
}