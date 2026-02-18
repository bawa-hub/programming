package patterns

import "sync"

// Advanced Pattern 4: Once with Error Handling
type SafeOnce struct {
	once sync.Once
	err  error
}

func (so *SafeOnce) Do(fn func() error) error {
	so.once.Do(func() {
		so.err = fn()
	})
	return so.err
}