package patterns

import (
	"sync"
	"sync/atomic"
)

// Advanced Pattern 8: Object Pool with Statistics
type StatsPool struct {
	pool     sync.Pool
	created  int64
	reused   int64
	returned int64
}

func NewStatsPool(newFunc func() interface{}) *StatsPool {
	return &StatsPool{
		pool: sync.Pool{New: newFunc},
	}
}

func (sp *StatsPool) Get() interface{} {
	obj := sp.pool.Get()
	if obj == nil {
		atomic.AddInt64(&sp.created, 1)
	} else {
		atomic.AddInt64(&sp.reused, 1)
	}
	return obj
}

func (sp *StatsPool) Put(obj interface{}) {
	atomic.AddInt64(&sp.returned, 1)
	sp.pool.Put(obj)
}

func (sp *StatsPool) GetStats() (created, reused, returned int64) {
	return atomic.LoadInt64(&sp.created),
		atomic.LoadInt64(&sp.reused),
		atomic.LoadInt64(&sp.returned)
}