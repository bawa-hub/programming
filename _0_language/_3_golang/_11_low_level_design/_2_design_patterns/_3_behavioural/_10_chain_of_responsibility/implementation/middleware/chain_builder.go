package middleware


type MiddlewareChainBuilder struct {
    middlewares []Middleware
}

func NewMiddlewareChainBuilder() *MiddlewareChainBuilder {
    return &MiddlewareChainBuilder{
        middlewares: make([]Middleware, 0),
    }
}

func (b *MiddlewareChainBuilder) Add(m Middleware) *MiddlewareChainBuilder {
    b.middlewares = append(b.middlewares, m)
    return b
}

func (b *MiddlewareChainBuilder) Build() Middleware {
    if len(b.middlewares) == 0 {
        return nil
    }
    for i := 0; i < len(b.middlewares)-1; i++ {
        b.middlewares[i].SetNext(b.middlewares[i+1])
    }
    return b.middlewares[0]
}


