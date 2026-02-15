package main

import (
	"fmt"
	"time"
)

// Mapping to real go GMP
// G struct --> goroutine
// P.runq   --> P local run queue
// M.run()  --> OS thread
// NewScheduler(2) --> GOMAXPROCS(2)
// goroutines --> goroutines

type G struct {
	id   int
	work func()
}

type P struct {
	runq chan *G
}

type M struct {
	id int
}

type Scheduler struct {
	ps []*P
}

func NewScheduler(numP int) *Scheduler {
	ps := make([]*P, numP)
	for i := 0; i < numP; i++ {
		ps[i] = &P{
			runq: make(chan *G, 100),
		}
	}
	return &Scheduler{ps: ps}
}

func (s *Scheduler) Go(g *G) {
	s.ps[0].runq <- g
}

func (s *Scheduler) Run() {
	for i, p := range s.ps {
		m := &M{id: i}
		go m.run(p)
	}
}

func (m *M) run(p *P) {
	for g := range p.runq {
		fmt.Printf("M%d running G%d\n", m.id, g.id)
		g.work()
	}
}

func main() {
	s := NewScheduler(2) // 👈 THIS IS GOMAXPROCS

	s.Run()

	for i := 0; i < 5; i++ {
		id := i
		s.Go(&G{
			id: id,
			work: func() {
				time.Sleep(1 * time.Second)
			},
		})
	}

	time.Sleep(6 * time.Second)
}
