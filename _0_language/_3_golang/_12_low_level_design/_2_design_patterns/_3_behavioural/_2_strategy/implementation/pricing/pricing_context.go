package pricing

type PricingContext struct {
	strategy PricingStrategy
}

func NewPricingContext(strategy PricingStrategy) *PricingContext {
	return &PricingContext{strategy: strategy}
}

func (pc *PricingContext) SetStrategy(strategy PricingStrategy) {
	pc.strategy = strategy
}

func (pc *PricingContext) CalculatePrice(basePrice float64, quantity int) float64 {
	return pc.strategy.CalculatePrice(basePrice, quantity)
}

func (pc *PricingContext) GetCurrentStrategy() string {
	return pc.strategy.GetName()
}