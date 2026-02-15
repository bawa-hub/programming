package payment

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy: strategy}
}

func (pp *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	pp.strategy = strategy
}

func (pp *PaymentProcessor) ProcessPayment(amount float64, currency string) (string, error) {
	return pp.strategy.ProcessPayment(amount, currency)
}

func (pp *PaymentProcessor) GetFee(amount float64) float64 {
	return amount * pp.strategy.GetFee()
}

func (pp *PaymentProcessor) GetCurrentStrategy() string {
	return pp.strategy.GetName()
}