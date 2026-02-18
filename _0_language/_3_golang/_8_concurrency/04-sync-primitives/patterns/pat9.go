package patterns

import "sync"

// Advanced Pattern 9: Barrier Synchronization
type Barrier struct {
	mu      sync.Mutex
	cond    *sync.Cond
	count   int
	target  int
	phase   int
}

func NewBarrier(target int) *Barrier {
	b := &Barrier{target: target}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Barrier) Wait() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	phase := b.phase
	b.count++
	
	if b.count == b.target {
		b.count = 0
		b.phase++
		b.cond.Broadcast()
	} else {
		for phase == b.phase {
			b.cond.Wait()
		}
	}
	
	return phase
}