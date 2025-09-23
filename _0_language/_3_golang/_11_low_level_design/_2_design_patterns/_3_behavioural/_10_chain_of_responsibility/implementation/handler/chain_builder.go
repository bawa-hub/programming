package handler

type ChainBuilder struct {
	handlers []Handler
}

func NewChainBuilder() *ChainBuilder {
	return &ChainBuilder{
		handlers: make([]Handler, 0),
	}
}

func (cb *ChainBuilder) AddHandler(handler Handler) *ChainBuilder {
	cb.handlers = append(cb.handlers, handler)
	return cb
}

func (cb *ChainBuilder) Build() Handler {
	if len(cb.handlers) == 0 {
		return nil
	}
	
	// Link handlers in sequence
	for i := 0; i < len(cb.handlers)-1; i++ {
		cb.handlers[i].SetNext(cb.handlers[i+1])
	}
	
	return cb.handlers[0]
}
