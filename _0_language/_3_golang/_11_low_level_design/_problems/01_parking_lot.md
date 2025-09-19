# Parking Lot System Design

## Problem Statement
Design a parking lot system that can handle different types of vehicles (cars, motorcycles, trucks) with different parking spot sizes and requirements. The system should support parking, unparking, and provide information about available spots.

## Requirements Analysis

### Functional Requirements
1. **Parking Operations**
   - Park a vehicle in an available spot
   - Unpark a vehicle from a spot
   - Find available spots for a specific vehicle type
   - Check if parking lot is full

2. **Vehicle Management**
   - Support different vehicle types (Car, Motorcycle, Truck)
   - Track vehicle information (license plate, type, entry time)
   - Handle vehicle identification and retrieval

3. **Spot Management**
   - Different spot sizes for different vehicle types
   - Track spot availability and occupancy
   - Support spot reservations

4. **Payment System**
   - Calculate parking fees based on duration
   - Support different payment methods
   - Handle payment processing

### Non-Functional Requirements
1. **Scalability**: Support multiple parking lots
2. **Concurrency**: Handle multiple vehicles parking simultaneously
3. **Reliability**: Ensure data consistency
4. **Performance**: Fast spot allocation and deallocation

## Core Entities

### 1. Vehicle
- **Attributes**: License plate, type, size, entry time
- **Types**: Car, Motorcycle, Truck, Bus
- **Behavior**: Get vehicle info, calculate size

### 2. Parking Spot
- **Attributes**: ID, size, location, availability status
- **Types**: Compact, Regular, Large, Handicapped
- **Behavior**: Check availability, reserve, release

### 3. Parking Lot
- **Attributes**: ID, name, location, total spots, available spots
- **Behavior**: Park vehicle, unpark vehicle, find available spots

### 4. Ticket
- **Attributes**: ID, vehicle info, spot info, entry time, exit time
- **Behavior**: Calculate duration, calculate fee

### 5. Payment
- **Attributes**: Amount, method, status, timestamp
- **Behavior**: Process payment, validate payment

## Design Patterns Used

### 1. Factory Pattern
- Create different types of vehicles
- Create different types of parking spots
- Create different types of tickets

### 2. Strategy Pattern
- Different pricing strategies
- Different payment methods
- Different spot allocation strategies

### 3. Observer Pattern
- Notify when spots become available
- Notify when parking lot is full
- Notify about payment status

### 4. Singleton Pattern
- Parking lot manager
- Payment processor
- Notification service

### 5. Command Pattern
- Parking commands
- Unparking commands
- Payment commands

## Class Diagram

```
Vehicle (Abstract)
├── Car
├── Motorcycle
├── Truck
└── Bus

ParkingSpot (Abstract)
├── CompactSpot
├── RegularSpot
├── LargeSpot
└── HandicappedSpot

ParkingLot
├── spots: List<ParkingSpot>
├── vehicles: Map<String, Vehicle>
└── tickets: Map<String, Ticket>

Ticket
├── vehicle: Vehicle
├── spot: ParkingSpot
├── entryTime: Time
└── exitTime: Time

Payment
├── amount: float
├── method: PaymentMethod
└── status: PaymentStatus

ParkingLotManager
├── parkingLots: List<ParkingLot>
├── parkVehicle()
├── unparkVehicle()
└── findAvailableSpots()
```

## Key Design Decisions

### 1. Vehicle Hierarchy
- Use abstract base class for common vehicle properties
- Each vehicle type has specific size requirements
- Support for future vehicle types

### 2. Spot Allocation Strategy
- **First Available**: Assign first available spot
- **Nearest to Entrance**: Assign closest spot to entrance
- **Size-based**: Assign smallest suitable spot
- **Reserved**: Assign specific reserved spots

### 3. Pricing Strategy
- **Hourly**: Charge per hour
- **Daily**: Charge per day
- **Tiered**: Different rates for different time periods
- **Vehicle Type**: Different rates for different vehicle types

### 4. Concurrency Handling
- Use locks for spot allocation
- Thread-safe operations
- Atomic transactions for critical operations

## API Design

### Core Operations
```go
// Park a vehicle
func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) (*Ticket, error)

// Unpark a vehicle
func (pl *ParkingLot) UnparkVehicle(ticketID string) (*Payment, error)

// Find available spots
func (pl *ParkingLot) FindAvailableSpots(vehicleType VehicleType) []ParkingSpot

// Check if parking lot is full
func (pl *ParkingLot) IsFull() bool

// Get parking lot status
func (pl *ParkingLot) GetStatus() ParkingLotStatus
```

### Vehicle Operations
```go
// Create vehicle
func CreateVehicle(licensePlate string, vehicleType VehicleType) Vehicle

// Get vehicle info
func (v Vehicle) GetInfo() VehicleInfo

// Check if vehicle can fit in spot
func (v Vehicle) CanFitInSpot(spot ParkingSpot) bool
```

### Payment Operations
```go
// Process payment
func (p *PaymentProcessor) ProcessPayment(amount float64, method PaymentMethod) error

// Calculate parking fee
func (p *PricingStrategy) CalculateFee(ticket *Ticket) float64

// Validate payment
func (p *PaymentProcessor) ValidatePayment(payment *Payment) bool
```

## Error Handling

### Common Error Scenarios
1. **Parking Lot Full**: No available spots
2. **Invalid Vehicle**: Unsupported vehicle type
3. **Spot Already Occupied**: Spot is not available
4. **Invalid Ticket**: Ticket not found or expired
5. **Payment Failed**: Payment processing error
6. **Concurrent Access**: Multiple vehicles trying to park simultaneously

### Error Handling Strategy
- Use specific error types for different scenarios
- Provide meaningful error messages
- Implement retry mechanisms for transient errors
- Log errors for debugging and monitoring

## Scalability Considerations

### 1. Multiple Parking Lots
- Support multiple parking lot instances
- Centralized management system
- Load balancing across lots

### 2. Database Design
- Store parking lot data in database
- Use appropriate indexing for fast queries
- Implement data partitioning for large datasets

### 3. Caching
- Cache frequently accessed data
- Use Redis for session management
- Implement cache invalidation strategies

### 4. Microservices
- Split into multiple services
- Parking service, payment service, notification service
- Use message queues for communication

## Testing Strategy

### 1. Unit Tests
- Test individual components
- Mock dependencies
- Test edge cases and error scenarios

### 2. Integration Tests
- Test component interactions
- Test database operations
- Test external service integrations

### 3. Performance Tests
- Test concurrent operations
- Test with large datasets
- Measure response times

### 4. End-to-End Tests
- Test complete user workflows
- Test real-world scenarios
- Test system under load

## Future Enhancements

### 1. Advanced Features
- Spot reservations
- Valet parking
- Electric vehicle charging stations
- Automated payment systems

### 2. Analytics
- Parking utilization reports
- Revenue analytics
- Peak hour analysis
- Customer behavior insights

### 3. Mobile App
- Real-time spot availability
- Mobile payments
- Push notifications
- Navigation to spots

### 4. IoT Integration
- Sensor-based spot detection
- Automated spot allocation
- Real-time monitoring
- Predictive maintenance

## Interview Tips

### 1. Start Simple
- Begin with basic parking lot design
- Add complexity gradually
- Focus on core functionality first

### 2. Ask Clarifying Questions
- What types of vehicles to support?
- What are the pricing requirements?
- How many parking spots?
- Any special requirements?

### 3. Consider Edge Cases
- What happens when parking lot is full?
- How to handle concurrent access?
- What if payment fails?
- How to handle invalid tickets?

### 4. Discuss Trade-offs
- Performance vs. consistency
- Simplicity vs. flexibility
- Memory vs. computation
- Synchronous vs. asynchronous

### 5. Show System Thinking
- Discuss scalability
- Consider monitoring and logging
- Think about error handling
- Plan for future enhancements

## Conclusion

The Parking Lot System is an excellent example of a real-world design problem that tests your understanding of:
- Object-oriented design principles
- Design patterns
- System design concepts
- Error handling
- Concurrency
- Scalability

The key is to start with a simple design and gradually add complexity while maintaining clean, maintainable code. Focus on the core requirements first, then consider edge cases and future enhancements.
