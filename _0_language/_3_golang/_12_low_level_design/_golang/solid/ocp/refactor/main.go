package main

import "fmt"

type Payment interface {
	Pay(amount float64)
}

type Card struct{}

type UPI struct{}

func (Card) Pay(amount float64) { fmt.Printf("Processing CARD: %.2f\n", amount) }
func (UPI) Pay(amount float64)  { fmt.Printf("Processing UPI: %.2f\n", amount) }

func Checkout(p Payment, amount float64) {
	p.Pay(amount)
}

func main() {
	Checkout(Card{}, 120)
	Checkout(UPI{}, 80)
}
