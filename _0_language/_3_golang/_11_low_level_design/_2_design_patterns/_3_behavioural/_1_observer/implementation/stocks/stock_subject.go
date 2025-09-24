package stocks



type StockSubject interface {
	Attach(observer StockObserver)
	Detach(observer StockObserver)
	Notify()
	UpdatePrice(symbol string, price float64)
	GetPrice(symbol string) *StockPrice
}