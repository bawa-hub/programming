package patterns

import (
	"sync"
	"sync/atomic"
)

// Advanced Pattern 2: Read-Write Lock with Priority
type PriorityRWMutex struct {
	mu           sync.RWMutex
	readerCount  int64
	writerCount  int64
	readerWait   int64
	writerWait   int64
	readerSignal chan struct{}
	writerSignal chan struct{}
}

func NewPriorityRWMutex() *PriorityRWMutex {
	return &PriorityRWMutex{
		readerSignal: make(chan struct{}),
		writerSignal: make(chan struct{}),
	}
}

func (prw *PriorityRWMutex) RLock() {
	atomic.AddInt64(&prw.readerWait, 1)
	prw.mu.RLock()
	atomic.AddInt64(&prw.readerCount, 1)
	atomic.AddInt64(&prw.readerWait, -1)
}

func (prw *PriorityRWMutex) RUnlock() {
	atomic.AddInt64(&prw.readerCount, -1)
	prw.mu.RUnlock()
}

func (prw *PriorityRWMutex) Lock() {
	atomic.AddInt64(&prw.writerWait, 1)
	prw.mu.Lock()
	atomic.AddInt64(&prw.writerCount, 1)
	atomic.AddInt64(&prw.writerWait, -1)
}

func (prw *PriorityRWMutex) Unlock() {
	atomic.AddInt64(&prw.writerCount, -1)
	prw.mu.Unlock()
}

func (prw *PriorityRWMutex) GetStats() (readers, writers, readerWait, writerWait int64) {
	return atomic.LoadInt64(&prw.readerCount),
		atomic.LoadInt64(&prw.writerCount),
		atomic.LoadInt64(&prw.readerWait),
		atomic.LoadInt64(&prw.writerWait)
}