package patterns

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Advanced Pattern 4: Pipeline with Caching
type CachedPipeline struct {
	stages  []*CachedStage
	caches  []*PipelineCache
	mu      sync.RWMutex
}

type CachedStage struct {
	name        string
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	cache       *PipelineCache
}

type PipelineCache struct {
	data   map[string]ProcessedData
	hits   int64
	misses int64
	mu     sync.RWMutex
}

func NewCachedPipeline() *CachedPipeline {
	return &CachedPipeline{
		stages: make([]*CachedStage, 0),
		caches: make([]*PipelineCache, 0),
	}
}

func (cp *CachedPipeline) AddStage(name string, processFunc func(ProcessedData) ProcessedData) {
	cache := &PipelineCache{
		data: make(map[string]ProcessedData),
	}
	
	stage := &CachedStage{
		name:        name,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
		cache:       cache,
	}
	
	// Start stage worker
	go cp.stageWorker(stage)
	
	cp.stages = append(cp.stages, stage)
	cp.caches = append(cp.caches, cache)
}

func (cp *CachedPipeline) stageWorker(stage *CachedStage) {
	for data := range stage.input {
		// Check cache first
		cacheKey := fmt.Sprintf("%s_%s", stage.name, data.Key)
		if cached, found := stage.cache.Get(cacheKey); found {
			stage.output <- cached
		} else {
			processed := stage.processFunc(data)
			stage.cache.Set(cacheKey, processed)
			stage.output <- processed
		}
	}
}

func (pc *PipelineCache) Get(key string) (ProcessedData, bool) {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	
	if value, found := pc.data[key]; found {
		atomic.AddInt64(&pc.hits, 1)
		return value, true
	}
	atomic.AddInt64(&pc.misses, 1)
	return ProcessedData{}, false
}

func (pc *PipelineCache) Set(key string, value ProcessedData) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.data[key] = value
}

func (pc *PipelineCache) GetHits() int64 {
	return atomic.LoadInt64(&pc.hits)
}

func (pc *PipelineCache) GetMisses() int64 {
	return atomic.LoadInt64(&pc.misses)
}

func (cp *CachedPipeline) Submit(data ProcessedData) {
	if len(cp.stages) > 0 {
		cp.stages[0].input <- data
	}
}

func (cp *CachedPipeline) Close() {
	for _, stage := range cp.stages {
		close(stage.input)
	}
}