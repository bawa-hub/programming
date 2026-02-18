package patterns

import (
	"sync"
	"time"
)

// Advanced Pattern 7: Channel-based Context with Cancellation
type ChannelContext struct {
	done   chan struct{}
	cancel func()
	mu     sync.RWMutex
}

func NewChannelContext() *ChannelContext {
	ctx := &ChannelContext{
		done: make(chan struct{}),
	}
	
	ctx.cancel = func() {
		close(ctx.done)
	}
	
	return ctx
}

func (ctx *ChannelContext) Done() <-chan struct{} {
	return ctx.done
}

func (ctx *ChannelContext) Cancel() {
	ctx.cancel()
}

func (ctx *ChannelContext) WithTimeout(timeout time.Duration) *ChannelContext {
	newCtx := &ChannelContext{
		done: make(chan struct{}),
	}
	
	newCtx.cancel = func() {
		close(newCtx.done)
	}
	
	go func() {
		select {
		case <-time.After(timeout):
			newCtx.cancel()
		case <-ctx.done:
			newCtx.cancel()
		}
	}()
	
	return newCtx
}