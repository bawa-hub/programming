package patterns

import (
	"fmt"
	"sync"
	"time"
)

// Advanced Pattern 5: Pipeline with Metrics and Monitoring
type MonitoredPipeline struct {
	stages   []*MonitoredStage
	metrics  *PipelineMetrics
	mu       sync.RWMutex
}

type MonitoredStage struct {
	name        string
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	metrics     *StageMetrics
}

type PipelineMetrics struct {
	totalItems     int64
	totalTime      time.Duration
	stageMetrics   map[string]*StageMetrics
	mu             sync.RWMutex
}

func NewMonitoredPipeline() *MonitoredPipeline {
	return &MonitoredPipeline{
		stages: make([]*MonitoredStage, 0),
		metrics: &PipelineMetrics{
			stageMetrics: make(map[string]*StageMetrics),
		},
	}
}

func (mp *MonitoredPipeline) AddStage(name string, processFunc func(ProcessedData) ProcessedData) {
	stageMetrics := &StageMetrics{}
	
	stage := &MonitoredStage{
		name:        name,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
		metrics:     stageMetrics,
	}
	
	// Start stage worker
	go mp.stageWorker(stage)
	
	mp.stages = append(mp.stages, stage)
	mp.metrics.stageMetrics[name] = stageMetrics
}

func (mp *MonitoredPipeline) stageWorker(stage *MonitoredStage) {
	for data := range stage.input {
		start := time.Now()
		processed := stage.processFunc(data)
		duration := time.Since(start)
		
		// Update metrics
		mp.updateStageMetrics(stage.name, duration)
		
		stage.output <- processed
	}
}

func (mp *MonitoredPipeline) updateStageMetrics(stageName string, duration time.Duration) {
	mp.metrics.mu.Lock()
	defer mp.metrics.mu.Unlock()
	
	mp.metrics.totalItems++
	mp.metrics.totalTime += duration
	
	if stageMetrics, exists := mp.metrics.stageMetrics[stageName]; exists {
		stageMetrics.processedItems++
		stageMetrics.processingTime += duration
	}
}

func (mp *MonitoredPipeline) GetMetrics() *PipelineMetrics {
	mp.metrics.mu.RLock()
	defer mp.metrics.mu.RUnlock()
	
	// Return a copy of metrics
	return &PipelineMetrics{
		totalItems:   mp.metrics.totalItems,
		totalTime:    mp.metrics.totalTime,
		stageMetrics: mp.metrics.stageMetrics,
	}
}

func (mp *MonitoredPipeline) Submit(data ProcessedData) {
	if len(mp.stages) > 0 {
		mp.stages[0].input <- data
	}
}

func (mp *MonitoredPipeline) Close() {
	for _, stage := range mp.stages {
		close(stage.input)
	}
}

// Demo function to run all advanced patterns
func RunAdvancedPatterns() {
	fmt.Println("🚀 Advanced Pipeline Patterns")
	fmt.Println("=============================")
	
	// Pattern 1: Adaptive Pipeline
	fmt.Println("\n1. Adaptive Pipeline:")
	adaptivePipeline := pNewAdaptivePipeline()
	
	// Add stages
	adaptivePipeline.AddStage("stage1", 2, 5, func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Adaptive Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	adaptivePipeline.AddStage("stage2", 2, 5, func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Adaptive Stage2: %s", data.Value),
			Key:   data.Key,
			Stage: "stage2",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		adaptivePipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Adaptive Item %d", i),
			Key:   fmt.Sprintf("adaptive_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(2 * time.Second)
	adaptivePipeline.Close()
	
	// Pattern 2: Load Balanced Pipeline
	fmt.Println("\n2. Load Balanced Pipeline:")
	loadBalancedPipeline := NewLoadBalancedPipeline()
	
	loadBalancedPipeline.AddStage("stage1", 3, func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Load Balanced Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		loadBalancedPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Load Balanced Item %d", i),
			Key:   fmt.Sprintf("load_balanced_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	loadBalancedPipeline.Close()
	
	// Pattern 3: Circuit Breaker Pipeline
	fmt.Println("\n3. Circuit Breaker Pipeline:")
	circuitBreakerPipeline := NewCircuitBreakerPipeline()
	
	circuitBreakerPipeline.AddStage("stage1", 3, 500*time.Millisecond, func(data ProcessedData) (ProcessedData, error) {
		time.Sleep(50 * time.Millisecond)
		if data.ID%4 == 0 {
			return ProcessedData{}, fmt.Errorf("stage1 failed for item %d", data.ID)
		}
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Circuit Breaker Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}, nil
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		circuitBreakerPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Circuit Breaker Item %d", i),
			Key:   fmt.Sprintf("circuit_breaker_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	circuitBreakerPipeline.Close()
	
	// Pattern 4: Cached Pipeline
	fmt.Println("\n4. Cached Pipeline:")
	cachedPipeline := NewCachedPipeline()
	
	cachedPipeline.AddStage("stage1", func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Cached Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		cachedPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Cached Item %d", i),
			Key:   fmt.Sprintf("cached_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	cachedPipeline.Close()
	
	// Pattern 5: Monitored Pipeline
	fmt.Println("\n5. Monitored Pipeline:")
	monitoredPipeline := NewMonitoredPipeline()
	
	monitoredPipeline.AddStage("stage1", func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Monitored Stage1: %s", data.Value),
			Key:   data.Key,
			Stage: "stage1",
		}
	})
	
	monitoredPipeline.AddStage("stage2", func(data ProcessedData) ProcessedData {
		time.Sleep(50 * time.Millisecond)
		return ProcessedData{
			ID:    data.ID,
			Value: fmt.Sprintf("Monitored Stage2: %s", data.Value),
			Key:   data.Key,
			Stage: "stage2",
		}
	})
	
	// Submit data
	for i := 0; i < 10; i++ {
		monitoredPipeline.Submit(ProcessedData{
			ID:    i,
			Value: fmt.Sprintf("Monitored Item %d", i),
			Key:   fmt.Sprintf("monitored_key%d", i),
			Stage: "input",
		})
	}
	
	time.Sleep(1 * time.Second)
	
	// Print metrics
	metrics := monitoredPipeline.GetMetrics()
	fmt.Printf("  Total Items Processed: %d\n", metrics.totalItems)
	fmt.Printf("  Total Processing Time: %v\n", metrics.totalTime)
	
	monitoredPipeline.Close()
	
	fmt.Println("\n✅ All advanced patterns completed!")
}