package stocks

import (
	"fmt"
	"sync"
	"time"
)



type StockMarket struct {
	observers []StockObserver
	stocks    map[string]*StockPrice
	mu        sync.RWMutex
}

func NewStockMarket() *StockMarket {
	return &StockMarket{
		observers: make([]StockObserver, 0),
		stocks:    make(map[string]*StockPrice),
	}
}

func (sm *StockMarket) Attach(observer StockObserver) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.observers = append(sm.observers, observer)
	fmt.Printf("Stock observer %s attached\n", observer.GetID())
}

func (sm *StockMarket) Detach(observer StockObserver) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	for i, obs := range sm.observers {
		if obs.GetID() == observer.GetID() {
			sm.observers = append(sm.observers[:i], sm.observers[i+1:]...)
			fmt.Printf("Stock observer %s detached\n", observer.GetID())
			break
		}
	}
}

func (sm *StockMarket) Notify() {
	sm.mu.RLock()
	observers := make([]StockObserver, len(sm.observers))
	copy(observers, sm.observers)
	sm.mu.RUnlock()
	
	for _, observer := range observers {
		// Notify about all stocks
		for _, stock := range sm.stocks {
			observer.Update(stock)
		}
	}
}

func (sm *StockMarket) UpdatePrice(symbol string, price float64) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	if existingStock, exists := sm.stocks[symbol]; exists {
		oldPrice := existingStock.Price
		existingStock.Price = price
		existingStock.Change = price - oldPrice
		existingStock.ChangePercent = (existingStock.Change / oldPrice) * 100
		existingStock.Timestamp = time.Now()
	} else {
		sm.stocks[symbol] = &StockPrice{
			Symbol:        symbol,
			Price:         price,
			Change:        0,
			ChangePercent: 0,
			Timestamp:     time.Now(),
		}
	}
	
	fmt.Printf("Stock price updated: %s\n", sm.stocks[symbol])
	sm.mu.Unlock()
	sm.Notify()
	sm.mu.Lock()
}

func (sm *StockMarket) GetPrice(symbol string) *StockPrice {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.stocks[symbol]
}