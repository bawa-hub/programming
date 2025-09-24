package main

import "fmt"

// 2. OPEN/CLOSED PRINCIPLE (OCP)
// Open for extension, closed for modification

// PaymentProcessor interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

// CreditCardProcessor implements PaymentProcessor
type CreditCardProcessor struct {
	CardNumber string
}

func (ccp *CreditCardProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return nil
}

// PayPalProcessor implements PaymentProcessor
type PayPalProcessor struct {
	Email string
}

func (ppp *PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	return nil
}

// PaymentService - closed for modification, open for extension
type PaymentService struct {
	processor PaymentProcessor
}

func (ps *PaymentService) ProcessPayment(amount float64) error {
	return ps.processor.ProcessPayment(amount)
}