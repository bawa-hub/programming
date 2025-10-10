package payment

import (
	"fmt"
	"time"
)

type CreditCardStrategy struct {
	cardNumber string
	expiryDate string
	cvv        string
}

func NewCreditCardStrategy(cardNumber, expiryDate, cvv string) *CreditCardStrategy {
	return &CreditCardStrategy{
		cardNumber: cardNumber,
		expiryDate: expiryDate,
		cvv:        cvv,
	}
}

func (ccs *CreditCardStrategy) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Processing credit card payment: $%.2f %s\n", amount, currency)
	fmt.Printf("Card: %s, Expiry: %s\n", ccs.cardNumber, ccs.expiryDate)
	
	// Simulate processing
	time.Sleep(100 * time.Millisecond)
	
	transactionID := fmt.Sprintf("cc_%d", time.Now().UnixNano())
	return transactionID, nil
}

func (ccs *CreditCardStrategy) GetName() string {
	return "Credit Card"
}

func (ccs *CreditCardStrategy) GetFee() float64 {
	return 0.029 // 2.9% fee
}
