package main

import "fmt"

type Payment struct {
	Amount float64
	Method string // "card", "upi", "netbanking"
}

// OCP violation: every new method requires modifying this function.
func ProcessPayment(p Payment) {
	switch p.Method {
	case "card":
		fmt.Printf("Processing CARD: %.2f\n", p.Amount)
	case "upi":
		fmt.Printf("Processing UPI: %.2f\n", p.Amount)
	case "netbanking":
		fmt.Printf("Processing NETBANKING: %.2f\n", p.Amount)
	default:
		fmt.Println("Unsupported payment method")
	}
}

func main() {
	ProcessPayment(Payment{Amount: 100, Method: "card"})
	ProcessPayment(Payment{Amount: 50, Method: "upi"})
}
