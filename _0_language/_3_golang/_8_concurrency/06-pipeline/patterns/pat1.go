package patterns

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 1: Adaptive Pipeline

// ProcessedData represents data after processing
type ProcessedData struct {
	ID    int
	Value string
	Key   string
	Stage string
}

type AdaptivePipeline struct {
	stages       []*PipelineStage
	metrics      *AdaptiveMetrics
	mu           sync.RWMutex
	ctx          context.Context
	cancel       context.CancelFunc
}

type PipelineStage struct {
	name        string
	workers     int
	minWorkers  int
	maxWorkers  int
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	mu          sync.RWMutex
}

type AdaptiveMetrics struct {
	stageMetrics map[string]*StageMetrics
	mu           sync.RWMutex
}

type StageMetrics struct {
	processedItems int64
	processingTime time.Duration
	queueSize      int64
	workerUtilization float64
}

func NewAdaptivePipeline() *AdaptivePipeline {
	ctx, cancel := context.WithCancel(context.Background())
	
	pipeline := &AdaptivePipeline{
		stages:  make([]*PipelineStage, 0),
		metrics: &AdaptiveMetrics{
			stageMetrics: make(map[string]*StageMetrics),
		},
		ctx:    ctx,
		cancel: cancel,
	}
	
	go pipeline.adaptiveController()
	return pipeline
}

func (ap *AdaptivePipeline) AddStage(name string, minWorkers, maxWorkers int, processFunc func(ProcessedData) ProcessedData) {
	stage := &PipelineStage{
		name:        name,
		workers:     minWorkers,
		minWorkers:  minWorkers,
		maxWorkers:  maxWorkers,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
	}
	
	ap.stages = append(ap.stages, stage)
	ap.metrics.stageMetrics[name] = &StageMetrics{}
	
	// Start workers
	ap.startStageWorkers(stage)
}

func (ap *AdaptivePipeline) startStageWorkers(stage *PipelineStage) {
	for i := 0; i < stage.workers; i++ {
		go ap.stageWorker(stage, i)
	}
}

func (ap *AdaptivePipeline) stageWorker(stage *PipelineStage, workerID int) {
	for {
		select {
		case data, ok := <-stage.input:
			if !ok {
				return
			}
			
			start := time.Now()
			processed := stage.processFunc(data)
			duration := time.Since(start)
			
			// Update metrics
			ap.updateStageMetrics(stage.name, duration)
			
			// Send to next stage
			if len(ap.stages) > 0 {
				nextStageIndex := ap.getNextStageIndex(stage)
				if nextStageIndex < len(ap.stages) {
					ap.stages[nextStageIndex].input <- processed
				}
			}
			
		case <-ap.ctx.Done():
			return
		}
	}
}

func (ap *AdaptivePipeline) getNextStageIndex(currentStage *PipelineStage) int {
	for i, stage := range ap.stages {
		if stage == currentStage {
			return i + 1
		}
	}
	return -1
}

func (ap *AdaptivePipeline) updateStageMetrics(stageName string, duration time.Duration) {
	ap.metrics.mu.Lock()
	defer ap.metrics.mu.Unlock()
	
	if metrics, exists := ap.metrics.stageMetrics[stageName]; exists {
		metrics.processedItems++
		metrics.processingTime += duration
	}
}

func (ap *AdaptivePipeline) adaptiveController() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			ap.adjustStages()
		case <-ap.ctx.Done():
			return
		}
	}
}

func (ap *AdaptivePipeline) adjustStages() {
	ap.mu.Lock()
	defer ap.mu.Unlock()
	
	for _, stage := range ap.stages {
		queueSize := len(stage.input)
		utilization := ap.getStageUtilization(stage.name)
		
		if queueSize > 5 && stage.workers < stage.maxWorkers && utilization > 0.8 {
			// Add worker
			stage.workers++
			go ap.stageWorker(stage, stage.workers-1)
			fmt.Printf("Added worker to stage %s, total: %d\n", stage.name, stage.workers)
		} else if queueSize == 0 && stage.workers > stage.minWorkers && utilization < 0.3 {
			// Remove worker (simplified - in real implementation, you'd need to signal workers to stop)
			fmt.Printf("Would remove worker from stage %s, total: %d\n", stage.name, stage.workers)
		}
	}
}

func (ap *AdaptivePipeline) getStageUtilization(stageName string) float64 {
	ap.metrics.mu.RLock()
	defer ap.metrics.mu.RUnlock()
	
	if metrics, exists := ap.metrics.stageMetrics[stageName]; exists {
		return metrics.workerUtilization
	}
	return 0.0
}

func (ap *AdaptivePipeline) Submit(data ProcessedData) {
	if len(ap.stages) > 0 {
		ap.stages[0].input <- data
	}
}

func (ap *AdaptivePipeline) Close() {
	ap.cancel()
	for _, stage := range ap.stages {
		close(stage.input)
	}
}