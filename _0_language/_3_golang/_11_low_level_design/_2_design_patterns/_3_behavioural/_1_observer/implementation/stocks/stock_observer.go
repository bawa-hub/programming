package stocks


type StockObserver interface {
	Update(stock *StockPrice)
	GetID() string
}