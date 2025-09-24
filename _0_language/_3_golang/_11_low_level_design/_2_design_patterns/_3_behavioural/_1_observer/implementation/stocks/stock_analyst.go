package stocks

import "fmt"

type StockAnalyst struct {
	id string
}

func NewStockAnalyst(id string) *StockAnalyst {
	return &StockAnalyst{id: id}
}

func (sa *StockAnalyst) Update(stock *StockPrice) {
	fmt.Printf("Analyst %s: Analyzing %s\n", sa.id, stock)
	
	// Simple analysis logic
	if stock.ChangePercent > 0 {
		fmt.Printf("Analyst %s: %s is performing well (%.2f%% up)\n", 
			sa.id, stock.Symbol, stock.ChangePercent)
	} else {
		fmt.Printf("Analyst %s: %s is underperforming (%.2f%% down)\n", 
			sa.id, stock.Symbol, stock.ChangePercent)
	}
}

func (sa *StockAnalyst) GetID() string {
	return sa.id
}
