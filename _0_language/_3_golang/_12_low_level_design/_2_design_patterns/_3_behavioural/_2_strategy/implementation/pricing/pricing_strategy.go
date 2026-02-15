package pricing

type PricingStrategy interface {
	CalculatePrice(basePrice float64, quantity int) float64
	GetName() string
	GetDescription() string
}