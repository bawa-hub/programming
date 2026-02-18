package patterns

import (
	"context"
	"sync"
)

// Advanced Pattern 7: Select-based Context Manager
type SelectContextManager struct {
	contexts map[string]context.Context
	cancels  map[string]context.CancelFunc
	mu       sync.RWMutex
	requestCh chan ContextRequest
	responseCh chan context.Context
}

type ContextRequest struct {
	ID   string
	Type string // "get", "create", "cancel"
}

func NewSelectContextManager() *SelectContextManager {
	cm := &SelectContextManager{
		contexts:   make(map[string]context.Context),
		cancels:    make(map[string]context.CancelFunc),
		requestCh:  make(chan ContextRequest, 100),
		responseCh: make(chan context.Context, 100),
	}
	
	go cm.run()
	return cm
}

func (cm *SelectContextManager) run() {
	for {
		select {
		case req := <-cm.requestCh:
			cm.handleRequest(req)
		}
	}
}

func (cm *SelectContextManager) handleRequest(req ContextRequest) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	switch req.Type {
	case "get":
		if ctx, exists := cm.contexts[req.ID]; exists {
			cm.responseCh <- ctx
		} else {
			cm.responseCh <- nil
		}
	case "create":
		ctx, cancel := context.WithCancel(context.Background())
		cm.contexts[req.ID] = ctx
		cm.cancels[req.ID] = cancel
		cm.responseCh <- ctx
	case "cancel":
		if cancel, exists := cm.cancels[req.ID]; exists {
			cancel()
			delete(cm.contexts, req.ID)
			delete(cm.cancels, req.ID)
		}
		cm.responseCh <- nil
	}
}

func (cm *SelectContextManager) GetContext(id string) context.Context {
	cm.requestCh <- ContextRequest{ID: id, Type: "get"}
	return <-cm.responseCh
}

func (cm *SelectContextManager) CreateContext(id string) context.Context {
	cm.requestCh <- ContextRequest{ID: id, Type: "create"}
	return <-cm.responseCh
}

func (cm *SelectContextManager) CancelContext(id string) {
	cm.requestCh <- ContextRequest{ID: id, Type: "cancel"}
	<-cm.responseCh
}