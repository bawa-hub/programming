package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC STATE PATTERN
// =============================================================================

// State interface
type State interface {
	Handle(context *Context)
	GetName() string
}

// Concrete States
type ConcreteStateA struct{}

func (csa *ConcreteStateA) Handle(context *Context) {
	fmt.Println("Handling request in State A")
	context.SetState(&ConcreteStateB{})
}

func (csa *ConcreteStateA) GetName() string {
	return "State A"
}

type ConcreteStateB struct{}

func (csb *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Handling request in State B")
	context.SetState(&ConcreteStateA{})
}

func (csb *ConcreteStateB) GetName() string {
	return "State B"
}

// Context
type Context struct {
	state State
}

func NewContext(initialState State) *Context {
	return &Context{state: initialState}
}

func (c *Context) SetState(state State) {
	c.state = state
	fmt.Printf("Context: State changed to %s\n", state.GetName())
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func (c *Context) GetCurrentState() string {
	return c.state.GetName()
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. VENDING MACHINE STATE
type VendingMachineState interface {
	InsertMoney(amount float64) error
	SelectProduct(product string) error
	DispenseProduct() error
	RefundMoney() error
	GetName() string
}

type VendingMachine struct {
	state           VendingMachineState
	moneyInserted   float64
	selectedProduct string
	products        map[string]float64
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		moneyInserted:   0,
		selectedProduct: "",
		products: map[string]float64{
			"coke":     1.50,
			"water":    1.00,
			"chips":    2.00,
			"candy":    0.75,
		},
	}
	vm.state = &IdleState{vm: vm}
	return vm
}

func (vm *VendingMachine) SetState(state VendingMachineState) {
	vm.state = state
	fmt.Printf("Vending Machine: State changed to %s\n", state.GetName())
}

func (vm *VendingMachine) InsertMoney(amount float64) error {
	return vm.state.InsertMoney(amount)
}

func (vm *VendingMachine) SelectProduct(product string) error {
	return vm.state.SelectProduct(product)
}

func (vm *VendingMachine) DispenseProduct() error {
	return vm.state.DispenseProduct()
}

func (vm *VendingMachine) RefundMoney() error {
	return vm.state.RefundMoney()
}

func (vm *VendingMachine) GetCurrentState() string {
	return vm.state.GetName()
}

// Idle State
type IdleState struct {
	vm *VendingMachine
}

func (is *IdleState) InsertMoney(amount float64) error {
	is.vm.moneyInserted = amount
	fmt.Printf("Inserted $%.2f\n", amount)
	is.vm.SetState(&HasMoneyState{vm: is.vm})
	return nil
}

func (is *IdleState) SelectProduct(product string) error {
	return fmt.Errorf("please insert money first")
}

func (is *IdleState) DispenseProduct() error {
	return fmt.Errorf("please insert money first")
}

func (is *IdleState) RefundMoney() error {
	return fmt.Errorf("no money to refund")
}

func (is *IdleState) GetName() string {
	return "Idle"
}

// Has Money State
type HasMoneyState struct {
	vm *VendingMachine
}

func (hms *HasMoneyState) InsertMoney(amount float64) error {
	hms.vm.moneyInserted += amount
	fmt.Printf("Inserted additional $%.2f, total: $%.2f\n", amount, hms.vm.moneyInserted)
	return nil
}

func (hms *HasMoneyState) SelectProduct(product string) error {
	if price, exists := hms.vm.products[product]; exists {
		if hms.vm.moneyInserted >= price {
			hms.vm.selectedProduct = product
			fmt.Printf("Selected %s for $%.2f\n", product, price)
			hms.vm.SetState(&ProductSelectedState{vm: hms.vm})
			return nil
		}
		return fmt.Errorf("insufficient funds for %s (need $%.2f, have $%.2f)", product, price, hms.vm.moneyInserted)
	}
	return fmt.Errorf("product %s not available", product)
}

func (hms *HasMoneyState) DispenseProduct() error {
	return fmt.Errorf("please select a product first")
}

func (hms *HasMoneyState) RefundMoney() error {
	fmt.Printf("Refunding $%.2f\n", hms.vm.moneyInserted)
	hms.vm.moneyInserted = 0
	hms.vm.SetState(&IdleState{vm: hms.vm})
	return nil
}

func (hms *HasMoneyState) GetName() string {
	return "Has Money"
}

// Product Selected State
type ProductSelectedState struct {
	vm *VendingMachine
}

func (pss *ProductSelectedState) InsertMoney(amount float64) error {
	pss.vm.moneyInserted += amount
	fmt.Printf("Inserted additional $%.2f, total: $%.2f\n", amount, pss.vm.moneyInserted)
	return nil
}

func (pss *ProductSelectedState) SelectProduct(product string) error {
	return fmt.Errorf("product already selected, please dispense or refund")
}

func (pss *ProductSelectedState) DispenseProduct() error {
	price := pss.vm.products[pss.vm.selectedProduct]
	if pss.vm.moneyInserted >= price {
		change := pss.vm.moneyInserted - price
		fmt.Printf("Dispensing %s\n", pss.vm.selectedProduct)
		if change > 0 {
			fmt.Printf("Returning change: $%.2f\n", change)
		}
		pss.vm.moneyInserted = 0
		pss.vm.selectedProduct = ""
		pss.vm.SetState(&IdleState{vm: pss.vm})
		return nil
	}
	return fmt.Errorf("insufficient funds")
}

func (pss *ProductSelectedState) RefundMoney() error {
	fmt.Printf("Refunding $%.2f\n", pss.vm.moneyInserted)
	pss.vm.moneyInserted = 0
	pss.vm.selectedProduct = ""
	pss.vm.SetState(&IdleState{vm: pss.vm})
	return nil
}

func (pss *ProductSelectedState) GetName() string {
	return "Product Selected"
}

// 2. MEDIA PLAYER STATE
type MediaPlayerState interface {
	Play() error
	Pause() error
	Stop() error
	Next() error
	Previous() error
	GetName() string
}

type MediaPlayer struct {
	state      MediaPlayerState
	currentTrack string
	playlist   []string
	trackIndex int
}

func NewMediaPlayer(playlist []string) *MediaPlayer {
	mp := &MediaPlayer{
		currentTrack: "",
		playlist:     playlist,
		trackIndex:   0,
	}
	mp.state = &StoppedState{mp: mp}
	return mp
}

func (mp *MediaPlayer) SetState(state MediaPlayerState) {
	mp.state = state
	fmt.Printf("Media Player: State changed to %s\n", state.GetName())
}

func (mp *MediaPlayer) Play() error {
	return mp.state.Play()
}

func (mp *MediaPlayer) Pause() error {
	return mp.state.Pause()
}

func (mp *MediaPlayer) Stop() error {
	return mp.state.Stop()
}

func (mp *MediaPlayer) Next() error {
	return mp.state.Next()
}

func (mp *MediaPlayer) Previous() error {
	return mp.state.Previous()
}

func (mp *MediaPlayer) GetCurrentState() string {
	return mp.state.GetName()
}

func (mp *MediaPlayer) GetCurrentTrack() string {
	return mp.currentTrack
}

// Stopped State
type StoppedState struct {
	mp *MediaPlayer
}

func (ss *StoppedState) Play() error {
	if len(ss.mp.playlist) == 0 {
		return fmt.Errorf("no tracks in playlist")
	}
	ss.mp.currentTrack = ss.mp.playlist[ss.mp.trackIndex]
	fmt.Printf("Playing: %s\n", ss.mp.currentTrack)
	ss.mp.SetState(&PlayingState{mp: ss.mp})
	return nil
}

func (ss *StoppedState) Pause() error {
	return fmt.Errorf("cannot pause when stopped")
}

func (ss *StoppedState) Stop() error {
	return fmt.Errorf("already stopped")
}

func (ss *StoppedState) Next() error {
	if ss.mp.trackIndex < len(ss.mp.playlist)-1 {
		ss.mp.trackIndex++
		ss.mp.currentTrack = ss.mp.playlist[ss.mp.trackIndex]
		fmt.Printf("Next track: %s\n", ss.mp.currentTrack)
	} else {
		fmt.Println("Already at last track")
	}
	return nil
}

func (ss *StoppedState) Previous() error {
	if ss.mp.trackIndex > 0 {
		ss.mp.trackIndex--
		ss.mp.currentTrack = ss.mp.playlist[ss.mp.trackIndex]
		fmt.Printf("Previous track: %s\n", ss.mp.currentTrack)
	} else {
		fmt.Println("Already at first track")
	}
	return nil
}

func (ss *StoppedState) GetName() string {
	return "Stopped"
}

// Playing State
type PlayingState struct {
	mp *MediaPlayer
}

func (ps *PlayingState) Play() error {
	return fmt.Errorf("already playing")
}

func (ps *PlayingState) Pause() error {
	fmt.Printf("Pausing: %s\n", ps.mp.currentTrack)
	ps.mp.SetState(&PausedState{mp: ps.mp})
	return nil
}

func (ps *PlayingState) Stop() error {
	fmt.Printf("Stopping: %s\n", ps.mp.currentTrack)
	ps.mp.currentTrack = ""
	ps.mp.SetState(&StoppedState{mp: ps.mp})
	return nil
}

func (ps *PlayingState) Next() error {
	if ps.mp.trackIndex < len(ps.mp.playlist)-1 {
		ps.mp.trackIndex++
		ps.mp.currentTrack = ps.mp.playlist[ps.mp.trackIndex]
		fmt.Printf("Next track: %s\n", ps.mp.currentTrack)
	} else {
		fmt.Println("Already at last track")
	}
	return nil
}

func (ps *PlayingState) Previous() error {
	if ps.mp.trackIndex > 0 {
		ps.mp.trackIndex--
		ps.mp.currentTrack = ps.mp.playlist[ps.mp.trackIndex]
		fmt.Printf("Previous track: %s\n", ps.mp.currentTrack)
	} else {
		fmt.Println("Already at first track")
	}
	return nil
}

func (ps *PlayingState) GetName() string {
	return "Playing"
}

// Paused State
type PausedState struct {
	mp *MediaPlayer
}

func (ps *PausedState) Play() error {
	fmt.Printf("Resuming: %s\n", ps.mp.currentTrack)
	ps.mp.SetState(&PlayingState{mp: ps.mp})
	return nil
}

func (ps *PausedState) Pause() error {
	return fmt.Errorf("already paused")
}

func (ps *PausedState) Stop() error {
	fmt.Printf("Stopping: %s\n", ps.mp.currentTrack)
	ps.mp.currentTrack = ""
	ps.mp.SetState(&StoppedState{mp: ps.mp})
	return nil
}

func (ps *PausedState) Next() error {
	if ps.mp.trackIndex < len(ps.mp.playlist)-1 {
		ps.mp.trackIndex++
		ps.mp.currentTrack = ps.mp.playlist[ps.mp.trackIndex]
		fmt.Printf("Next track: %s\n", ps.mp.currentTrack)
	} else {
		fmt.Println("Already at last track")
	}
	return nil
}

func (ps *PausedState) Previous() error {
	if ps.mp.trackIndex > 0 {
		ps.mp.trackIndex--
		ps.mp.currentTrack = ps.mp.playlist[ps.mp.trackIndex]
		fmt.Printf("Previous track: %s\n", ps.mp.currentTrack)
	} else {
		fmt.Println("Already at first track")
	}
	return nil
}

func (ps *PausedState) GetName() string {
	return "Paused"
}

// 3. ORDER PROCESSING STATE
type OrderState interface {
	Process() error
	Ship() error
	Deliver() error
	Cancel() error
	GetName() string
}

type Order struct {
	state       OrderState
	orderID     string
	customer    string
	items       []string
	totalAmount float64
}

func NewOrder(orderID, customer string, items []string, totalAmount float64) *Order {
	order := &Order{
		orderID:     orderID,
		customer:    customer,
		items:       items,
		totalAmount: totalAmount,
	}
	order.state = &PendingState{order: order}
	return order
}

func (o *Order) SetState(state OrderState) {
	o.state = state
	fmt.Printf("Order %s: State changed to %s\n", o.orderID, state.GetName())
}

func (o *Order) Process() error {
	return o.state.Process()
}

func (o *Order) Ship() error {
	return o.state.Ship()
}

func (o *Order) Deliver() error {
	return o.state.Deliver()
}

func (o *Order) Cancel() error {
	return o.state.Cancel()
}

func (o *Order) GetCurrentState() string {
	return o.state.GetName()
}

// Pending State
type PendingState struct {
	order *Order
}

func (ps *PendingState) Process() error {
	fmt.Printf("Processing order %s for %s\n", ps.order.orderID, ps.order.customer)
	ps.order.SetState(&ConfirmedState{order: ps.order})
	return nil
}

func (ps *PendingState) Ship() error {
	return fmt.Errorf("cannot ship unprocessed order")
}

func (ps *PendingState) Deliver() error {
	return fmt.Errorf("cannot deliver unprocessed order")
}

func (ps *PendingState) Cancel() error {
	fmt.Printf("Cancelling order %s\n", ps.order.orderID)
	ps.order.SetState(&CancelledState{order: ps.order})
	return nil
}

func (ps *PendingState) GetName() string {
	return "Pending"
}

// Confirmed State
type ConfirmedState struct {
	order *Order
}

func (cs *ConfirmedState) Process() error {
	return fmt.Errorf("order already processed")
}

func (cs *ConfirmedState) Ship() error {
	fmt.Printf("Shipping order %s to %s\n", cs.order.orderID, cs.order.customer)
	cs.order.SetState(&ShippedState{order: cs.order})
	return nil
}

func (cs *ConfirmedState) Deliver() error {
	return fmt.Errorf("cannot deliver unshipped order")
}

func (cs *ConfirmedState) Cancel() error {
	fmt.Printf("Cancelling confirmed order %s\n", cs.order.orderID)
	cs.order.SetState(&CancelledState{order: cs.order})
	return nil
}

func (cs *ConfirmedState) GetName() string {
	return "Confirmed"
}

// Shipped State
type ShippedState struct {
	order *Order
}

func (ss *ShippedState) Process() error {
	return fmt.Errorf("order already processed")
}

func (ss *ShippedState) Ship() error {
	return fmt.Errorf("order already shipped")
}

func (ss *ShippedState) Deliver() error {
	fmt.Printf("Delivering order %s to %s\n", ss.order.orderID, ss.order.customer)
	ss.order.SetState(&DeliveredState{order: ss.order})
	return nil
}

func (ss *ShippedState) Cancel() error {
	return fmt.Errorf("cannot cancel shipped order")
}

func (ss *ShippedState) GetName() string {
	return "Shipped"
}

// Delivered State
type DeliveredState struct {
	order *Order
}

func (ds *DeliveredState) Process() error {
	return fmt.Errorf("order already delivered")
}

func (ds *DeliveredState) Ship() error {
	return fmt.Errorf("order already delivered")
}

func (ds *DeliveredState) Deliver() error {
	return fmt.Errorf("order already delivered")
}

func (ds *DeliveredState) Cancel() error {
	return fmt.Errorf("cannot cancel delivered order")
}

func (ds *DeliveredState) GetName() string {
	return "Delivered"
}

// Cancelled State
type CancelledState struct {
	order *Order
}

func (cs *CancelledState) Process() error {
	return fmt.Errorf("cannot process cancelled order")
}

func (cs *CancelledState) Ship() error {
	return fmt.Errorf("cannot ship cancelled order")
}

func (cs *CancelledState) Deliver() error {
	return fmt.Errorf("cannot deliver cancelled order")
}

func (cs *CancelledState) Cancel() error {
	return fmt.Errorf("order already cancelled")
}

func (cs *CancelledState) GetName() string {
	return "Cancelled"
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== STATE PATTERN DEMONSTRATION ===\n")

	// 1. BASIC STATE
	fmt.Println("1. BASIC STATE:")
	context := NewContext(&ConcreteStateA{})
	
	for i := 0; i < 5; i++ {
		fmt.Printf("Request %d: ", i+1)
		context.Request()
		fmt.Printf("Current state: %s\n", context.GetCurrentState())
	}
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Vending Machine State
	fmt.Println("Vending Machine State:")
	vm := NewVendingMachine()
	
	// Try to select product without money
	vm.SelectProduct("coke")
	
	// Insert money
	vm.InsertMoney(2.0)
	vm.SelectProduct("coke")
	vm.DispenseProduct()
	
	// Try to select another product
	vm.SelectProduct("water")
	vm.InsertMoney(1.0)
	vm.SelectProduct("water")
	vm.DispenseProduct()
	
	// Refund remaining money
	vm.InsertMoney(0.5)
	vm.RefundMoney()
	fmt.Println()

	// Media Player State
	fmt.Println("Media Player State:")
	playlist := []string{"Song 1", "Song 2", "Song 3", "Song 4"}
	mp := NewMediaPlayer(playlist)
	
	// Play music
	mp.Play()
	mp.Pause()
	mp.Play()
	mp.Next()
	mp.Previous()
	mp.Stop()
	
	// Try to pause when stopped
	mp.Pause()
	fmt.Println()

	// Order Processing State
	fmt.Println("Order Processing State:")
	order := NewOrder("ORD-001", "John Doe", []string{"Item 1", "Item 2"}, 100.0)
	
	// Process order
	order.Process()
	order.Ship()
	order.Deliver()
	
	// Try to cancel delivered order
	order.Cancel()
	
	// Create another order and cancel it
	order2 := NewOrder("ORD-002", "Jane Smith", []string{"Item 3"}, 50.0)
	order2.Process()
	order2.Cancel()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
