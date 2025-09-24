package main

import "fmt"

// BankAccount demonstrates encapsulation
type BankAccount struct {
	accountNumber string
	balance      float64
	owner        string
	// private fields - cannot be accessed directly from outside the package
}

// Constructor function (Go doesn't have constructors, but we can create factory functions)
func NewBankAccount(accountNumber, owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		accountNumber: accountNumber,
		balance:      initialBalance,
		owner:        owner,
	}
}

// Public methods to access private data (encapsulation)
func (ba *BankAccount) GetBalance() float64 {
	return ba.balance
}

func (ba *BankAccount) GetAccountNumber() string {
	return ba.accountNumber
}

func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	ba.balance += amount
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > ba.balance {
		return fmt.Errorf("insufficient funds")
	}
	ba.balance -= amount
	return nil
}

func main() {
		fmt.Println("1. ENCAPSULATION:")
	account := NewBankAccount("12345", "John Doe", 1000.0)
	fmt.Printf("Initial balance: $%.2f\n", account.GetBalance())
	
	err := account.Deposit(500.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After deposit: $%.2f\n", account.GetBalance())
	}

	err = account.Withdraw(200.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("After withdrawal: $%.2f\n", account.GetBalance())
	}
	fmt.Println()

}