package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// CORE ENTITIES
// =============================================================================

// Vehicle Types
type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Truck
	Bus
)

func (vt VehicleType) String() string {
	switch vt {
	case Motorcycle:
		return "Motorcycle"
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	case Bus:
		return "Bus"
	default:
		return "Unknown"
	}
}

// Vehicle Interface
type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
	GetSize() int
	GetInfo() VehicleInfo
}

// Vehicle Info
type VehicleInfo struct {
	LicensePlate string
	Type         VehicleType
	Size         int
	EntryTime    time.Time
}

// Concrete Vehicle Types
type MotorcycleVehicle struct {
	licensePlate string
	entryTime    time.Time
}

func NewMotorcycleVehicle(licensePlate string) *MotorcycleVehicle {
	return &MotorcycleVehicle{
		licensePlate: licensePlate,
		entryTime:    time.Now(),
	}
}

func (mv *MotorcycleVehicle) GetLicensePlate() string {
	return mv.licensePlate
}

func (mv *MotorcycleVehicle) GetType() VehicleType {
	return Motorcycle
}

func (mv *MotorcycleVehicle) GetSize() int {
	return 1
}

func (mv *MotorcycleVehicle) GetInfo() VehicleInfo {
	return VehicleInfo{
		LicensePlate: mv.licensePlate,
		Type:         mv.GetType(),
		Size:         mv.GetSize(),
		EntryTime:    mv.entryTime,
	}
}

type CarVehicle struct {
	licensePlate string
	entryTime    time.Time
}

func NewCarVehicle(licensePlate string) *CarVehicle {
	return &CarVehicle{
		licensePlate: licensePlate,
		entryTime:    time.Now(),
	}
}

func (cv *CarVehicle) GetLicensePlate() string {
	return cv.licensePlate
}

func (cv *CarVehicle) GetType() VehicleType {
	return Car
}

func (cv *CarVehicle) GetSize() int {
	return 2
}

func (cv *CarVehicle) GetInfo() VehicleInfo {
	return VehicleInfo{
		LicensePlate: cv.licensePlate,
		Type:         cv.GetType(),
		Size:         cv.GetSize(),
		EntryTime:    cv.entryTime,
	}
}

type TruckVehicle struct {
	licensePlate string
	entryTime    time.Time
}

func NewTruckVehicle(licensePlate string) *TruckVehicle {
	return &TruckVehicle{
		licensePlate: licensePlate,
		entryTime:    time.Now(),
	}
}

func (tv *TruckVehicle) GetLicensePlate() string {
	return tv.licensePlate
}

func (tv *TruckVehicle) GetType() VehicleType {
	return Truck
}

func (tv *TruckVehicle) GetSize() int {
	return 4
}

func (tv *TruckVehicle) GetInfo() VehicleInfo {
	return VehicleInfo{
		LicensePlate: tv.licensePlate,
		Type:         tv.GetType(),
		Size:         tv.GetSize(),
		EntryTime:    tv.entryTime,
	}
}

// =============================================================================
// PARKING SPOT SYSTEM
// =============================================================================

// Spot Types
type SpotType int

const (
	CompactSpot SpotType = iota
	RegularSpot
	LargeSpot
	HandicappedSpot
)

func (st SpotType) String() string {
	switch st {
	case CompactSpot:
		return "Compact"
	case RegularSpot:
		return "Regular"
	case LargeSpot:
		return "Large"
	case HandicappedSpot:
		return "Handicapped"
	default:
		return "Unknown"
	}
}

// Parking Spot Interface
type ParkingSpot interface {
	GetID() string
	GetType() SpotType
	GetSize() int
	IsAvailable() bool
	IsHandicapped() bool
	CanFitVehicle(vehicle Vehicle) bool
	ParkVehicle(vehicle Vehicle) error
	UnparkVehicle() (Vehicle, error)
	GetVehicle() Vehicle
	GetInfo() SpotInfo
}

// Spot Info
type SpotInfo struct {
	ID          string
	Type        SpotType
	Size        int
	IsAvailable bool
	IsHandicapped bool
	Vehicle     Vehicle
}

// Concrete Spot Types
type CompactParkingSpot struct {
	id        string
	vehicle   Vehicle
	available bool
	mu        sync.RWMutex
}

func NewCompactParkingSpot(id string) *CompactParkingSpot {
	return &CompactParkingSpot{
		id:        id,
		vehicle:   nil,
		available: true,
	}
}

func (cps *CompactParkingSpot) GetID() string {
	return cps.id
}

func (cps *CompactParkingSpot) GetType() SpotType {
	return CompactSpot
}

func (cps *CompactParkingSpot) GetSize() int {
	return 1
}

func (cps *CompactParkingSpot) IsAvailable() bool {
	cps.mu.RLock()
	defer cps.mu.RUnlock()
	return cps.available
}

func (cps *CompactParkingSpot) IsHandicapped() bool {
	return false
}

func (cps *CompactParkingSpot) CanFitVehicle(vehicle Vehicle) bool {
	return vehicle.GetSize() <= cps.GetSize()
}

func (cps *CompactParkingSpot) ParkVehicle(vehicle Vehicle) error {
	cps.mu.Lock()
	defer cps.mu.Unlock()
	
	if !cps.available {
		return fmt.Errorf("spot %s is not available", cps.id)
	}
	
	if !cps.CanFitVehicle(vehicle) {
		return fmt.Errorf("vehicle %s cannot fit in spot %s", vehicle.GetLicensePlate(), cps.id)
	}
	
	cps.vehicle = vehicle
	cps.available = false
	return nil
}

func (cps *CompactParkingSpot) UnparkVehicle() (Vehicle, error) {
	cps.mu.Lock()
	defer cps.mu.Unlock()
	
	if cps.available {
		return nil, fmt.Errorf("spot %s is already empty", cps.id)
	}
	
	vehicle := cps.vehicle
	cps.vehicle = nil
	cps.available = true
	return vehicle, nil
}

func (cps *CompactParkingSpot) GetVehicle() Vehicle {
	cps.mu.RLock()
	defer cps.mu.RUnlock()
	return cps.vehicle
}

func (cps *CompactParkingSpot) GetInfo() SpotInfo {
	cps.mu.RLock()
	defer cps.mu.RUnlock()
	return SpotInfo{
		ID:          cps.id,
		Type:        cps.GetType(),
		Size:        cps.GetSize(),
		IsAvailable: cps.available,
		IsHandicapped: cps.IsHandicapped(),
		Vehicle:     cps.vehicle,
	}
}

type RegularParkingSpot struct {
	id        string
	vehicle   Vehicle
	available bool
	mu        sync.RWMutex
}

func NewRegularParkingSpot(id string) *RegularParkingSpot {
	return &RegularParkingSpot{
		id:        id,
		vehicle:   nil,
		available: true,
	}
}

func (rps *RegularParkingSpot) GetID() string {
	return rps.id
}

func (rps *RegularParkingSpot) GetType() SpotType {
	return RegularSpot
}

func (rps *RegularParkingSpot) GetSize() int {
	return 2
}

func (rps *RegularParkingSpot) IsAvailable() bool {
	rps.mu.RLock()
	defer rps.mu.RUnlock()
	return rps.available
}

func (rps *RegularParkingSpot) IsHandicapped() bool {
	return false
}

func (rps *RegularParkingSpot) CanFitVehicle(vehicle Vehicle) bool {
	return vehicle.GetSize() <= rps.GetSize()
}

func (rps *RegularParkingSpot) ParkVehicle(vehicle Vehicle) error {
	rps.mu.Lock()
	defer rps.mu.Unlock()
	
	if !rps.available {
		return fmt.Errorf("spot %s is not available", rps.id)
	}
	
	if !rps.CanFitVehicle(vehicle) {
		return fmt.Errorf("vehicle %s cannot fit in spot %s", vehicle.GetLicensePlate(), rps.id)
	}
	
	rps.vehicle = vehicle
	rps.available = false
	return nil
}

func (rps *RegularParkingSpot) UnparkVehicle() (Vehicle, error) {
	rps.mu.Lock()
	defer rps.mu.Unlock()
	
	if rps.available {
		return nil, fmt.Errorf("spot %s is already empty", rps.id)
	}
	
	vehicle := rps.vehicle
	rps.vehicle = nil
	rps.available = true
	return vehicle, nil
}

func (rps *RegularParkingSpot) GetVehicle() Vehicle {
	rps.mu.RLock()
	defer rps.mu.RUnlock()
	return rps.vehicle
}

func (rps *RegularParkingSpot) GetInfo() SpotInfo {
	rps.mu.RLock()
	defer rps.mu.RUnlock()
	return SpotInfo{
		ID:          rps.id,
		Type:        rps.GetType(),
		Size:        rps.GetSize(),
		IsAvailable: rps.available,
		IsHandicapped: rps.IsHandicapped(),
		Vehicle:     rps.vehicle,
	}
}

type LargeParkingSpot struct {
	id        string
	vehicle   Vehicle
	available bool
	mu        sync.RWMutex
}

func NewLargeParkingSpot(id string) *LargeParkingSpot {
	return &LargeParkingSpot{
		id:        id,
		vehicle:   nil,
		available: true,
	}
}

func (lps *LargeParkingSpot) GetID() string {
	return lps.id
}

func (lps *LargeParkingSpot) GetType() SpotType {
	return LargeSpot
}

func (lps *LargeParkingSpot) GetSize() int {
	return 4
}

func (lps *LargeParkingSpot) IsAvailable() bool {
	lps.mu.RLock()
	defer lps.mu.RUnlock()
	return lps.available
}

func (lps *LargeParkingSpot) IsHandicapped() bool {
	return false
}

func (lps *LargeParkingSpot) CanFitVehicle(vehicle Vehicle) bool {
	return vehicle.GetSize() <= lps.GetSize()
}

func (lps *LargeParkingSpot) ParkVehicle(vehicle Vehicle) error {
	lps.mu.Lock()
	defer lps.mu.Unlock()
	
	if !lps.available {
		return fmt.Errorf("spot %s is not available", lps.id)
	}
	
	if !lps.CanFitVehicle(vehicle) {
		return fmt.Errorf("vehicle %s cannot fit in spot %s", vehicle.GetLicensePlate(), lps.id)
	}
	
	lps.vehicle = vehicle
	lps.available = false
	return nil
}

func (lps *LargeParkingSpot) UnparkVehicle() (Vehicle, error) {
	lps.mu.Lock()
	defer lps.mu.Unlock()
	
	if lps.available {
		return nil, fmt.Errorf("spot %s is already empty", lps.id)
	}
	
	vehicle := lps.vehicle
	lps.vehicle = nil
	lps.available = true
	return vehicle, nil
}

func (lps *LargeParkingSpot) GetVehicle() Vehicle {
	lps.mu.RLock()
	defer lps.mu.RUnlock()
	return lps.vehicle
}

func (lps *LargeParkingSpot) GetInfo() SpotInfo {
	lps.mu.RLock()
	defer lps.mu.RUnlock()
	return SpotInfo{
		ID:          lps.id,
		Type:        lps.GetType(),
		Size:        lps.GetSize(),
		IsAvailable: lps.available,
		IsHandicapped: lps.IsHandicapped(),
		Vehicle:     lps.vehicle,
	}
}

// =============================================================================
// TICKET SYSTEM
// =============================================================================

type Ticket struct {
	ID        string
	Vehicle   Vehicle
	Spot      ParkingSpot
	EntryTime time.Time
	ExitTime  *time.Time
	Paid      bool
}

func NewTicket(vehicle Vehicle, spot ParkingSpot) *Ticket {
	return &Ticket{
		ID:        fmt.Sprintf("T%d", time.Now().UnixNano()),
		Vehicle:   vehicle,
		Spot:      spot,
		EntryTime: time.Now(),
		ExitTime:  nil,
		Paid:      false,
	}
}

func (t *Ticket) GetDuration() time.Duration {
	if t.ExitTime == nil {
		return time.Since(t.EntryTime)
	}
	return t.ExitTime.Sub(t.EntryTime)
}

func (t *Ticket) MarkExit() {
	now := time.Now()
	t.ExitTime = &now
}

func (t *Ticket) MarkPaid() {
	t.Paid = true
}

// =============================================================================
// PRICING SYSTEM
// =============================================================================

type PricingStrategy interface {
	CalculateFee(ticket *Ticket) float64
}

type HourlyPricingStrategy struct {
	hourlyRate float64
}

func NewHourlyPricingStrategy(hourlyRate float64) *HourlyPricingStrategy {
	return &HourlyPricingStrategy{hourlyRate: hourlyRate}
}

func (hps *HourlyPricingStrategy) CalculateFee(ticket *Ticket) float64 {
	duration := ticket.GetDuration()
	hours := duration.Hours()
	if hours < 1 {
		hours = 1 // Minimum 1 hour
	}
	return hours * hps.hourlyRate
}

type TieredPricingStrategy struct {
	rates map[VehicleType]float64
}

func NewTieredPricingStrategy() *TieredPricingStrategy {
	return &TieredPricingStrategy{
		rates: map[VehicleType]float64{
			Motorcycle: 2.0,
			Car:        5.0,
			Truck:      10.0,
			Bus:        15.0,
		},
	}
}

func (tps *TieredPricingStrategy) CalculateFee(ticket *Ticket) float64 {
	duration := ticket.GetDuration()
	hours := duration.Hours()
	if hours < 1 {
		hours = 1 // Minimum 1 hour
	}
	
	rate, exists := tps.rates[ticket.Vehicle.GetType()]
	if !exists {
		rate = 5.0 // Default rate
	}
	
	return hours * rate
}

// =============================================================================
// PAYMENT SYSTEM
// =============================================================================

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

type Payment struct {
	ID          string
	Amount      float64
	Method      PaymentMethod
	Status      PaymentStatus
	Timestamp   time.Time
	TicketID    string
}

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

type PaymentProcessor struct {
	pricingStrategy PricingStrategy
}

func NewPaymentProcessor(pricingStrategy PricingStrategy) *PaymentProcessor {
	return &PaymentProcessor{
		pricingStrategy: pricingStrategy,
	}
}

func (pp *PaymentProcessor) ProcessPayment(ticket *Ticket, method PaymentMethod) (*Payment, error) {
	amount := pp.pricingStrategy.CalculateFee(ticket)
	
	payment := &Payment{
		ID:        fmt.Sprintf("P%d", time.Now().UnixNano()),
		Amount:    amount,
		Method:    method,
		Status:    Pending,
		Timestamp: time.Now(),
		TicketID:  ticket.ID,
	}
	
	// Simulate payment processing
	time.Sleep(100 * time.Millisecond)
	
	// For demo purposes, all payments succeed
	payment.Status = Completed
	ticket.MarkPaid()
	
	return payment, nil
}

// =============================================================================
// PARKING LOT SYSTEM
// =============================================================================

type ParkingLot struct {
	ID           string
	Name         string
	Spots        []ParkingSpot
	Vehicles     map[string]Vehicle
	Tickets      map[string]*Ticket
	PaymentProcessor *PaymentProcessor
	mu           sync.RWMutex
}

func NewParkingLot(id, name string, spots []ParkingSpot) *ParkingLot {
	return &ParkingLot{
		ID:           id,
		Name:         name,
		Spots:        spots,
		Vehicles:     make(map[string]Vehicle),
		Tickets:      make(map[string]*Ticket),
		PaymentProcessor: NewPaymentProcessor(NewTieredPricingStrategy()),
	}
}

func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) (*Ticket, error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	
	// Check if vehicle is already parked
	if _, exists := pl.Vehicles[vehicle.GetLicensePlate()]; exists {
		return nil, fmt.Errorf("vehicle %s is already parked", vehicle.GetLicensePlate())
	}
	
	// Find available spot
	spot := pl.findAvailableSpot(vehicle)
	if spot == nil {
		return nil, fmt.Errorf("no available spots for vehicle type %s", vehicle.GetType())
	}
	
	// Park vehicle
	err := spot.ParkVehicle(vehicle)
	if err != nil {
		return nil, err
	}
	
	// Create ticket
	ticket := NewTicket(vehicle, spot)
	pl.Tickets[ticket.ID] = ticket
	pl.Vehicles[vehicle.GetLicensePlate()] = vehicle
	
	fmt.Printf("Vehicle %s parked in spot %s\n", vehicle.GetLicensePlate(), spot.GetID())
	return ticket, nil
}

func (pl *ParkingLot) UnparkVehicle(ticketID string) (*Payment, error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	
	ticket, exists := pl.Tickets[ticketID]
	if !exists {
		return nil, fmt.Errorf("ticket %s not found", ticketID)
	}
	
	// Unpark vehicle
	vehicle, err := ticket.Spot.UnparkVehicle()
	if err != nil {
		return nil, err
	}
	
	// Mark exit time
	ticket.MarkExit()
	
	// Process payment
	payment, err := pl.PaymentProcessor.ProcessPayment(ticket, CreditCard)
	if err != nil {
		return nil, err
	}
	
	// Remove from tracking
	delete(pl.Vehicles, vehicle.GetLicensePlate())
	
	fmt.Printf("Vehicle %s unparked from spot %s\n", vehicle.GetLicensePlate(), ticket.Spot.GetID())
	return payment, nil
}

func (pl *ParkingLot) FindAvailableSpots(vehicleType VehicleType) []ParkingSpot {
	pl.mu.RLock()
	defer pl.mu.RUnlock()
	
	var availableSpots []ParkingSpot
	for _, spot := range pl.Spots {
		if spot.IsAvailable() {
			// Create a temporary vehicle to check if it fits
			var tempVehicle Vehicle
			switch vehicleType {
			case Motorcycle:
				tempVehicle = NewMotorcycleVehicle("temp")
			case Car:
				tempVehicle = NewCarVehicle("temp")
			case Truck:
				tempVehicle = NewTruckVehicle("temp")
			default:
				continue
			}
			
			if spot.CanFitVehicle(tempVehicle) {
				availableSpots = append(availableSpots, spot)
			}
		}
	}
	
	return availableSpots
}

func (pl *ParkingLot) IsFull() bool {
	pl.mu.RLock()
	defer pl.mu.RUnlock()
	
	for _, spot := range pl.Spots {
		if spot.IsAvailable() {
			return false
		}
	}
	return true
}

func (pl *ParkingLot) GetStatus() ParkingLotStatus {
	pl.mu.RLock()
	defer pl.mu.RUnlock()
	
	totalSpots := len(pl.Spots)
	occupiedSpots := 0
	availableSpots := 0
	
	for _, spot := range pl.Spots {
		if spot.IsAvailable() {
			availableSpots++
		} else {
			occupiedSpots++
		}
	}
	
	return ParkingLotStatus{
		TotalSpots:     totalSpots,
		OccupiedSpots:  occupiedSpots,
		AvailableSpots: availableSpots,
		IsFull:         availableSpots == 0,
	}
}

func (pl *ParkingLot) findAvailableSpot(vehicle Vehicle) ParkingSpot {
	for _, spot := range pl.Spots {
		if spot.IsAvailable() && spot.CanFitVehicle(vehicle) {
			return spot
		}
	}
	return nil
}

type ParkingLotStatus struct {
	TotalSpots     int
	OccupiedSpots  int
	AvailableSpots int
	IsFull         bool
}

// =============================================================================
// VEHICLE FACTORY
// =============================================================================

type VehicleFactory struct{}

func NewVehicleFactory() *VehicleFactory {
	return &VehicleFactory{}
}

func (vf *VehicleFactory) CreateVehicle(licensePlate string, vehicleType VehicleType) Vehicle {
	switch vehicleType {
	case Motorcycle:
		return NewMotorcycleVehicle(licensePlate)
	case Car:
		return NewCarVehicle(licensePlate)
	case Truck:
		return NewTruckVehicle(licensePlate)
	default:
		return NewCarVehicle(licensePlate)
	}
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== PARKING LOT SYSTEM DEMONSTRATION ===\n")

	// Create parking lot with different types of spots
	spots := []ParkingSpot{
		NewCompactParkingSpot("C1"),
		NewCompactParkingSpot("C2"),
		NewRegularParkingSpot("R1"),
		NewRegularParkingSpot("R2"),
		NewRegularParkingSpot("R3"),
		NewLargeParkingSpot("L1"),
		NewLargeParkingSpot("L2"),
	}
	
	parkingLot := NewParkingLot("PL001", "Downtown Parking", spots)
	vehicleFactory := NewVehicleFactory()
	
	// Test parking different vehicles
	fmt.Println("1. PARKING VEHICLES:")
	
	// Park a motorcycle
	motorcycle := vehicleFactory.CreateVehicle("M123", Motorcycle)
	ticket1, err := parkingLot.ParkVehicle(motorcycle)
	if err != nil {
		fmt.Printf("Error parking motorcycle: %v\n", err)
	} else {
		fmt.Printf("Motorcycle parked. Ticket: %s\n", ticket1.ID)
	}
	
	// Park a car
	car := vehicleFactory.CreateVehicle("C456", Car)
	ticket2, err := parkingLot.ParkVehicle(car)
	if err != nil {
		fmt.Printf("Error parking car: %v\n", err)
	} else {
		fmt.Printf("Car parked. Ticket: %s\n", ticket2.ID)
	}
	
	// Park a truck
	truck := vehicleFactory.CreateVehicle("T789", Truck)
	ticket3, err := parkingLot.ParkVehicle(truck)
	if err != nil {
		fmt.Printf("Error parking truck: %v\n", err)
	} else {
		fmt.Printf("Truck parked. Ticket: %s\n", ticket3.ID)
	}
	
	// Try to park another car
	car2 := vehicleFactory.CreateVehicle("C999", Car)
	ticket4, err := parkingLot.ParkVehicle(car2)
	if err != nil {
		fmt.Printf("Error parking second car: %v\n", err)
	} else {
		fmt.Printf("Second car parked. Ticket: %s\n", ticket4.ID)
	}
	
	fmt.Println()
	
	// Test finding available spots
	fmt.Println("2. FINDING AVAILABLE SPOTS:")
	availableSpots := parkingLot.FindAvailableSpots(Car)
	fmt.Printf("Available spots for cars: %d\n", len(availableSpots))
	for _, spot := range availableSpots {
		fmt.Printf("  Spot %s (%s)\n", spot.GetID(), spot.GetType())
	}
	
	availableSpots = parkingLot.FindAvailableSpots(Truck)
	fmt.Printf("Available spots for trucks: %d\n", len(availableSpots))
	for _, spot := range availableSpots {
		fmt.Printf("  Spot %s (%s)\n", spot.GetID(), spot.GetType())
	}
	
	fmt.Println()
	
	// Test parking lot status
	fmt.Println("3. PARKING LOT STATUS:")
	status := parkingLot.GetStatus()
	fmt.Printf("Total spots: %d\n", status.TotalSpots)
	fmt.Printf("Occupied spots: %d\n", status.OccupiedSpots)
	fmt.Printf("Available spots: %d\n", status.AvailableSpots)
	fmt.Printf("Is full: %t\n", status.IsFull)
	
	fmt.Println()
	
	// Test unparking
	fmt.Println("4. UNPARKING VEHICLES:")
	
	// Unpark motorcycle
	payment1, err := parkingLot.UnparkVehicle(ticket1.ID)
	if err != nil {
		fmt.Printf("Error unparking motorcycle: %v\n", err)
	} else {
		fmt.Printf("Motorcycle unparked. Payment: $%.2f via %s\n", payment1.Amount, payment1.Method)
	}
	
	// Unpark car
	payment2, err := parkingLot.UnparkVehicle(ticket2.ID)
	if err != nil {
		fmt.Printf("Error unparking car: %v\n", err)
	} else {
		fmt.Printf("Car unparked. Payment: $%.2f via %s\n", payment2.Amount, payment2.Method)
	}
	
	fmt.Println()
	
	// Test final status
	fmt.Println("5. FINAL PARKING LOT STATUS:")
	status = parkingLot.GetStatus()
	fmt.Printf("Total spots: %d\n", status.TotalSpots)
	fmt.Printf("Occupied spots: %d\n", status.OccupiedSpots)
	fmt.Printf("Available spots: %d\n", status.AvailableSpots)
	fmt.Printf("Is full: %t\n", status.IsFull)
	
	fmt.Println()
	
	// Test error scenarios
	fmt.Println("6. ERROR SCENARIOS:")
	
	// Try to park vehicle with same license plate
	duplicateCar := vehicleFactory.CreateVehicle("C456", Car)
	_, err = parkingLot.ParkVehicle(duplicateCar)
	if err != nil {
		fmt.Printf("Expected error for duplicate license plate: %v\n", err)
	}
	
	// Try to unpark with invalid ticket
	_, err = parkingLot.UnparkVehicle("INVALID")
	if err != nil {
		fmt.Printf("Expected error for invalid ticket: %v\n", err)
	}
	
	fmt.Println()
	fmt.Println("=== END OF DEMONSTRATION ===")
}
