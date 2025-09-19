# Vending Machine System Design

## Problem Statement
Design a vending machine system that can dispense different types of products (snacks, drinks, etc.) and accept different forms of payment (coins, bills, cards). The system should handle inventory management, payment processing, and change calculation.

## Requirements Analysis

### Functional Requirements
1. **Product Management**
   - Display available products and prices
   - Check product availability
   - Dispense products when payment is complete
   - Handle product selection and validation

2. **Payment Processing**
   - Accept coins and bills
   - Process credit/debit card payments
   - Calculate and return change
   - Handle insufficient funds scenarios

3. **Inventory Management**
   - Track product quantities
   - Restock products
   - Handle out-of-stock scenarios
   - Monitor product expiration dates

4. **Transaction Management**
   - Record all transactions
   - Handle refunds and cancellations
   - Generate receipts
   - Track sales analytics

### Non-Functional Requirements
1. **Reliability**: System should be fault-tolerant
2. **Security**: Secure payment processing
3. **Usability**: Simple and intuitive interface
4. **Maintainability**: Easy to add new products and payment methods

## Core Entities

### 1. Product
- **Attributes**: ID, name, price, quantity, category, expiration date
- **Types**: Snack, Drink, Candy, Gum
- **Behavior**: Check availability, update quantity, validate expiration

### 2. Vending Machine
- **Attributes**: ID, location, products, cash drawer, card reader
- **Behavior**: Display products, process payments, dispense products

### 3. Payment
- **Attributes**: Amount, method, status, timestamp
- **Types**: Cash, Card, Mobile Payment
- **Behavior**: Process payment, validate amount, calculate change

### 4. Transaction
- **Attributes**: ID, product, payment, timestamp, status
- **Behavior**: Record transaction, generate receipt, handle refunds

### 5. Cash Drawer
- **Attributes**: Coins, bills, total amount
- **Behavior**: Accept money, dispense change, calculate balance

## Design Patterns Used

### 1. State Pattern
- Different states: Idle, ProductSelected, PaymentInProgress, Dispensing
- State transitions based on user actions
- Encapsulate behavior for each state

### 2. Strategy Pattern
- Different payment strategies (cash, card, mobile)
- Different pricing strategies
- Different inventory management strategies

### 3. Observer Pattern
- Notify when products are out of stock
- Notify when maintenance is required
- Notify about transaction completion

### 4. Factory Pattern
- Create different types of products
- Create different types of payments
- Create different types of transactions

### 5. Command Pattern
- Product selection commands
- Payment commands
- Refund commands

## Class Diagram

```
VendingMachine
├── products: Map<String, Product>
├── cashDrawer: CashDrawer
├── cardReader: CardReader
├── currentState: VendingMachineState
└── currentTransaction: Transaction

Product
├── id: String
├── name: String
├── price: double
├── quantity: int
└── category: ProductCategory

VendingMachineState (Abstract)
├── IdleState
├── ProductSelectedState
├── PaymentInProgressState
└── DispensingState

Payment
├── amount: double
├── method: PaymentMethod
└── status: PaymentStatus

CashDrawer
├── coins: Map<CoinType, Integer>
├── bills: Map<BillType, Integer>
└── totalAmount: double
```

## Key Design Decisions

### 1. State Management
- Use State pattern to manage machine states
- Clear state transitions and validations
- Handle state-specific operations

### 2. Payment Processing
- Support multiple payment methods
- Secure payment validation
- Proper change calculation

### 3. Inventory Management
- Real-time inventory tracking
- Automatic restocking alerts
- Product expiration management

### 4. Error Handling
- Graceful error handling for all scenarios
- User-friendly error messages
- Automatic recovery mechanisms

## API Design

### Core Operations
```go
// Select a product
func (vm *VendingMachine) SelectProduct(productID string) error

// Insert money
func (vm *VendingMachine) InsertMoney(amount float64) error

// Process payment
func (vm *VendingMachine) ProcessPayment() (*Transaction, error)

// Dispense product
func (vm *VendingMachine) DispenseProduct() (*Product, error)

// Cancel transaction
func (vm *VendingMachine) CancelTransaction() error

// Get available products
func (vm *VendingMachine) GetAvailableProducts() []Product
```

### Product Operations
```go
// Add product to inventory
func (vm *VendingMachine) AddProduct(product Product, quantity int) error

// Update product quantity
func (vm *VendingMachine) UpdateProductQuantity(productID string, quantity int) error

// Check product availability
func (vm *VendingMachine) IsProductAvailable(productID string) bool

// Get product details
func (vm *VendingMachine) GetProduct(productID string) (*Product, error)
```

### Payment Operations
```go
// Process cash payment
func (vm *VendingMachine) ProcessCashPayment(amount float64) error

// Process card payment
func (vm *VendingMachine) ProcessCardPayment(cardNumber string) error

// Calculate change
func (vm *VendingMachine) CalculateChange(amountPaid, productPrice float64) float64

// Dispense change
func (vm *VendingMachine) DispenseChange(amount float64) error
```

## Error Handling

### Common Error Scenarios
1. **Product Out of Stock**: Selected product not available
2. **Insufficient Funds**: Not enough money inserted
3. **Invalid Product**: Product ID not found
4. **Payment Failed**: Card payment declined
5. **Machine Malfunction**: Hardware failure
6. **Network Error**: Card reader connection issue

### Error Handling Strategy
- Use specific error types for different scenarios
- Provide clear error messages to users
- Implement retry mechanisms for transient errors
- Log errors for debugging and monitoring

## Security Considerations

### 1. Payment Security
- Encrypt card information
- Validate payment methods
- Secure communication with payment processors

### 2. Physical Security
- Secure cash drawer
- Tamper-proof product storage
- Surveillance and monitoring

### 3. Data Security
- Encrypt sensitive data
- Secure transaction logs
- Regular security audits

## Testing Strategy

### 1. Unit Tests
- Test individual components
- Mock external dependencies
- Test edge cases and error scenarios

### 2. Integration Tests
- Test component interactions
- Test payment processing
- Test inventory management

### 3. End-to-End Tests
- Test complete user workflows
- Test real-world scenarios
- Test system under load

### 4. Security Tests
- Test payment security
- Test data encryption
- Test access controls

## Future Enhancements

### 1. Advanced Features
- Touch screen interface
- Voice commands
- Mobile app integration
- Loyalty programs

### 2. Analytics
- Sales analytics
- Product performance metrics
- Customer behavior analysis
- Predictive maintenance

### 3. IoT Integration
- Remote monitoring
- Predictive restocking
- Real-time inventory tracking
- Automated maintenance alerts

### 4. AI Features
- Product recommendations
- Demand forecasting
- Dynamic pricing
- Fraud detection

## Interview Tips

### 1. Start Simple
- Begin with basic vending machine functionality
- Add complexity gradually
- Focus on core requirements first

### 2. Ask Clarifying Questions
- What types of products to support?
- What payment methods to accept?
- How to handle inventory management?
- Any special requirements?

### 3. Consider Edge Cases
- What happens when product is out of stock?
- How to handle payment failures?
- What if machine runs out of change?
- How to handle power outages?

### 4. Discuss Trade-offs
- Performance vs. security
- Simplicity vs. flexibility
- Cost vs. features
- Reliability vs. complexity

### 5. Show System Thinking
- Discuss scalability
- Consider maintenance and monitoring
- Think about error handling
- Plan for future enhancements

## Conclusion

The Vending Machine System is an excellent example of a real-world design problem that tests your understanding of:
- State management
- Payment processing
- Inventory management
- Error handling
- Security considerations
- User experience design

The key is to start with a simple design and gradually add complexity while maintaining clean, maintainable code. Focus on the core requirements first, then consider edge cases and future enhancements.
