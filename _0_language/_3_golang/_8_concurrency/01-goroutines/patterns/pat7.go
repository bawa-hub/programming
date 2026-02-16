package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 7: Goroutine with Graceful Shutdown
type GracefulGoroutine struct {
	shutdown chan struct{}
	done     chan struct{}
	wg       sync.WaitGroup
}

func NewGracefulGoroutine() *GracefulGoroutine {
	return &GracefulGoroutine{
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}
}

func (g *GracefulGoroutine) Start() {
	g.wg.Add(1)
	go g.run()
}

func (g *GracefulGoroutine) run() {
	defer g.wg.Done()
	defer close(g.done)
	
	for {
		select {
		case <-g.shutdown:
			fmt.Println("Goroutine: Received shutdown signal")
			return
		default:
			// Do work
			fmt.Println("Goroutine: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (g *GracefulGoroutine) Shutdown() {
	close(g.shutdown)
	g.wg.Wait()
}

func (g *GracefulGoroutine) Wait() {
	<-g.done
}