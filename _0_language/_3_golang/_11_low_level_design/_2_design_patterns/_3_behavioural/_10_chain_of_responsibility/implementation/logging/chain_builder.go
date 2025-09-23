package logging


type LoggingChainBuilder struct {
    handlers []LogHandler
}

func NewLoggingChainBuilder() *LoggingChainBuilder {
    return &LoggingChainBuilder{handlers: make([]LogHandler, 0)}
}

func (b *LoggingChainBuilder) Add(h LogHandler) *LoggingChainBuilder {
    b.handlers = append(b.handlers, h)
    return b
}

func (b *LoggingChainBuilder) Build() LogHandler {
    if len(b.handlers) == 0 {
        return nil
    }
    for i := 0; i < len(b.handlers)-1; i++ {
        b.handlers[i].SetNext(b.handlers[i+1])
    }
    return b.handlers[0]
}


