package payment

type PaymentStrategy interface {
	ProcessPayment(amount float64, currency string) (string, error)
	GetName() string
	GetFee() float64
}
