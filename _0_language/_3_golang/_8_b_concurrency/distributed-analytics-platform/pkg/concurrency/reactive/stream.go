package reactive

import (
	"context"
	"sync"
	"time"
)

// Stream represents a reactive stream
type Stream[T any] struct {
	source    chan T
	operators []Operator[T]
	mu        sync.RWMutex
	stopCh    chan struct{}
	ctx       context.Context
	cancel    context.CancelFunc
	subscribers []chan T
	subMu     sync.RWMutex
}

// Operator represents a stream operator
type Operator[T any] interface {
	Process(input <-chan T) <-chan T
}

// NewStream creates a new reactive stream
func NewStream[T any]() *Stream[T] {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &Stream[T]{
		source:      make(chan T, 1000),
		operators:   make([]Operator[T], 0),
		stopCh:      make(chan struct{}),
		ctx:         ctx,
		cancel:      cancel,
		subscribers: make([]chan T, 0),
	}
}

// Emit emits a value to the stream
func (s *Stream[T]) Emit(value T) {
	select {
	case s.source <- value:
	case <-s.ctx.Done():
		return
	default:
		// Stream is full, drop value
	}
}

// EmitBatch emits multiple values to the stream
func (s *Stream[T]) EmitBatch(values []T) {
	for _, value := range values {
		s.Emit(value)
	}
}

// Subscribe subscribes to the stream
func (s *Stream[T]) Subscribe() <-chan T {
	s.subMu.Lock()
	defer s.subMu.Unlock()
	
	output := make(chan T, 1000)
	s.subscribers = append(s.subscribers, output)
	
	return output
}

// Map applies a map operation to the stream
func (s *Stream[T]) Map(mapper func(T) T) *Stream[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	operator := &MapOperator[T]{
		mapper: mapper,
		ctx:    s.ctx,
	}
	
	s.operators = append(s.operators, operator)
	return s
}

// Filter applies a filter operation to the stream
func (s *Stream[T]) Filter(predicate func(T) bool) *Stream[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	operator := &FilterOperator[T]{
		predicate: predicate,
		ctx:       s.ctx,
	}
	
	s.operators = append(s.operators, operator)
	return s
}

// Reduce applies a reduce operation to the stream
func (s *Stream[T]) Reduce(initial T, reducer func(T, T) T) *Stream[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	operator := &ReduceOperator[T]{
		initial: initial,
		reducer: reducer,
		ctx:     s.ctx,
	}
	
	s.operators = append(s.operators, operator)
	return s
}

// Take takes the first n values from the stream
func (s *Stream[T]) Take(n int) *Stream[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	operator := &TakeOperator[T]{
		count: n,
		ctx:   s.ctx,
	}
	
	s.operators = append(s.operators, operator)
	return s
}

// Skip skips the first n values from the stream
func (s *Stream[T]) Skip(n int) *Stream[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	operator := &SkipOperator[T]{
		count: n,
		ctx:   s.ctx,
	}
	
	s.operators = append(s.operators, operator)
	return s
}

// Distinct removes duplicate values from the stream
func (s *Stream[T]) Distinct() *Stream[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	operator := &DistinctOperator[T]{
		seen: make(map[interface{}]bool),
		ctx:  s.ctx,
	}
	
	s.operators = append(s.operators, operator)
	return s
}

// Window creates a windowed stream
func (s *Stream[T]) Window(size int, interval time.Duration) *Stream[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	operator := &WindowOperator[T]{
		size:     size,
		interval: interval,
		ctx:      s.ctx,
	}
	
	s.operators = append(s.operators, operator)
	return s
}

// Start starts the stream processing
func (s *Stream[T]) Start() {
	go s.process()
}

// Stop stops the stream processing
func (s *Stream[T]) Stop() {
	s.cancel()
	close(s.stopCh)
}

// Flush flushes the stream
func (s *Stream[T]) Flush() {
	// Close source to signal end of stream
	close(s.source)
}

// process processes the stream
func (s *Stream[T]) process() {
	defer close(s.stopCh)
	
	// Start with source
	var current <-chan T = s.source
	
	// Apply operators
	for _, operator := range s.operators {
		current = operator.Process(current)
	}
	
	// Distribute to subscribers
	go s.distribute(current)
}

// distribute distributes values to subscribers
func (s *Stream[T]) distribute(input <-chan T) {
	for value := range input {
		s.subMu.RLock()
		for _, subscriber := range s.subscribers {
			select {
			case subscriber <- value:
			case <-s.ctx.Done():
				return
			default:
				// Subscriber is full, skip
			}
		}
		s.subMu.RUnlock()
	}
}

// MapOperator represents a map operation
type MapOperator[T any] struct {
	mapper func(T) T
	ctx    context.Context
}

// Process processes the input stream
func (m *MapOperator[T]) Process(input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		for value := range input {
			select {
			case output <- m.mapper(value):
			case <-m.ctx.Done():
				return
			}
		}
	}()
	
	return output
}

// FilterOperator represents a filter operation
type FilterOperator[T any] struct {
	predicate func(T) bool
	ctx       context.Context
}

// Process processes the input stream
func (f *FilterOperator[T]) Process(input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		for value := range input {
			if f.predicate(value) {
				select {
				case output <- value:
				case <-f.ctx.Done():
					return
				}
			}
		}
	}()
	
	return output
}

// ReduceOperator represents a reduce operation
type ReduceOperator[T any] struct {
	initial T
	reducer func(T, T) T
	ctx     context.Context
}

// Process processes the input stream
func (r *ReduceOperator[T]) Process(input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		
		acc := r.initial
		for value := range input {
			acc = r.reducer(acc, value)
		}
		
		select {
		case output <- acc:
		case <-r.ctx.Done():
			return
		}
	}()
	
	return output
}

// TakeOperator represents a take operation
type TakeOperator[T any] struct {
	count int
	ctx   context.Context
}

// Process processes the input stream
func (t *TakeOperator[T]) Process(input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		
		taken := 0
		for value := range input {
			if taken >= t.count {
				return
			}
			
			select {
			case output <- value:
				taken++
			case <-t.ctx.Done():
				return
			}
		}
	}()
	
	return output
}

// SkipOperator represents a skip operation
type SkipOperator[T any] struct {
	count int
	ctx   context.Context
}

// Process processes the input stream
func (s *SkipOperator[T]) Process(input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		
		skipped := 0
		for value := range input {
			if skipped < s.count {
				skipped++
				continue
			}
			
			select {
			case output <- value:
			case <-s.ctx.Done():
				return
			}
		}
	}()
	
	return output
}

// DistinctOperator represents a distinct operation
type DistinctOperator[T any] struct {
	seen map[interface{}]bool
	mu   sync.Mutex
	ctx  context.Context
}

// Process processes the input stream
func (d *DistinctOperator[T]) Process(input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		for value := range input {
			d.mu.Lock()
			if !d.seen[value] {
				d.seen[value] = true
				d.mu.Unlock()
				
				select {
				case output <- value:
				case <-d.ctx.Done():
					return
				}
			} else {
				d.mu.Unlock()
			}
		}
	}()
	
	return output
}

// WindowOperator represents a window operation
type WindowOperator[T any] struct {
	size     int
	interval time.Duration
	ctx      context.Context
}

// Process processes the input stream
func (w *WindowOperator[T]) Process(input <-chan T) <-chan T {
	output := make(chan T)
	
	go func() {
		defer close(output)
		
		window := make([]T, 0, w.size)
		ticker := time.NewTicker(w.interval)
		defer ticker.Stop()
		
		for {
			select {
			case value, ok := <-input:
				if !ok {
					// Flush remaining window
					if len(window) > 0 {
						for _, v := range window {
							select {
							case output <- v:
							case <-w.ctx.Done():
								return
							}
						}
					}
					return
				}
				
				window = append(window, value)
				if len(window) >= w.size {
					// Flush window
					for _, v := range window {
						select {
						case output <- v:
						case <-w.ctx.Done():
							return
						}
					}
					window = window[:0]
				}
				
			case <-ticker.C:
				// Flush window on interval
				if len(window) > 0 {
					for _, v := range window {
						select {
						case output <- v:
						case <-w.ctx.Done():
							return
						}
					}
					window = window[:0]
				}
				
			case <-w.ctx.Done():
				return
			}
		}
	}()
	
	return output
}

// Merge merges multiple streams into one
func Merge[T any](streams ...*Stream[T]) *Stream[T] {
	merged := NewStream[T]()
	
	for _, stream := range streams {
		go func(s *Stream[T]) {
			for value := range s.source {
				merged.Emit(value)
			}
		}(stream)
	}
	
	return merged
}

// Zip zips multiple streams together
func Zip[T any](streams ...*Stream[T]) *Stream[T] {
	zipped := NewStream[T]()
	
	go func() {
		defer zipped.Stop()
		
		channels := make([]<-chan T, len(streams))
		for i, stream := range streams {
			channels[i] = stream.source
		}
		
		for {
			values := make([]T, len(channels))
			allOk := true
			
			for i, ch := range channels {
				value, ok := <-ch
				if !ok {
					allOk = false
					break
				}
				values[i] = value
			}
			
			if !allOk {
				break
			}
			
			// Emit all values as a slice
			// This is a simplified version - in practice you'd want to handle this differently
			for _, value := range values {
				zipped.Emit(value)
			}
		}
	}()
	
	return zipped
}
