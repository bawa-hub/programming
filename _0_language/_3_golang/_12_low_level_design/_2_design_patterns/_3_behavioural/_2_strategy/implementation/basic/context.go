package basic

type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(data interface{}) interface{} {
	return c.strategy.Execute(data)
}

func (c *Context) GetCurrentStrategy() string {
	return c.strategy.GetName()
}