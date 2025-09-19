package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// =============================================================================
// BASIC STRATEGY PATTERN
// =============================================================================

// Strategy interface
type Strategy interface {
	Execute(data interface{}) interface{}
	GetName() string
}

// Concrete Strategies
type ConcreteStrategyA struct{}

func (csa *ConcreteStrategyA) Execute(data interface{}) interface{} {
	return fmt.Sprintf("Strategy A processed: %v", data)
}

func (csa *ConcreteStrategyA) GetName() string {
	return "Strategy A"
}

type ConcreteStrategyB struct{}

func (csb *ConcreteStrategyB) Execute(data interface{}) interface{} {
	return fmt.Sprintf("Strategy B processed: %v", data)
}

func (csb *ConcreteStrategyB) GetName() string {
	return "Strategy B"
}

type ConcreteStrategyC struct{}

func (csc *ConcreteStrategyC) Execute(data interface{}) interface{} {
	return fmt.Sprintf("Strategy C processed: %v", data)
}

func (csc *ConcreteStrategyC) GetName() string {
	return "Strategy C"
}

// Context
type Context struct {
	strategy Strategy
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(data interface{}) interface{} {
	return c.strategy.Execute(data)
}

func (c *Context) GetCurrentStrategy() string {
	return c.strategy.GetName()
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. PAYMENT PROCESSING STRATEGY
type PaymentStrategy interface {
	ProcessPayment(amount float64, currency string) (string, error)
	GetName() string
	GetFee() float64
}

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

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy: strategy}
}

func (pp *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	pp.strategy = strategy
}

func (pp *PaymentProcessor) ProcessPayment(amount float64, currency string) (string, error) {
	return pp.strategy.ProcessPayment(amount, currency)
}

func (pp *PaymentProcessor) GetFee(amount float64) float64 {
	return amount * pp.strategy.GetFee()
}

func (pp *PaymentProcessor) GetCurrentStrategy() string {
	return pp.strategy.GetName()
}

// 2. SORTING STRATEGY
type SortingStrategy interface {
	Sort(data []int) []int
	GetName() string
	GetTimeComplexity() string
}

type BubbleSortStrategy struct{}

func (bss *BubbleSortStrategy) Sort(data []int) []int {
	fmt.Printf("Bubble sorting: %v\n", data)
	
	// Create a copy to avoid modifying original
	sorted := make([]int, len(data))
	copy(sorted, data)
	
	n := len(sorted)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	
	return sorted
}

func (bss *BubbleSortStrategy) GetName() string {
	return "Bubble Sort"
}

func (bss *BubbleSortStrategy) GetTimeComplexity() string {
	return "O(n²)"
}

type QuickSortStrategy struct{}

func (qss *QuickSortStrategy) Sort(data []int) []int {
	fmt.Printf("Quick sorting: %v\n", data)
	
	// Create a copy to avoid modifying original
	sorted := make([]int, len(data))
	copy(sorted, data)
	
	qss.quickSort(sorted, 0, len(sorted)-1)
	return sorted
}

func (qss *QuickSortStrategy) quickSort(arr []int, low, high int) {
	if low < high {
		pi := qss.partition(arr, low, high)
		qss.quickSort(arr, low, pi-1)
		qss.quickSort(arr, pi+1, high)
	}
}

func (qss *QuickSortStrategy) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (qss *QuickSortStrategy) GetName() string {
	return "Quick Sort"
}

func (qss *QuickSortStrategy) GetTimeComplexity() string {
	return "O(n log n)"
}

type MergeSortStrategy struct{}

func (mss *MergeSortStrategy) Sort(data []int) []int {
	fmt.Printf("Merge sorting: %v\n", data)
	
	// Create a copy to avoid modifying original
	sorted := make([]int, len(data))
	copy(sorted, data)
	
	mss.mergeSort(sorted, 0, len(sorted)-1)
	return sorted
}

func (mss *MergeSortStrategy) mergeSort(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mss.mergeSort(arr, left, mid)
		mss.mergeSort(arr, mid+1, right)
		mss.merge(arr, left, mid, right)
	}
}

func (mss *MergeSortStrategy) merge(arr []int, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid
	
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)
	
	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}
	
	i, j, k := 0, 0, left
	
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func (mss *MergeSortStrategy) GetName() string {
	return "Merge Sort"
}

func (mss *MergeSortStrategy) GetTimeComplexity() string {
	return "O(n log n)"
}

type SortingContext struct {
	strategy SortingStrategy
}

func NewSortingContext(strategy SortingStrategy) *SortingContext {
	return &SortingContext{strategy: strategy}
}

func (sc *SortingContext) SetStrategy(strategy SortingStrategy) {
	sc.strategy = strategy
}

func (sc *SortingContext) Sort(data []int) []int {
	start := time.Now()
	result := sc.strategy.Sort(data)
	duration := time.Since(start)
	
	fmt.Printf("Sorted with %s in %v: %v\n", 
		sc.strategy.GetName(), duration, result)
	return result
}

func (sc *SortingContext) GetCurrentStrategy() string {
	return sc.strategy.GetName()
}

// 3. VALIDATION STRATEGY
type ValidationStrategy interface {
	Validate(data string) (bool, string)
	GetName() string
}

type EmailValidationStrategy struct{}

func (evs *EmailValidationStrategy) Validate(data string) (bool, string) {
	if strings.Contains(data, "@") && strings.Contains(data, ".") {
		return true, "Valid email format"
	}
	return false, "Invalid email format"
}

func (evs *EmailValidationStrategy) GetName() string {
	return "Email Validation"
}

type PhoneValidationStrategy struct{}

func (pvs *PhoneValidationStrategy) Validate(data string) (bool, string) {
	// Simple phone validation (10 digits)
	if len(data) == 10 {
		if _, err := strconv.Atoi(data); err == nil {
			return true, "Valid phone number"
		}
	}
	return false, "Invalid phone number format"
}

func (pvs *PhoneValidationStrategy) GetName() string {
	return "Phone Validation"
}

type PasswordValidationStrategy struct{}

func (pvs *PasswordValidationStrategy) Validate(data string) (bool, string) {
	if len(data) < 8 {
		return false, "Password must be at least 8 characters long"
	}
	
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	
	for _, char := range data {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case char >= 33 && char <= 126:
			hasSpecial = true
		}
	}
	
	if !hasUpper {
		return false, "Password must contain at least one uppercase letter"
	}
	if !hasLower {
		return false, "Password must contain at least one lowercase letter"
	}
	if !hasDigit {
		return false, "Password must contain at least one digit"
	}
	if !hasSpecial {
		return false, "Password must contain at least one special character"
	}
	
	return true, "Valid password"
}

func (pvs *PasswordValidationStrategy) GetName() string {
	return "Password Validation"
}

type ValidationContext struct {
	strategy ValidationStrategy
}

func NewValidationContext(strategy ValidationStrategy) *ValidationContext {
	return &ValidationContext{strategy: strategy}
}

func (vc *ValidationContext) SetStrategy(strategy ValidationStrategy) {
	vc.strategy = strategy
}

func (vc *ValidationContext) Validate(data string) (bool, string) {
	return vc.strategy.Validate(data)
}

func (vc *ValidationContext) GetCurrentStrategy() string {
	return vc.strategy.GetName()
}

// 4. PRICING STRATEGY
type PricingStrategy interface {
	CalculatePrice(basePrice float64, quantity int) float64
	GetName() string
	GetDescription() string
}

type RegularPricingStrategy struct{}

func (rps *RegularPricingStrategy) CalculatePrice(basePrice float64, quantity int) float64 {
	return basePrice * float64(quantity)
}

func (rps *RegularPricingStrategy) GetName() string {
	return "Regular Pricing"
}

func (rps *RegularPricingStrategy) GetDescription() string {
	return "Standard pricing with no discounts"
}

type BulkDiscountStrategy struct {
	discountPercent float64
	minQuantity     int
}

func NewBulkDiscountStrategy(discountPercent float64, minQuantity int) *BulkDiscountStrategy {
	return &BulkDiscountStrategy{
		discountPercent: discountPercent,
		minQuantity:     minQuantity,
	}
}

func (bds *BulkDiscountStrategy) CalculatePrice(basePrice float64, quantity int) float64 {
	total := basePrice * float64(quantity)
	if quantity >= bds.minQuantity {
		discount := total * (bds.discountPercent / 100)
		total -= discount
	}
	return total
}

func (bds *BulkDiscountStrategy) GetName() string {
	return "Bulk Discount"
}

func (bds *BulkDiscountStrategy) GetDescription() string {
	return fmt.Sprintf("%.1f%% discount for %d+ items", bds.discountPercent, bds.minQuantity)
}

type PremiumPricingStrategy struct {
	premiumMultiplier float64
}

func NewPremiumPricingStrategy(premiumMultiplier float64) *PremiumPricingStrategy {
	return &PremiumPricingStrategy{
		premiumMultiplier: premiumMultiplier,
	}
}

func (pps *PremiumPricingStrategy) CalculatePrice(basePrice float64, quantity int) float64 {
	return basePrice * float64(quantity) * pps.premiumMultiplier
}

func (pps *PremiumPricingStrategy) GetName() string {
	return "Premium Pricing"
}

func (pps *PremiumPricingStrategy) GetDescription() string {
	return fmt.Sprintf("Premium pricing with %.1fx multiplier", pps.premiumMultiplier)
}

type PricingContext struct {
	strategy PricingStrategy
}

func NewPricingContext(strategy PricingStrategy) *PricingContext {
	return &PricingContext{strategy: strategy}
}

func (pc *PricingContext) SetStrategy(strategy PricingStrategy) {
	pc.strategy = strategy
}

func (pc *PricingContext) CalculatePrice(basePrice float64, quantity int) float64 {
	return pc.strategy.CalculatePrice(basePrice, quantity)
}

func (pc *PricingContext) GetCurrentStrategy() string {
	return pc.strategy.GetName()
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== STRATEGY PATTERN DEMONSTRATION ===\n")

	// 1. BASIC STRATEGY
	fmt.Println("1. BASIC STRATEGY:")
	context := NewContext(&ConcreteStrategyA{})
	
	fmt.Printf("Current strategy: %s\n", context.GetCurrentStrategy())
	fmt.Println(context.ExecuteStrategy("Hello"))
	
	context.SetStrategy(&ConcreteStrategyB{})
	fmt.Printf("Current strategy: %s\n", context.GetCurrentStrategy())
	fmt.Println(context.ExecuteStrategy("World"))
	
	context.SetStrategy(&ConcreteStrategyC{})
	fmt.Printf("Current strategy: %s\n", context.GetCurrentStrategy())
	fmt.Println(context.ExecuteStrategy("Strategy Pattern"))
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Payment Processing Strategy
	fmt.Println("Payment Processing Strategy:")
	paymentProcessor := NewPaymentProcessor(&CreditCardStrategy{
		cardNumber: "1234-5678-9012-3456",
		expiryDate: "12/25",
		cvv:        "123",
	})
	
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
	paymentProcessor.SetStrategy(&PayPalStrategy{
		email:    "user@example.com",
		password: "password123",
	})
	
	transactionID, err = paymentProcessor.ProcessPayment(amount, currency)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
	} else {
		fmt.Printf("Payment successful: %s\n", transactionID)
		fmt.Printf("Fee: $%.2f\n", paymentProcessor.GetFee(amount))
	}
	fmt.Println()

	// Sorting Strategy
	fmt.Println("Sorting Strategy:")
	data := []int{64, 34, 25, 12, 22, 11, 90}
	
	sortingContext := NewSortingContext(&BubbleSortStrategy{})
	sortingContext.Sort(data)
	
	sortingContext.SetStrategy(&QuickSortStrategy{})
	sortingContext.Sort(data)
	
	sortingContext.SetStrategy(&MergeSortStrategy{})
	sortingContext.Sort(data)
	fmt.Println()

	// Validation Strategy
	fmt.Println("Validation Strategy:")
	validationContext := NewValidationContext(&EmailValidationStrategy{})
	
	testData := []string{
		"user@example.com",
		"invalid-email",
		"1234567890",
		"123",
		"Password123!",
		"weak",
	}
	
	strategies := []ValidationStrategy{
		&EmailValidationStrategy{},
		&PhoneValidationStrategy{},
		&PasswordValidationStrategy{},
	}
	
	for _, strategy := range strategies {
		validationContext.SetStrategy(strategy)
		fmt.Printf("Testing %s:\n", strategy.GetName())
		
		for _, data := range testData {
			valid, message := validationContext.Validate(data)
			fmt.Printf("  %s: %t - %s\n", data, valid, message)
		}
		fmt.Println()
	}

	// Pricing Strategy
	fmt.Println("Pricing Strategy:")
	basePrice := 10.0
	quantities := []int{1, 5, 10, 20}
	
	pricingContext := NewPricingContext(&RegularPricingStrategy{})
	fmt.Printf("Regular Pricing:\n")
	for _, qty := range quantities {
		price := pricingContext.CalculatePrice(basePrice, qty)
		fmt.Printf("  %d items: $%.2f\n", qty, price)
	}
	
	pricingContext.SetStrategy(NewBulkDiscountStrategy(10.0, 5))
	fmt.Printf("\nBulk Discount (10%% for 5+ items):\n")
	for _, qty := range quantities {
		price := pricingContext.CalculatePrice(basePrice, qty)
		fmt.Printf("  %d items: $%.2f\n", qty, price)
	}
	
	pricingContext.SetStrategy(NewPremiumPricingStrategy(1.5))
	fmt.Printf("\nPremium Pricing (1.5x multiplier):\n")
	for _, qty := range quantities {
		price := pricingContext.CalculatePrice(basePrice, qty)
		fmt.Printf("  %d items: $%.2f\n", qty, price)
	}
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
