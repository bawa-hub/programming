package patterns

import "sync"

// Advanced Pattern 2: Pipeline with Load Balancing
type LoadBalancedPipeline struct {
	stages    []*LoadBalancedStage
	balancers []*LoadBalancer
	mu        sync.RWMutex
}

type LoadBalancedStage struct {
	name        string
	workers     []*Worker
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	balancer    *LoadBalancer
}

type Worker struct {
	id          int
	input       chan ProcessedData
	output      chan ProcessedData
	processFunc func(ProcessedData) ProcessedData
	load        int64
	mu          sync.RWMutex
}

type LoadBalancer struct {
	workers []*Worker
	index   int64
	mu      sync.RWMutex
}

func NewLoadBalancedPipeline() *LoadBalancedPipeline {
	return &LoadBalancedPipeline{
		stages:    make([]*LoadBalancedStage, 0),
		balancers: make([]*LoadBalancer, 0),
	}
}

func (lbp *LoadBalancedPipeline) AddStage(name string, numWorkers int, processFunc func(ProcessedData) ProcessedData) {
	workers := make([]*Worker, numWorkers)
	balancer := &LoadBalancer{
		workers: workers,
		index:   0,
	}
	
	stage := &LoadBalancedStage{
		name:        name,
		workers:     workers,
		input:       make(chan ProcessedData, 100),
		output:      make(chan ProcessedData, 100),
		processFunc: processFunc,
		balancer:    balancer,
	}
	
	// Create workers
	for i := 0; i < numWorkers; i++ {
		worker := &Worker{
			id:          i,
			input:       make(chan ProcessedData, 10),
			output:      stage.output,
			processFunc: processFunc,
		}
		workers[i] = worker
		balancer.workers[i] = worker
		
		// Start worker
		go lbp.workerLoop(worker)
	}
	
	// Start load balancer
	go lbp.loadBalancerLoop(stage)
	
	lbp.stages = append(lbp.stages, stage)
	lbp.balancers = append(lbp.balancers, balancer)
}

func (lbp *LoadBalancedPipeline) workerLoop(worker *Worker) {
	for data := range worker.input {
		processed := worker.processFunc(data)
		
		// Update load
		worker.mu.Lock()
		worker.load++
		worker.mu.Unlock()
		
		worker.output <- processed
	}
}

func (lbp *LoadBalancedPipeline) loadBalancerLoop(stage *LoadBalancedStage) {
	for data := range stage.input {
		worker := stage.balancer.GetWorker()
		worker.input <- data
	}
}

func (lb *LoadBalancer) GetWorker() *Worker {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	
	// Round-robin selection
	worker := lb.workers[lb.index%int64(len(lb.workers))]
	lb.index++
	return worker
}

func (lbp *LoadBalancedPipeline) Submit(data ProcessedData) {
	if len(lbp.stages) > 0 {
		lbp.stages[0].input <- data
	}
}

func (lbp *LoadBalancedPipeline) Close() {
	for _, stage := range lbp.stages {
		close(stage.input)
		for _, worker := range stage.workers {
			close(worker.input)
		}
	}
}