package exercies

import (
	"fmt"
	"sync"
	"time"
)

// Exercise 8: Object Pool
// Implement an object pool using sync.Pool.
func Exercise8() {
	fmt.Println("\nExercise 8: Object Pool")
	fmt.Println("=======================")
	
	var pool = sync.Pool{
		New: func() interface{} {
			return &g.Buffer{ID: time.Now().UnixNano()}
		},
	}
	
	var wg sync.WaitGroup
	
	// Get objects from pool
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Get buffer from pool
			buf := pool.Get().(*Buffer)
			fmt.Printf("Worker %d got buffer %d\n", id, buf.ID)
			
			// Use buffer
			buf.WriteString(fmt.Sprintf("Data from worker %d", id))
			time.Sleep(100 * time.Millisecond)
			
			// Put buffer back to pool
			buf.Reset()
			pool.Put(buf)
			fmt.Printf("Worker %d returned buffer %d to pool\n", id, buf.ID)
		}(i)
	}
	
	wg.Wait()
}

type Buffer struct {
	ID   int64
	Data string
}

func (b *Buffer) WriteString(s string) {
	b.Data += s
}

func (b *Buffer) Reset() {
	b.Data = ""
}
