package payment

import (
	"fmt"
	"time"
)

type BankTransferStrategy struct {
	accountNumber string
	routingNumber string
}

func NewBankTransferStrategy(accountNumber, routingNumber string) *BankTransferStrategy {
	return &BankTransferStrategy{
		accountNumber: accountNumber,
		routingNumber: routingNumber,
	}
}

func (bts *BankTransferStrategy) ProcessPayment(amount float64, currency string) (string, error) {
	fmt.Printf("Processing bank transfer: $%.2f %s\n", amount, currency)
	fmt.Printf("Account: %s, Routing: %s\n", bts.accountNumber, bts.routingNumber)
	
	// Simulate processing
	time.Sleep(200 * time.Millisecond)
	
	transactionID := fmt.Sprintf("bt_%d", time.Now().UnixNano())
	return transactionID, nil
}

func (bts *BankTransferStrategy) GetName() string {
	return "Bank Transfer"
}

func (bts *BankTransferStrategy) GetFee() float64 {
	return 0.01 // 1% fee
}