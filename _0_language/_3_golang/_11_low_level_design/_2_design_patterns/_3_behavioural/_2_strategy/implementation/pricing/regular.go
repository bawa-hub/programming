package pricing

type RegularPricingStrategy struct{}

func (rps *RegularPricingStrategy) CalculatePrice(basePrice float64, quantity int) float64 {
	return basePrice * float64(quantity)
}

func (rps *RegularPricingStrategy) GetName() string {
	return "Regular Pricing"
}

func (rps *RegularPricingStrategy) GetDescription() string {
	return "Standard pricing with no discounts"
}