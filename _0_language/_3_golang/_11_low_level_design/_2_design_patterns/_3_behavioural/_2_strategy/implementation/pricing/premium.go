package pricing

import "fmt"

type PremiumPricingStrategy struct {
	premiumMultiplier float64
}

func NewPremiumPricingStrategy(premiumMultiplier float64) *PremiumPricingStrategy {
	return &PremiumPricingStrategy{
		premiumMultiplier: premiumMultiplier,
	}
}

func (pps *PremiumPricingStrategy) CalculatePrice(basePrice float64, quantity int) float64 {
	return basePrice * float64(quantity) * pps.premiumMultiplier
}

func (pps *PremiumPricingStrategy) GetName() string {
	return "Premium Pricing"
}

func (pps *PremiumPricingStrategy) GetDescription() string {
	return fmt.Sprintf("Premium pricing with %.1fx multiplier", pps.premiumMultiplier)
}