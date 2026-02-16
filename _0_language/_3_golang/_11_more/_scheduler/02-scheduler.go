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
	fn func()
}

///////////////////////////////////////////////////////////////
// P = Processor (Owns run queue)
///////////////////////////////////////////////////////////////

type P struct {
	id    int
	runq  []*G
	lock  sync.Mutex
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
// M = Machine (Worker thread)
///////////////////////////////////////////////////////////////

type M struct {
	id int
	p  *P
}

func (m *M) start(allP []*P, wg *sync.WaitGroup) {
	go func() {
		for {
			g := m.p.runqGet()

			if g == nil {
				// Try to steal
				g = stealWork(m.p, allP)
			}

			if g == nil {
				// Nothing to do
				time.Sleep(10 * time.Millisecond)
				continue
			}

			fmt.Printf("M%d executing G%d on P%d\n", m.id, g.id, m.p.id)
			g.fn()
			wg.Done()
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
			// steal half
			half := n / 2
			stolen := p.runq[:half]
			p.runq = p.runq[half:]
			p.lock.Unlock()

			// return one of stolen tasks
			return stolen[0]
		}
		p.lock.Unlock()
	}
	return nil
}

///////////////////////////////////////////////////////////////
// Main
///////////////////////////////////////////////////////////////

func main() {
	rand.Seed(time.Now().UnixNano())

	numP := 2
	numTasks := 10

	var wg sync.WaitGroup
	wg.Add(numTasks)

	// Create Ps
	allP := make([]*P, numP)
	for i := 0; i < numP; i++ {
		allP[i] = &P{id: i}
	}

	// Create Ms (1 M per P)
	for i := 0; i < numP; i++ {
		m := &M{id: i, p: allP[i]}
		m.start(allP, &wg)
	}

	// Create Gs (Tasks)
	for i := 0; i < numTasks; i++ {
		g := &G{
			id: i,
			fn: func() {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			},
		}

		// Randomly assign to a P
		p := allP[rand.Intn(numP)]
		p.runqPut(g)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}

// 1️⃣ Each P has its own run queue
// 2️⃣ Each M owns exactly one P
// 3️⃣ If local queue empty → steal from another P
// 4️⃣ No global scheduling lock

// This mirrors Go runtime design.

// | Our Simulation | Real Go Runtime     |
// | -------------- | ------------------- |
// | G struct       | runtime.g           |
// | P struct       | runtime.p           |
// | M struct       | runtime.m           |
// | runq slice     | runq [256]guintptr  |
// | stealWork()    | runtime.stealWork() |
// | M.start()      | schedule() loop     |
