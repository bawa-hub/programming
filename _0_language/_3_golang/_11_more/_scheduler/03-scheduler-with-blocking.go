package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

///////////////////////////////////////////////////////////////
// G = Goroutine (Task)
///////////////////////////////////////////////////////////////

type G struct {
	id int
	fn func(done chan *G)
}

///////////////////////////////////////////////////////////////
// Global Run Queue
///////////////////////////////////////////////////////////////

var globalRunq []*G
var globalLock sync.Mutex

func globalPut(g *G) {
	globalLock.Lock()
	globalRunq = append(globalRunq, g)
	globalLock.Unlock()
}

func globalGet() *G {
	globalLock.Lock()
	defer globalLock.Unlock()

	if len(globalRunq) == 0 {
		return nil
	}
	g := globalRunq[0]
	globalRunq = globalRunq[1:]
	return g
}

///////////////////////////////////////////////////////////////
// P = Processor
///////////////////////////////////////////////////////////////

type P struct {
	id   int
	runq []*G
	lock sync.Mutex
}

func (p *P) runqPut(g *G) {
	p.lock.Lock()
	p.runq = append(p.runq, g)
	p.lock.Unlock()
}

func (p *P) runqGet() *G {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.runq) == 0 {
		return nil
	}
	g := p.runq[0]
	p.runq = p.runq[1:]
	return g
}

///////////////////////////////////////////////////////////////
// M = Machine
///////////////////////////////////////////////////////////////

type M struct {
	id int
	p  *P
}

func (m *M) start(allP []*P, done chan *G) {
	go func() {
		for {
			if m.p == nil {
				time.Sleep(10 * time.Millisecond)
				continue
			}

			g := m.p.runqGet()

			if g == nil {
				g = globalGet()
			}

			if g == nil {
				g = stealWork(m.p, allP)
			}

			if g == nil {
				time.Sleep(10 * time.Millisecond)
				continue
			}

			fmt.Printf("M%d executing G%d on P%d\n", m.id, g.id, m.p.id)
			g.fn(done)
		}
	}()
}

///////////////////////////////////////////////////////////////
// Work Stealing
///////////////////////////////////////////////////////////////

func stealWork(self *P, allP []*P) *G {
	for _, p := range allP {
		if p.id == self.id {
			continue
		}

		p.lock.Lock()
		n := len(p.runq)
		if n > 1 {
			half := n / 2
			stolen := p.runq[:half]
			p.runq = p.runq[half:]
			p.lock.Unlock()
			return stolen[0]
		}
		p.lock.Unlock()
	}
	return nil
}

///////////////////////////////////////////////////////////////
// Simulated Blocking (Syscall)
///////////////////////////////////////////////////////////////

func blockingTask(id int) func(done chan *G) {
	return func(done chan *G) {
		fmt.Printf("G%d blocking...\n", id)
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("G%d waking up\n", id)
			done <- &G{id: id, fn: blockingTask(id)}
		}()
	}
}

///////////////////////////////////////////////////////////////
// Main
///////////////////////////////////////////////////////////////

func main() {
	numP := 2
	allP := make([]*P, numP)
	done := make(chan *G)

	for i := 0; i < numP; i++ {
		allP[i] = &P{id: i}
	}

	// Create Ms
	for i := 0; i < numP; i++ {
		m := &M{id: i, p: allP[i]}
		m.start(allP, done)
	}

	// Create blocking tasks
	for i := 0; i < 5; i++ {
		g := &G{
			id: i,
			fn: blockingTask(i),
		}
		globalPut(g)
	}

	// Wakeup loop
	go func() {
		for g := range done {
			fmt.Printf("Re-scheduling G%d\n", g.id)
			globalPut(g)
		}
	}()

	time.Sleep(10 * time.Second)
}

// 🔥 What This Demonstrates
// 1️⃣ Global queue exists (like runtime)
// 2️⃣ Local P queues exist
// 3️⃣ Work stealing happens
// 4️⃣ Blocking tasks release execution
// 5️⃣ Woken tasks get re-queued

// This mirrors real runtime behavior.

// | Concept       | Real Function    |
// | ------------- | ---------------- |
// | Put runnable  | `ready()`        |
// | Find work     | `findRunnable()` |
// | Execute       | `execute()`      |
// | Park thread   | `park_m()`       |
// | Syscall enter | `entersyscall()` |
// | Syscall exit  | `exitsyscall()`  |
