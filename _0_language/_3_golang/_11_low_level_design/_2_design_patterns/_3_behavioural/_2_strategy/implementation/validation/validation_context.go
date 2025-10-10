package validation

type ValidationContext struct {
	strategy ValidationStrategy
}

func NewValidationContext(strategy ValidationStrategy) *ValidationContext {
	return &ValidationContext{strategy: strategy}
}

func (vc *ValidationContext) SetStrategy(strategy ValidationStrategy) {
	vc.strategy = strategy
}

func (vc *ValidationContext) Validate(data string) (bool, string) {
	return vc.strategy.Validate(data)
}

func (vc *ValidationContext) GetCurrentStrategy() string {
	return vc.strategy.GetName()
}