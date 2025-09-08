package main

import (
	"fmt"
	"sync"
	"time"
)

// Bank account with mutex protection
type BankAccount struct {
	mu      sync.RWMutex
	balance int
}

func NewBankAccount(initialBalance int) *BankAccount {
	return &BankAccount{balance: initialBalance}
}

func (ba *BankAccount) Deposit(amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	ba.balance += amount
	fmt.Printf("Deposited %d, new balance: %d\n", amount, ba.balance)
}

func (ba *BankAccount) Withdraw(amount int) bool {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	if ba.balance >= amount {
		ba.balance -= amount
		fmt.Printf("Withdrew %d, new balance: %d\n", amount, ba.balance)
		return true
	}
	fmt.Printf("Insufficient funds for withdrawal of %d\n", amount)
	return false
}

func (ba *BankAccount) GetBalance() int {
	ba.mu.RLock()
	defer ba.mu.RUnlock()
	return ba.balance
}

func main() {
	fmt.Println("=== Combined Sync Primitives Example ===")
	
	account := NewBankAccount(1000)
	var wg sync.WaitGroup
	
	// Start multiple depositors
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				account.Deposit(100)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	// Start multiple withdrawers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				account.Withdraw(50)
				time.Sleep(150 * time.Millisecond)
			}
		}(i)
	}
	
	// Start a balance checker
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			balance := account.GetBalance()
			fmt.Printf("Balance check: %d\n", balance)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	wg.Wait()
	fmt.Printf("Final balance: %d\n", account.GetBalance())
}
