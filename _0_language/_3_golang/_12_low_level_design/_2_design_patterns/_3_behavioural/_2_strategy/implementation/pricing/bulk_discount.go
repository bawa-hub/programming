package pricing

import "fmt"

type BulkDiscountStrategy struct {
	discountPercent float64
	minQuantity     int
}

func NewBulkDiscountStrategy(discountPercent float64, minQuantity int) *BulkDiscountStrategy {
	return &BulkDiscountStrategy{
		discountPercent: discountPercent,
		minQuantity:     minQuantity,
	}
}

func (bds *BulkDiscountStrategy) CalculatePrice(basePrice float64, quantity int) float64 {
	total := basePrice * float64(quantity)
	if quantity >= bds.minQuantity {
		discount := total * (bds.discountPercent / 100)
		total -= discount
	}
	return total
}

func (bds *BulkDiscountStrategy) GetName() string {
	return "Bulk Discount"
}

func (bds *BulkDiscountStrategy) GetDescription() string {
	return fmt.Sprintf("%.1f%% discount for %d+ items", bds.discountPercent, bds.minQuantity)
}