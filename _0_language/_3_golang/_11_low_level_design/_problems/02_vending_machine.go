package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// CORE ENTITIES
// =============================================================================

// Product Categories
type ProductCategory int

const (
	Snack ProductCategory = iota
	Drink
	Candy
	Gum
)

func (pc ProductCategory) String() string {
	switch pc {
	case Snack:
		return "Snack"
	case Drink:
		return "Drink"
	case Candy:
		return "Candy"
	case Gum:
		return "Gum"
	default:
		return "Unknown"
	}
}

// Product
type Product struct {
	ID           string
	Name         string
	Price        float64
	Quantity     int
	Category     ProductCategory
	ExpirationDate *time.Time
}

func NewProduct(id, name string, price float64, quantity int, category ProductCategory) *Product {
	return &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
		Category: category,
	}
}

func (p *Product) IsAvailable() bool {
	return p.Quantity > 0
}

func (p *Product) IsExpired() bool {
	if p.ExpirationDate == nil {
		return false
	}
	return time.Now().After(*p.ExpirationDate)
}

func (p *Product) DecreaseQuantity(amount int) error {
	if p.Quantity < amount {
		return fmt.Errorf("insufficient quantity: requested %d, available %d", amount, p.Quantity)
	}
	p.Quantity -= amount
	return nil
}

func (p *Product) IncreaseQuantity(amount int) {
	p.Quantity += amount
}

// =============================================================================
// PAYMENT SYSTEM
// =============================================================================

// Payment Methods
type PaymentMethod int

const (
	Cash PaymentMethod = iota
	CreditCard
	DebitCard
	MobilePayment
)

func (pm PaymentMethod) String() string {
	switch pm {
	case Cash:
		return "Cash"
	case CreditCard:
		return "Credit Card"
	case DebitCard:
		return "Debit Card"
	case MobilePayment:
		return "Mobile Payment"
	default:
		return "Unknown"
	}
}

// Payment Status
type PaymentStatus int

const (
	Pending PaymentStatus = iota
	Completed
	Failed
	Refunded
)

func (ps PaymentStatus) String() string {
	switch ps {
	case Pending:
		return "Pending"
	case Completed:
		return "Completed"
	case Failed:
		return "Failed"
	case Refunded:
		return "Refunded"
	default:
		return "Unknown"
	}
}

// Payment
type Payment struct {
	ID        string
	Amount    float64
	Method    PaymentMethod
	Status    PaymentStatus
	Timestamp time.Time
}

func NewPayment(amount float64, method PaymentMethod) *Payment {
	return &Payment{
		ID:        fmt.Sprintf("PAY%d", time.Now().UnixNano()),
		Amount:    amount,
		Method:    method,
		Status:    Pending,
		Timestamp: time.Now(),
	}
}

// =============================================================================
// CASH DRAWER SYSTEM
// =============================================================================

// Coin Types
type CoinType int

const (
	Penny CoinType = iota
	Nickel
	Dime
	Quarter
	Dollar
)

func (ct CoinType) String() string {
	switch ct {
	case Penny:
		return "Penny"
	case Nickel:
		return "Nickel"
	case Dime:
		return "Dime"
	case Quarter:
		return "Quarter"
	case Dollar:
		return "Dollar"
	default:
		return "Unknown"
	}
}

func (ct CoinType) Value() float64 {
	switch ct {
	case Penny:
		return 0.01
	case Nickel:
		return 0.05
	case Dime:
		return 0.10
	case Quarter:
		return 0.25
	case Dollar:
		return 1.00
	default:
		return 0.0
	}
}

// Bill Types
type BillType int

const (
	OneDollar BillType = iota
	FiveDollar
	TenDollar
	TwentyDollar
)

func (bt BillType) String() string {
	switch bt {
	case OneDollar:
		return "One Dollar"
	case FiveDollar:
		return "Five Dollar"
	case TenDollar:
		return "Ten Dollar"
	case TwentyDollar:
		return "Twenty Dollar"
	default:
		return "Unknown"
	}
}

func (bt BillType) Value() float64 {
	switch bt {
	case OneDollar:
		return 1.00
	case FiveDollar:
		return 5.00
	case TenDollar:
		return 10.00
	case TwentyDollar:
		return 20.00
	default:
		return 0.0
	}
}

// Cash Drawer
type CashDrawer struct {
	Coins      map[CoinType]int
	Bills      map[BillType]int
	TotalAmount float64
	mu         sync.RWMutex
}

func NewCashDrawer() *CashDrawer {
	return &CashDrawer{
		Coins:      make(map[CoinType]int),
		Bills:      make(map[BillType]int),
		TotalAmount: 0.0,
	}
}

func (cd *CashDrawer) AddCoin(coinType CoinType, count int) {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	
	cd.Coins[coinType] += count
	cd.TotalAmount += float64(count) * coinType.Value()
}

func (cd *CashDrawer) AddBill(billType BillType, count int) {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	
	cd.Bills[billType] += count
	cd.TotalAmount += float64(count) * billType.Value()
}

func (cd *CashDrawer) GetTotalAmount() float64 {
	cd.mu.RLock()
	defer cd.mu.RUnlock()
	return cd.TotalAmount
}

func (cd *CashDrawer) CalculateChange(amountPaid, productPrice float64) float64 {
	change := amountPaid - productPrice
	if change < 0 {
		return 0
	}
	return change
}

func (cd *CashDrawer) DispenseChange(changeAmount float64) error {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	
	if changeAmount > cd.TotalAmount {
		return fmt.Errorf("insufficient change: requested %.2f, available %.2f", changeAmount, cd.TotalAmount)
	}
	
	// Simple change dispensing logic
	remaining := changeAmount
	
	// Dispense bills first
	billTypes := []BillType{TwentyDollar, TenDollar, FiveDollar, OneDollar}
	for _, billType := range billTypes {
		if remaining <= 0 {
			break
		}
		
		billValue := billType.Value()
		if remaining >= billValue && cd.Bills[billType] > 0 {
			count := int(remaining / billValue)
			if count > cd.Bills[billType] {
				count = cd.Bills[billType]
			}
			
			cd.Bills[billType] -= count
			remaining -= float64(count) * billValue
		}
	}
	
	// Dispense coins
	coinTypes := []CoinType{Dollar, Quarter, Dime, Nickel, Penny}
	for _, coinType := range coinTypes {
		if remaining <= 0 {
			break
		}
		
		coinValue := coinType.Value()
		if remaining >= coinValue && cd.Coins[coinType] > 0 {
			count := int(remaining / coinValue)
			if count > cd.Coins[coinType] {
				count = cd.Coins[coinType]
			}
			
			cd.Coins[coinType] -= count
			remaining -= float64(count) * coinValue
		}
	}
	
	cd.TotalAmount -= changeAmount
	return nil
}

// =============================================================================
// VENDING MACHINE STATES
// =============================================================================

// Vending Machine State Interface
type VendingMachineState interface {
	SelectProduct(productID string) error
	InsertMoney(amount float64) error
	ProcessPayment() error
	DispenseProduct() (*Product, error)
	CancelTransaction() error
	GetStateName() string
}

// Idle State
type IdleState struct {
	vendingMachine *VendingMachine
}

func NewIdleState(vm *VendingMachine) *IdleState {
	return &IdleState{vendingMachine: vm}
}

func (is *IdleState) SelectProduct(productID string) error {
	product, exists := is.vendingMachine.Products[productID]
	if !exists {
		return fmt.Errorf("product %s not found", productID)
	}
	
	if !product.IsAvailable() {
		return fmt.Errorf("product %s is out of stock", productID)
	}
	
	is.vendingMachine.SelectedProduct = product
	is.vendingMachine.SetState(NewProductSelectedState(is.vendingMachine))
	fmt.Printf("Product selected: %s (Price: $%.2f)\n", product.Name, product.Price)
	return nil
}

func (is *IdleState) InsertMoney(amount float64) error {
	return fmt.Errorf("please select a product first")
}

func (is *IdleState) ProcessPayment() error {
	return fmt.Errorf("please select a product first")
}

func (is *IdleState) DispenseProduct() (*Product, error) {
	return nil, fmt.Errorf("please select a product first")
}

func (is *IdleState) CancelTransaction() error {
	return fmt.Errorf("no transaction to cancel")
}

func (is *IdleState) GetStateName() string {
	return "Idle"
}

// Product Selected State
type ProductSelectedState struct {
	vendingMachine *VendingMachine
}

func NewProductSelectedState(vm *VendingMachine) *ProductSelectedState {
	return &ProductSelectedState{vendingMachine: vm}
}

func (pss *ProductSelectedState) SelectProduct(productID string) error {
	return fmt.Errorf("product already selected, please complete current transaction")
}

func (pss *ProductSelectedState) InsertMoney(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %.2f", amount)
	}
	
	pss.vendingMachine.AmountInserted += amount
	pss.vendingMachine.CashDrawer.AddCoin(Quarter, int(amount/0.25)) // Simplified
	
	if pss.vendingMachine.AmountInserted >= pss.vendingMachine.SelectedProduct.Price {
		pss.vendingMachine.SetState(NewPaymentInProgressState(pss.vendingMachine))
		fmt.Printf("Amount inserted: $%.2f, Product price: $%.2f\n", 
			pss.vendingMachine.AmountInserted, pss.vendingMachine.SelectedProduct.Price)
	} else {
		fmt.Printf("Amount inserted: $%.2f, Remaining: $%.2f\n", 
			pss.vendingMachine.AmountInserted, 
			pss.vendingMachine.SelectedProduct.Price - pss.vendingMachine.AmountInserted)
	}
	
	return nil
}

func (pss *ProductSelectedState) ProcessPayment() error {
	return fmt.Errorf("please insert more money")
}

func (pss *ProductSelectedState) DispenseProduct() (*Product, error) {
	return nil, fmt.Errorf("please complete payment first")
}

func (pss *ProductSelectedState) CancelTransaction() error {
	pss.vendingMachine.AmountInserted = 0
	pss.vendingMachine.SelectedProduct = nil
	pss.vendingMachine.SetState(NewIdleState(pss.vendingMachine))
	fmt.Println("Transaction cancelled")
	return nil
}

func (pss *ProductSelectedState) GetStateName() string {
	return "Product Selected"
}

// Payment In Progress State
type PaymentInProgressState struct {
	vendingMachine *VendingMachine
}

func NewPaymentInProgressState(vm *VendingMachine) *PaymentInProgressState {
	return &PaymentInProgressState{vendingMachine: vm}
}

func (pips *PaymentInProgressState) SelectProduct(productID string) error {
	return fmt.Errorf("product already selected, please complete current transaction")
}

func (pips *PaymentInProgressState) InsertMoney(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %.2f", amount)
	}
	
	pips.vendingMachine.AmountInserted += amount
	pips.vendingMachine.CashDrawer.AddCoin(Quarter, int(amount/0.25)) // Simplified
	
	fmt.Printf("Amount inserted: $%.2f, Product price: $%.2f\n", 
		pips.vendingMachine.AmountInserted, pips.vendingMachine.SelectedProduct.Price)
	
	return nil
}

func (pips *PaymentInProgressState) ProcessPayment() error {
	if pips.vendingMachine.AmountInserted < pips.vendingMachine.SelectedProduct.Price {
		return fmt.Errorf("insufficient funds: %.2f < %.2f", 
			pips.vendingMachine.AmountInserted, pips.vendingMachine.SelectedProduct.Price)
	}
	
	// Create payment record
	payment := NewPayment(pips.vendingMachine.AmountInserted, Cash)
	payment.Status = Completed
	pips.vendingMachine.CurrentPayment = payment
	
	// Calculate change
	change := pips.vendingMachine.CashDrawer.CalculateChange(
		pips.vendingMachine.AmountInserted, 
		pips.vendingMachine.SelectedProduct.Price)
	
	if change > 0 {
		fmt.Printf("Change to return: $%.2f\n", change)
		pips.vendingMachine.CashDrawer.DispenseChange(change)
	}
	
	pips.vendingMachine.SetState(NewDispensingState(pips.vendingMachine))
	fmt.Println("Payment processed successfully")
	return nil
}

func (pips *PaymentInProgressState) DispenseProduct() (*Product, error) {
	return nil, fmt.Errorf("please complete payment first")
}

func (pips *PaymentInProgressState) CancelTransaction() error {
	pips.vendingMachine.AmountInserted = 0
	pips.vendingMachine.SelectedProduct = nil
	pips.vendingMachine.SetState(NewIdleState(pips.vendingMachine))
	fmt.Println("Transaction cancelled")
	return nil
}

func (pips *PaymentInProgressState) GetStateName() string {
	return "Payment In Progress"
}

// Dispensing State
type DispensingState struct {
	vendingMachine *VendingMachine
}

func NewDispensingState(vm *VendingMachine) *DispensingState {
	return &DispensingState{vendingMachine: vm}
}

func (ds *DispensingState) SelectProduct(productID string) error {
	return fmt.Errorf("please wait for current transaction to complete")
}

func (ds *DispensingState) InsertMoney(amount float64) error {
	return fmt.Errorf("please wait for current transaction to complete")
}

func (ds *DispensingState) ProcessPayment() error {
	return fmt.Errorf("payment already processed")
}

func (ds *DispensingState) DispenseProduct() (*Product, error) {
	product := ds.vendingMachine.SelectedProduct
	
	// Decrease product quantity
	err := product.DecreaseQuantity(1)
	if err != nil {
		return nil, err
	}
	
	// Create transaction record
	transaction := &Transaction{
		ID:        fmt.Sprintf("TXN%d", time.Now().UnixNano()),
		Product:   product,
		Payment:   ds.vendingMachine.CurrentPayment,
		Timestamp: time.Now(),
		Status:    "Completed",
	}
	
	ds.vendingMachine.Transactions = append(ds.vendingMachine.Transactions, transaction)
	
	// Reset machine state
	ds.vendingMachine.AmountInserted = 0
	ds.vendingMachine.SelectedProduct = nil
	ds.vendingMachine.CurrentPayment = nil
	ds.vendingMachine.SetState(NewIdleState(ds.vendingMachine))
	
	fmt.Printf("Product dispensed: %s\n", product.Name)
	return product, nil
}

func (ds *DispensingState) CancelTransaction() error {
	return fmt.Errorf("cannot cancel transaction in dispensing state")
}

func (ds *DispensingState) GetStateName() string {
	return "Dispensing"
}

// =============================================================================
// TRANSACTION SYSTEM
// =============================================================================

type Transaction struct {
	ID        string
	Product   *Product
	Payment   *Payment
	Timestamp time.Time
	Status    string
}

// =============================================================================
// VENDING MACHINE
// =============================================================================

type VendingMachine struct {
	ID              string
	Products        map[string]*Product
	CashDrawer      *CashDrawer
	CurrentState    VendingMachineState
	SelectedProduct *Product
	AmountInserted  float64
	CurrentPayment  *Payment
	Transactions    []*Transaction
	mu              sync.RWMutex
}

func NewVendingMachine(id string) *VendingMachine {
	vm := &VendingMachine{
		ID:           id,
		Products:     make(map[string]*Product),
		CashDrawer:   NewCashDrawer(),
		Transactions: make([]*Transaction, 0),
	}
	vm.SetState(NewIdleState(vm))
	return vm
}

func (vm *VendingMachine) SetState(state VendingMachineState) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	vm.CurrentState = state
}

func (vm *VendingMachine) GetState() VendingMachineState {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	return vm.CurrentState
}

func (vm *VendingMachine) SelectProduct(productID string) error {
	return vm.GetState().SelectProduct(productID)
}

func (vm *VendingMachine) InsertMoney(amount float64) error {
	return vm.GetState().InsertMoney(amount)
}

func (vm *VendingMachine) ProcessPayment() error {
	return vm.GetState().ProcessPayment()
}

func (vm *VendingMachine) DispenseProduct() (*Product, error) {
	return vm.GetState().DispenseProduct()
}

func (vm *VendingMachine) CancelTransaction() error {
	return vm.GetState().CancelTransaction()
}

func (vm *VendingMachine) AddProduct(product *Product) {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	vm.Products[product.ID] = product
}

func (vm *VendingMachine) GetAvailableProducts() []*Product {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	
	var availableProducts []*Product
	for _, product := range vm.Products {
		if product.IsAvailable() {
			availableProducts = append(availableProducts, product)
		}
	}
	return availableProducts
}

func (vm *VendingMachine) GetProduct(productID string) (*Product, error) {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	
	product, exists := vm.Products[productID]
	if !exists {
		return nil, fmt.Errorf("product %s not found", productID)
	}
	return product, nil
}

func (vm *VendingMachine) RestockProduct(productID string, quantity int) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()
	
	product, exists := vm.Products[productID]
	if !exists {
		return fmt.Errorf("product %s not found", productID)
	}
	
	product.IncreaseQuantity(quantity)
	fmt.Printf("Restocked %s: +%d units (Total: %d)\n", product.Name, quantity, product.Quantity)
	return nil
}

func (vm *VendingMachine) GetMachineStatus() MachineStatus {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	
	totalProducts := len(vm.Products)
	availableProducts := 0
	totalValue := 0.0
	
	for _, product := range vm.Products {
		if product.IsAvailable() {
			availableProducts++
		}
		totalValue += float64(product.Quantity) * product.Price
	}
	
	return MachineStatus{
		ID:                vm.ID,
		State:             vm.GetState().GetStateName(),
		TotalProducts:     totalProducts,
		AvailableProducts: availableProducts,
		TotalValue:        totalValue,
		CashAmount:        vm.CashDrawer.GetTotalAmount(),
	}
}

type MachineStatus struct {
	ID                string
	State             string
	TotalProducts     int
	AvailableProducts int
	TotalValue        float64
	CashAmount        float64
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== VENDING MACHINE SYSTEM DEMONSTRATION ===\n")

	// Create vending machine
	vm := NewVendingMachine("VM001")
	
	// Add products
	products := []*Product{
		NewProduct("P001", "Coca Cola", 1.50, 10, Drink),
		NewProduct("P002", "Pepsi", 1.50, 8, Drink),
		NewProduct("P003", "Chips", 2.00, 15, Snack),
		NewProduct("P004", "Chocolate Bar", 1.75, 12, Candy),
		NewProduct("P005", "Gum", 0.75, 20, Gum),
	}
	
	for _, product := range products {
		vm.AddProduct(product)
	}
	
	// Initialize cash drawer
	vm.CashDrawer.AddCoin(Quarter, 50)
	vm.CashDrawer.AddCoin(Dime, 30)
	vm.CashDrawer.AddCoin(Nickel, 20)
	vm.CashDrawer.AddBill(OneDollar, 10)
	vm.CashDrawer.AddBill(FiveDollar, 5)
	
	// Display initial status
	fmt.Println("1. INITIAL MACHINE STATUS:")
	status := vm.GetMachineStatus()
	fmt.Printf("Machine ID: %s\n", status.ID)
	fmt.Printf("State: %s\n", status.State)
	fmt.Printf("Total Products: %d\n", status.TotalProducts)
	fmt.Printf("Available Products: %d\n", status.AvailableProducts)
	fmt.Printf("Total Value: $%.2f\n", status.TotalValue)
	fmt.Printf("Cash Amount: $%.2f\n", status.CashAmount)
	
	// Display available products
	fmt.Println("\n2. AVAILABLE PRODUCTS:")
	availableProducts := vm.GetAvailableProducts()
	for _, product := range availableProducts {
		fmt.Printf("  %s - %s: $%.2f (Qty: %d)\n", 
			product.ID, product.Name, product.Price, product.Quantity)
	}
	
	fmt.Println()
	
	// Test successful transaction
	fmt.Println("3. SUCCESSFUL TRANSACTION:")
	
	// Select product
	err := vm.SelectProduct("P001")
	if err != nil {
		fmt.Printf("Error selecting product: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	// Insert money
	err = vm.InsertMoney(1.00)
	if err != nil {
		fmt.Printf("Error inserting money: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	// Insert more money
	err = vm.InsertMoney(0.75)
	if err != nil {
		fmt.Printf("Error inserting money: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	// Process payment
	err = vm.ProcessPayment()
	if err != nil {
		fmt.Printf("Error processing payment: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	// Dispense product
	product, err := vm.DispenseProduct()
	if err != nil {
		fmt.Printf("Error dispensing product: %v\n", err)
	} else {
		fmt.Printf("Product dispensed: %s\n", product.Name)
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	fmt.Println()
	
	// Test insufficient funds scenario
	fmt.Println("4. INSUFFICIENT FUNDS SCENARIO:")
	
	// Select product
	err = vm.SelectProduct("P003")
	if err != nil {
		fmt.Printf("Error selecting product: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	// Insert insufficient money
	err = vm.InsertMoney(1.00)
	if err != nil {
		fmt.Printf("Error inserting money: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	// Try to process payment
	err = vm.ProcessPayment()
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
	
	// Cancel transaction
	err = vm.CancelTransaction()
	if err != nil {
		fmt.Printf("Error cancelling transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", vm.GetState().GetStateName())
	}
	
	fmt.Println()
	
	// Test out of stock scenario
	fmt.Println("5. OUT OF STOCK SCENARIO:")
	
	// Restock a product to 0
	err = vm.RestockProduct("P002", -8) // Remove all stock
	if err != nil {
		fmt.Printf("Error restocking: %v\n", err)
	}
	
	// Try to select out of stock product
	err = vm.SelectProduct("P002")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
	
	fmt.Println()
	
	// Test restocking
	fmt.Println("6. RESTOCKING:")
	
	// Restock the product
	err = vm.RestockProduct("P002", 5)
	if err != nil {
		fmt.Printf("Error restocking: %v\n", err)
	}
	
	// Try to select the product again
	err = vm.SelectProduct("P002")
	if err != nil {
		fmt.Printf("Error selecting product: %v\n", err)
	} else {
		fmt.Printf("Product selected successfully after restocking\n")
	}
	
	// Cancel transaction
	vm.CancelTransaction()
	
	fmt.Println()
	
	// Display final status
	fmt.Println("7. FINAL MACHINE STATUS:")
	status = vm.GetMachineStatus()
	fmt.Printf("Machine ID: %s\n", status.ID)
	fmt.Printf("State: %s\n", status.State)
	fmt.Printf("Total Products: %d\n", status.TotalProducts)
	fmt.Printf("Available Products: %d\n", status.AvailableProducts)
	fmt.Printf("Total Value: $%.2f\n", status.TotalValue)
	fmt.Printf("Cash Amount: $%.2f\n", status.CashAmount)
	
	fmt.Println()
	fmt.Println("=== END OF DEMONSTRATION ===")
}
