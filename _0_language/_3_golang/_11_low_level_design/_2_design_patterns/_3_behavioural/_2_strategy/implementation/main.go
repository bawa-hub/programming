package main

import (
	"fmt"
	"strategy/basic"
	"strategy/payment"
)

func main() {

	// 1. BASIC STRATEGY
	fmt.Println("1. BASIC STRATEGY:")
	context := basic.NewContext(&basic.ConcreteStrategyA{})
	
	fmt.Printf("Current strategy: %s\n", context.GetCurrentStrategy())
	fmt.Println(context.ExecuteStrategy("Hello"))
	
	context.SetStrategy(&basic.ConcreteStrategyB{})
	fmt.Printf("Current strategy: %s\n", context.GetCurrentStrategy())
	fmt.Println(context.ExecuteStrategy("World"))
	
	context.SetStrategy(&basic.ConcreteStrategyC{})
	fmt.Printf("Current strategy: %s\n", context.GetCurrentStrategy())
	fmt.Println(context.ExecuteStrategy("Strategy Pattern"))
	fmt.Println()

	// // 2. REAL-WORLD EXAMPLES
	fmt.Println("2. Payment Strategy:")

	// Payment Processing Strategy
	fmt.Println("Payment Processing Strategy:")
	// paymentProcessor := payment.NewPaymentProcessor(&CreditCardStrategy{
	// 	cardNumber: "1234-5678-9012-3456",
	// 	expiryDate: "12/25",
	// 	cvv:        "123",
	// })

	paymentProcessor := payment.NewPaymentProcessor(payment.NewCreditCardStrategy("1234-5678-9012-3456","12/25","123"))
	
	amount := 100.0
	currency := "USD"
	
	transactionID, err := paymentProcessor.ProcessPayment(amount, currency)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
	} else {
		fmt.Printf("Payment successful: %s\n", transactionID)
		fmt.Printf("Fee: $%.2f\n", paymentProcessor.GetFee(amount))
	}
	
	// Switch to PayPal
	// paymentProcessor.SetStrategy(&PayPalStrategy{
	// 	email:    "user@example.com",
	// 	password: "password123",
	// })

	paymentProcessor.SetStrategy(payment.NewPayPalStrategy("user@example.com","password123"))
	
	transactionID, err = paymentProcessor.ProcessPayment(amount, currency)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
	} else {
		fmt.Printf("Payment successful: %s\n", transactionID)
		fmt.Printf("Fee: $%.2f\n", paymentProcessor.GetFee(amount))
	}
	fmt.Println()

	// // Sorting Strategy
	// fmt.Println("Sorting Strategy:")
	// data := []int{64, 34, 25, 12, 22, 11, 90}
	
	// sortingContext := NewSortingContext(&BubbleSortStrategy{})
	// sortingContext.Sort(data)
	
	// sortingContext.SetStrategy(&QuickSortStrategy{})
	// sortingContext.Sort(data)
	
	// sortingContext.SetStrategy(&MergeSortStrategy{})
	// sortingContext.Sort(data)
	// fmt.Println()

	// // Validation Strategy
	// fmt.Println("Validation Strategy:")
	// validationContext := NewValidationContext(&EmailValidationStrategy{})
	
	// testData := []string{
	// 	"user@example.com",
	// 	"invalid-email",
	// 	"1234567890",
	// 	"123",
	// 	"Password123!",
	// 	"weak",
	// }
	
	// strategies := []ValidationStrategy{
	// 	&EmailValidationStrategy{},
	// 	&PhoneValidationStrategy{},
	// 	&PasswordValidationStrategy{},
	// }
	
	// for _, strategy := range strategies {
	// 	validationContext.SetStrategy(strategy)
	// 	fmt.Printf("Testing %s:\n", strategy.GetName())
		
	// 	for _, data := range testData {
	// 		valid, message := validationContext.Validate(data)
	// 		fmt.Printf("  %s: %t - %s\n", data, valid, message)
	// 	}
	// 	fmt.Println()
	// }

	// // Pricing Strategy
	// fmt.Println("Pricing Strategy:")
	// basePrice := 10.0
	// quantities := []int{1, 5, 10, 20}
	
	// pricingContext := NewPricingContext(&RegularPricingStrategy{})
	// fmt.Printf("Regular Pricing:\n")
	// for _, qty := range quantities {
	// 	price := pricingContext.CalculatePrice(basePrice, qty)
	// 	fmt.Printf("  %d items: $%.2f\n", qty, price)
	// }
	
	// pricingContext.SetStrategy(NewBulkDiscountStrategy(10.0, 5))
	// fmt.Printf("\nBulk Discount (10%% for 5+ items):\n")
	// for _, qty := range quantities {
	// 	price := pricingContext.CalculatePrice(basePrice, qty)
	// 	fmt.Printf("  %d items: $%.2f\n", qty, price)
	// }
	
	// pricingContext.SetStrategy(NewPremiumPricingStrategy(1.5))
	// fmt.Printf("\nPremium Pricing (1.5x multiplier):\n")
	// for _, qty := range quantities {
	// 	price := pricingContext.CalculatePrice(basePrice, qty)
	// 	fmt.Printf("  %d items: $%.2f\n", qty, price)
	// }
	// fmt.Println()

	// fmt.Println("=== END OF DEMONSTRATION ===")
}