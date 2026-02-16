package patterns

import (
	"sync"
	"time"
)

// Advanced Pattern 3: Goroutine with Heartbeat
type HeartbeatGoroutine struct {
	heartbeat chan time.Time
	done      chan bool
	wg        sync.WaitGroup
}

func NewHeartbeatGoroutine() *HeartbeatGoroutine {
	return &HeartbeatGoroutine{
		heartbeat: make(chan time.Time),
		done:      make(chan bool),
	}
}

func (h *HeartbeatGoroutine) Start() {
	h.wg.Add(1)
	go h.run()
}

func (h *HeartbeatGoroutine) run() {
	defer h.wg.Done()
	
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case t := <-ticker.C:
			select {
			case h.heartbeat <- t:
			default:
				// Heartbeat channel is full, skip
			}
		case <-h.done:
			return
		}
	}
}

func (h *HeartbeatGoroutine) GetHeartbeat() <-chan time.Time {
	return h.heartbeat
}

func (h *HeartbeatGoroutine) Stop() {
	close(h.done)
	h.wg.Wait()
}