package stocks

import (
	"fmt"
	"time"
)

// 1. STOCK PRICE MONITORING SYSTEM
type StockPrice struct {
	Symbol    string
	Price     float64
	Change    float64
	ChangePercent float64
	Timestamp time.Time
}

func (sp *StockPrice) String() string {
	return fmt.Sprintf("%s: $%.2f (%.2f, %.2f%%)", 
		sp.Symbol, sp.Price, sp.Change, sp.ChangePercent)
}
