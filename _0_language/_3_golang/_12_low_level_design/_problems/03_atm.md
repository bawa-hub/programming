# ATM System Design

## Problem Statement
Design an ATM (Automated Teller Machine) system that allows users to perform banking operations like checking balance, withdrawing cash, depositing money, and transferring funds. The system should handle multiple accounts, security, and transaction logging.

## Requirements Analysis

### Functional Requirements
1. **Authentication**
   - Card insertion and validation
   - PIN verification
   - Account verification
   - Session management

2. **Account Operations**
   - Check account balance
   - Withdraw cash
   - Deposit money
   - Transfer funds between accounts
   - View transaction history

3. **Cash Management**
   - Dispense cash in different denominations
   - Accept cash deposits
   - Count and validate cash
   - Handle insufficient cash scenarios

4. **Transaction Management**
   - Record all transactions
   - Generate receipts
   - Handle transaction failures
   - Support transaction reversal

### Non-Functional Requirements
1. **Security**: Secure authentication and transaction processing
2. **Reliability**: High availability and fault tolerance
3. **Performance**: Fast response times for all operations
4. **Scalability**: Support multiple ATMs and accounts

## Core Entities

### 1. ATM
- **Attributes**: ID, location, cash available, status
- **Behavior**: Process transactions, dispense cash, accept deposits

### 2. Bank Account
- **Attributes**: Account number, balance, account type, owner
- **Behavior**: Update balance, validate transactions, check limits

### 3. Card
- **Attributes**: Card number, PIN, account number, expiry date
- **Behavior**: Validate PIN, check expiry, link to account

### 4. Transaction
- **Attributes**: ID, type, amount, timestamp, status
- **Types**: Withdrawal, Deposit, Transfer, Balance Inquiry
- **Behavior**: Record transaction, validate amount, update status

### 5. User Session
- **Attributes**: Session ID, card, start time, end time
- **Behavior**: Manage session lifecycle, track activities

## Design Patterns Used

### 1. State Pattern
- Different ATM states: Idle, CardInserted, PINEntered, TransactionInProgress
- State transitions based on user actions
- Encapsulate behavior for each state

### 2. Strategy Pattern
- Different transaction strategies (withdrawal, deposit, transfer)
- Different authentication strategies
- Different cash dispensing strategies

### 3. Observer Pattern
- Notify bank about transactions
- Notify about security events
- Notify about maintenance requirements

### 4. Factory Pattern
- Create different types of transactions
- Create different types of cards
- Create different types of accounts

### 5. Command Pattern
- Transaction commands
- Authentication commands
- Cash management commands

## Class Diagram

```
ATM
├── id: String
├── location: String
├── cashDrawer: CashDrawer
├── cardReader: CardReader
├── currentState: ATMState
├── currentSession: UserSession
└── bankService: BankService

BankAccount
├── accountNumber: String
├── balance: double
├── accountType: AccountType
├── owner: Customer
└── transactions: List<Transaction>

Card
├── cardNumber: String
├── pin: String
├── accountNumber: String
├── expiryDate: Date
└── isActive: boolean

Transaction
├── id: String
├── type: TransactionType
├── amount: double
├── timestamp: Date
├── status: TransactionStatus
└── accountNumber: String
```

## Key Design Decisions

### 1. Authentication Flow
- Card insertion → PIN verification → Account validation
- Session-based authentication
- Secure PIN handling and validation

### 2. Transaction Processing
- Atomic transaction processing
- Rollback on failure
- Transaction logging and auditing

### 3. Cash Management
- Multiple denomination support
- Optimal cash dispensing algorithm
- Cash level monitoring and alerts

### 4. Error Handling
- Comprehensive error handling
- User-friendly error messages
- Automatic recovery mechanisms

## API Design

### Core Operations
```go
// Insert card
func (atm *ATM) InsertCard(cardNumber string) error

// Enter PIN
func (atm *ATM) EnterPIN(pin string) error

// Check balance
func (atm *ATM) CheckBalance() (float64, error)

// Withdraw cash
func (atm *ATM) WithdrawCash(amount float64) error

// Deposit money
func (atm *ATM) DepositMoney(amount float64) error

// Transfer funds
func (atm *ATM) TransferFunds(toAccount string, amount float64) error

// End session
func (atm *ATM) EndSession() error
```

### Account Operations
```go
// Get account details
func (atm *ATM) GetAccountDetails() (*Account, error)

// Get transaction history
func (atm *ATM) GetTransactionHistory() ([]*Transaction, error)

// Validate transaction
func (atm *ATM) ValidateTransaction(amount float64, type TransactionType) error
```

### Cash Management
```go
// Dispense cash
func (atm *ATM) DispenseCash(amount float64) error

// Accept cash deposit
func (atm *ATM) AcceptCashDeposit(amount float64) error

// Get cash levels
func (atm *ATM) GetCashLevels() map[Denomination]int

// Refill cash
func (atm *ATM) RefillCash(denominations map[Denomination]int) error
```

## Security Considerations

### 1. Authentication Security
- Encrypt PIN during transmission
- Limit PIN attempts
- Session timeout
- Card validation

### 2. Transaction Security
- Encrypt transaction data
- Validate transaction limits
- Monitor for suspicious activity
- Secure communication with bank

### 3. Physical Security
- Secure cash storage
- Tamper detection
- Surveillance monitoring
- Access control

### 4. Data Security
- Encrypt sensitive data
- Secure data transmission
- Regular security audits
- Compliance with regulations

## Error Handling

### Common Error Scenarios
1. **Invalid Card**: Card not recognized or expired
2. **Invalid PIN**: Incorrect PIN entered
3. **Insufficient Funds**: Account balance too low
4. **Insufficient Cash**: ATM out of cash
5. **Network Error**: Communication with bank failed
6. **Card Blocked**: Card is blocked or frozen

### Error Handling Strategy
- Use specific error types for different scenarios
- Provide clear error messages to users
- Implement retry mechanisms for transient errors
- Log errors for debugging and monitoring

## Testing Strategy

### 1. Unit Tests
- Test individual components
- Mock external dependencies
- Test edge cases and error scenarios

### 2. Integration Tests
- Test component interactions
- Test bank service integration
- Test cash management systems

### 3. Security Tests
- Test authentication mechanisms
- Test transaction security
- Test data encryption

### 4. Performance Tests
- Test response times
- Test concurrent access
- Test under load

## Future Enhancements

### 1. Advanced Features
- Biometric authentication
- Mobile app integration
- Contactless payments
- Multi-language support

### 2. Analytics
- Transaction analytics
- Usage patterns
- Performance metrics
- Predictive maintenance

### 3. IoT Integration
- Remote monitoring
- Predictive maintenance
- Real-time alerts
- Automated refilling

### 4. AI Features
- Fraud detection
- Personalized services
- Predictive analytics
- Smart cash management

## Interview Tips

### 1. Start Simple
- Begin with basic ATM functionality
- Add complexity gradually
- Focus on core requirements first

### 2. Ask Clarifying Questions
- What types of accounts to support?
- What are the transaction limits?
- How to handle cash management?
- Any special security requirements?

### 3. Consider Edge Cases
- What happens when ATM runs out of cash?
- How to handle network failures?
- What if card is blocked?
- How to handle concurrent access?

### 4. Discuss Trade-offs
- Security vs. usability
- Performance vs. reliability
- Cost vs. features
- Simplicity vs. flexibility

### 5. Show System Thinking
- Discuss scalability
- Consider monitoring and logging
- Think about error handling
- Plan for future enhancements

## Conclusion

The ATM System is an excellent example of a real-world design problem that tests your understanding of:
- State management
- Security considerations
- Transaction processing
- Error handling
- System integration
- User experience design

The key is to start with a simple design and gradually add complexity while maintaining security and reliability. Focus on the core requirements first, then consider edge cases and future enhancements.
