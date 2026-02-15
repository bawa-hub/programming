package payment

import (
	"fmt"
	"time"
)

type PayPalStrategy struct {
	email    string
	password string
}

func NewPayPalStrategy(email, password string) *PayPalStrategy {
	return &PayPalStrategy{
		email:    email,
		password: password,
	}
}

func (pps *PayPalStrategy) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Processing PayPal payment: $%.2f %s\n", amount, currency)
	fmt.Printf("Email: %s\n", pps.email)
	
	// Simulate processing
	time.Sleep(150 * time.Millisecond)
	
	transactionID := fmt.Sprintf("pp_%d", time.Now().UnixNano())
	return transactionID, nil
}

func (pps *PayPalStrategy) GetName() string {
	return "PayPal"
}

func (pps *PayPalStrategy) GetFee() float64 {
	return 0.034 // 3.4% fee
}