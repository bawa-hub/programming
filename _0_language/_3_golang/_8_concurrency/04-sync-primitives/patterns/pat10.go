package patterns

// Advanced Pattern 10: Semaphore
type Semaphore struct {
	permits chan struct{}
}

func NewSemaphore(permits int) *Semaphore {
	s := &Semaphore{
		permits: make(chan struct{}, permits),
	}
	
	// Fill with permits
	for i := 0; i < permits; i++ {
		s.permits <- struct{}{}
	}
	
	return s
}

func (s *Semaphore) Acquire() {
	<-s.permits
}

func (s *Semaphore) Release() {
	select {
	case s.permits <- struct{}{}:
	default:
		// Semaphore is full
	}
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case <-s.permits:
		return true
	default:
		return false
	}
}