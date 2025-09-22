package main

import (
	"errors"
	"fmt"
)

type account struct {
	owner string
	balance int
}

func NewAccount(owner string, initial int) (*account, error) {
	if initial < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}
	return &account{owner: owner, balance: initial}, nil
}

func (a *account) Owner() string { return a.owner }
func (a *account) Balance() int  { return a.balance }

func (a *account) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("deposit must be positive")
	}
	a.balance += amount
	return nil
}

func (a *account) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("withdraw must be positive")
	}
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func main() {
	acc, _ := NewAccount("Alice", 100)
	_ = acc.Deposit(50)
	_ = acc.Withdraw(30)
	fmt.Println(acc.Owner(), acc.Balance())
}
