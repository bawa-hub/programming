package stocks

import "fmt"

type StockTrader struct {
	id        string
	portfolio map[string]int
}

func NewStockTrader(id string) *StockTrader {
	return &StockTrader{
		id:        id,
		portfolio: make(map[string]int),
	}
}

func (st *StockTrader) Update(stock *StockPrice) {
	fmt.Printf("Trader %s: %s\n", st.id, stock)
	
	// Simple trading logic
	if stock.ChangePercent > 5.0 {
		fmt.Printf("Trader %s: BUY signal for %s (%.2f%% increase)\n", 
			st.id, stock.Symbol, stock.ChangePercent)
	} else if stock.ChangePercent < -5.0 {
		fmt.Printf("Trader %s: SELL signal for %s (%.2f%% decrease)\n", 
			st.id, stock.Symbol, stock.ChangePercent)
	}
}

func (st *StockTrader) GetID() string {
	return st.id
}