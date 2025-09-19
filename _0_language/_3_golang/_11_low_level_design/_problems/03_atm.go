package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// CORE ENTITIES
// =============================================================================

// Account Types
type AccountType int

const (
	Savings AccountType = iota
	Checking
	Business
)

func (at AccountType) String() string {
	switch at {
	case Savings:
		return "Savings"
	case Checking:
		return "Checking"
	case Business:
		return "Business"
	default:
		return "Unknown"
	}
}

// Transaction Types
type TransactionType int

const (
	BalanceInquiry TransactionType = iota
	Withdrawal
	Deposit
	Transfer
	MiniStatement
)

func (tt TransactionType) String() string {
	switch tt {
	case BalanceInquiry:
		return "Balance Inquiry"
	case Withdrawal:
		return "Withdrawal"
	case Deposit:
		return "Deposit"
	case Transfer:
		return "Transfer"
	case MiniStatement:
		return "Mini Statement"
	default:
		return "Unknown"
	}
}

// Transaction Status
type TransactionStatus int

const (
	Pending TransactionStatus = iota
	Completed
	Failed
	Cancelled
)

func (ts TransactionStatus) String() string {
	switch ts {
	case Pending:
		return "Pending"
	case Completed:
		return "Completed"
	case Failed:
		return "Failed"
	case Cancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

// =============================================================================
// BANK ACCOUNT SYSTEM
// =============================================================================

// Bank Account
type BankAccount struct {
	AccountNumber string
	AccountType   AccountType
	Balance       float64
	Owner         *Customer
	Transactions  []*Transaction
	IsActive      bool
	mu            sync.RWMutex
}

func NewBankAccount(accountNumber string, accountType AccountType, owner *Customer) *BankAccount {
	return &BankAccount{
		AccountNumber: accountNumber,
		AccountType:   accountType,
		Balance:       0.0,
		Owner:         owner,
		Transactions:  make([]*Transaction, 0),
		IsActive:      true,
	}
}

func (ba *BankAccount) GetBalance() float64 {
	ba.mu.RLock()
	defer ba.mu.RUnlock()
	return ba.Balance
}

func (ba *BankAccount) Deposit(amount float64) error {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	
	ba.Balance += amount
	return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	
	if amount > ba.Balance {
		return fmt.Errorf("insufficient funds: requested %.2f, available %.2f", amount, ba.Balance)
	}
	
	ba.Balance -= amount
	return nil
}

func (ba *BankAccount) Transfer(toAccount *BankAccount, amount float64) error {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	
	if amount <= 0 {
		return fmt.Errorf("transfer amount must be positive")
	}
	
	if amount > ba.Balance {
		return fmt.Errorf("insufficient funds for transfer: requested %.2f, available %.2f", amount, ba.Balance)
	}
	
	// Perform atomic transfer
	ba.Balance -= amount
	toAccount.mu.Lock()
	toAccount.Balance += amount
	toAccount.mu.Unlock()
	
	return nil
}

func (ba *BankAccount) AddTransaction(transaction *Transaction) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	ba.Transactions = append(ba.Transactions, transaction)
}

func (ba *BankAccount) GetRecentTransactions(limit int) []*Transaction {
	ba.mu.RLock()
	defer ba.mu.RUnlock()
	
	start := len(ba.Transactions) - limit
	if start < 0 {
		start = 0
	}
	
	return ba.Transactions[start:]
}

func (ba *BankAccount) IsAccountActive() bool {
	ba.mu.RLock()
	defer ba.mu.RUnlock()
	return ba.IsActive
}

func (ba *BankAccount) Deactivate() {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	ba.IsActive = false
}

// Customer
type Customer struct {
	CustomerID   string
	Name         string
	Email        string
	Phone        string
	Address      string
	Accounts     []*BankAccount
	IsActive     bool
	mu           sync.RWMutex
}

func NewCustomer(customerID, name, email, phone, address string) *Customer {
	return &Customer{
		CustomerID: customerID,
		Name:       name,
		Email:      email,
		Phone:      phone,
		Address:    address,
		Accounts:   make([]*BankAccount, 0),
		IsActive:   true,
	}
}

func (c *Customer) AddAccount(account *BankAccount) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Accounts = append(c.Accounts, account)
}

func (c *Customer) GetAccount(accountNumber string) *BankAccount {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	for _, account := range c.Accounts {
		if account.AccountNumber == accountNumber {
			return account
		}
	}
	return nil
}

func (c *Customer) GetAllAccounts() []*BankAccount {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	accounts := make([]*BankAccount, len(c.Accounts))
	copy(accounts, c.Accounts)
	return accounts
}

// =============================================================================
// CARD SYSTEM
// =============================================================================

// Card
type Card struct {
	CardNumber    string
	PIN           string
	AccountNumber string
	ExpiryDate    time.Time
	IsActive      bool
	IsBlocked     bool
	mu            sync.RWMutex
}

func NewCard(cardNumber, pin, accountNumber string, expiryDate time.Time) *Card {
	return &Card{
		CardNumber:    cardNumber,
		PIN:           pin,
		AccountNumber: accountNumber,
		ExpiryDate:    expiryDate,
		IsActive:      true,
		IsBlocked:     false,
	}
}

func (c *Card) ValidatePIN(enteredPIN string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	if !c.IsActive || c.IsBlocked {
		return false
	}
	
	return c.PIN == enteredPIN
}

func (c *Card) IsExpired() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return time.Now().After(c.ExpiryDate)
}

func (c *Card) IsValid() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.IsActive && !c.IsBlocked && !c.IsExpired()
}

func (c *Card) Block() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.IsBlocked = true
}

func (c *Card) Unblock() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.IsBlocked = false
}

// =============================================================================
// TRANSACTION SYSTEM
// =============================================================================

// Transaction
type Transaction struct {
	TransactionID   string
	AccountNumber   string
	Type            TransactionType
	Amount          float64
	Description     string
	Timestamp       time.Time
	Status          TransactionStatus
	RelatedAccount  string // For transfers
}

func NewTransaction(accountNumber string, transactionType TransactionType, amount float64, description string) *Transaction {
	return &Transaction{
		TransactionID: fmt.Sprintf("TXN%d", time.Now().UnixNano()),
		AccountNumber: accountNumber,
		Type:          transactionType,
		Amount:        amount,
		Description:   description,
		Timestamp:     time.Now(),
		Status:        Pending,
	}
}

func (t *Transaction) MarkCompleted() {
	t.Status = Completed
}

func (t *Transaction) MarkFailed() {
	t.Status = Failed
}

func (t *Transaction) MarkCancelled() {
	t.Status = Cancelled
}

// =============================================================================
// ATM STATES
// =============================================================================

// ATM State Interface
type ATMState interface {
	InsertCard(cardNumber string) error
	EnterPIN(pin string) error
	SelectAccount(accountNumber string) error
	SelectTransaction(transactionType TransactionType) error
	EnterAmount(amount float64) error
	ProcessTransaction() error
	EjectCard() error
	GetStateName() string
}

// Idle State
type ATMIdleState struct {
	atm *ATM
}

func NewATMIdleState(atm *ATM) *ATMIdleState {
	return &ATMIdleState{atm: atm}
}

func (ais *ATMIdleState) InsertCard(cardNumber string) error {
	card := ais.atm.BankService.GetCard(cardNumber)
	if card == nil {
		return fmt.Errorf("card not found")
	}
	
	if !card.IsValid() {
		return fmt.Errorf("card is invalid or expired")
	}
	
	ais.atm.CurrentCard = card
	ais.atm.SetState(NewATMCardInsertedState(ais.atm))
	fmt.Println("Card inserted successfully")
	return nil
}

func (ais *ATMIdleState) EnterPIN(pin string) error {
	return fmt.Errorf("please insert card first")
}

func (ais *ATMIdleState) SelectAccount(accountNumber string) error {
	return fmt.Errorf("please insert card first")
}

func (ais *ATMIdleState) SelectTransaction(transactionType TransactionType) error {
	return fmt.Errorf("please insert card first")
}

func (ais *ATMIdleState) EnterAmount(amount float64) error {
	return fmt.Errorf("please insert card first")
}

func (ais *ATMIdleState) ProcessTransaction() error {
	return fmt.Errorf("please insert card first")
}

func (ais *ATMIdleState) EjectCard() error {
	return fmt.Errorf("no card to eject")
}

func (ais *ATMIdleState) GetStateName() string {
	return "Idle"
}

// Card Inserted State
type ATMCardInsertedState struct {
	atm *ATM
}

func NewATMCardInsertedState(atm *ATM) *ATMCardInsertedState {
	return &ATMCardInsertedState{atm: atm}
}

func (cis *ATMCardInsertedState) InsertCard(cardNumber string) error {
	return fmt.Errorf("card already inserted")
}

func (cis *ATMCardInsertedState) EnterPIN(pin string) error {
	if !cis.atm.CurrentCard.ValidatePIN(pin) {
		cis.atm.PinAttempts++
		if cis.atm.PinAttempts >= 3 {
			cis.atm.CurrentCard.Block()
			cis.atm.SetState(NewATMIdleState(cis.atm))
			return fmt.Errorf("card blocked due to too many incorrect PIN attempts")
		}
		return fmt.Errorf("incorrect PIN, %d attempts remaining", 3-cis.atm.PinAttempts)
	}
	
	cis.atm.PinAttempts = 0
	cis.atm.SetState(NewATMPINEnteredState(cis.atm))
	fmt.Println("PIN verified successfully")
	return nil
}

func (cis *ATMCardInsertedState) SelectAccount(accountNumber string) error {
	return fmt.Errorf("please enter PIN first")
}

func (cis *ATMCardInsertedState) SelectTransaction(transactionType TransactionType) error {
	return fmt.Errorf("please enter PIN first")
}

func (cis *ATMCardInsertedState) EnterAmount(amount float64) error {
	return fmt.Errorf("please enter PIN first")
}

func (cis *ATMCardInsertedState) ProcessTransaction() error {
	return fmt.Errorf("please enter PIN first")
}

func (cis *ATMCardInsertedState) EjectCard() error {
	cis.atm.CurrentCard = nil
	cis.atm.PinAttempts = 0
	cis.atm.SetState(NewATMIdleState(cis.atm))
	fmt.Println("Card ejected")
	return nil
}

func (cis *ATMCardInsertedState) GetStateName() string {
	return "Card Inserted"
}

// PIN Entered State
type ATMPINEnteredState struct {
	atm *ATM
}

func NewATMPINEnteredState(atm *ATM) *ATMPINEnteredState {
	return &ATMPINEnteredState{atm: atm}
}

func (pes *ATMPINEnteredState) InsertCard(cardNumber string) error {
	return fmt.Errorf("card already inserted")
}

func (pes *ATMPINEnteredState) EnterPIN(pin string) error {
	return fmt.Errorf("PIN already entered")
}

func (pes *ATMPINEnteredState) SelectAccount(accountNumber string) error {
	account := pes.atm.BankService.GetAccount(accountNumber)
	if account == nil {
		return fmt.Errorf("account not found")
	}
	
	if !account.IsAccountActive() {
		return fmt.Errorf("account is inactive")
	}
	
	pes.atm.CurrentAccount = account
	pes.atm.SetState(NewATMAccountSelectedState(pes.atm))
	fmt.Printf("Account selected: %s\n", accountNumber)
	return nil
}

func (pes *ATMPINEnteredState) SelectTransaction(transactionType TransactionType) error {
	return fmt.Errorf("please select account first")
}

func (pes *ATMPINEnteredState) EnterAmount(amount float64) error {
	return fmt.Errorf("please select account first")
}

func (pes *ATMPINEnteredState) ProcessTransaction() error {
	return fmt.Errorf("please select account first")
}

func (pes *ATMPINEnteredState) EjectCard() error {
	pes.atm.CurrentCard = nil
	pes.atm.CurrentAccount = nil
	pes.atm.PinAttempts = 0
	pes.atm.SetState(NewATMIdleState(pes.atm))
	fmt.Println("Card ejected")
	return nil
}

func (pes *ATMPINEnteredState) GetStateName() string {
	return "PIN Entered"
}

// Account Selected State
type ATMAccountSelectedState struct {
	atm *ATM
}

func NewATMAccountSelectedState(atm *ATM) *ATMAccountSelectedState {
	return &ATMAccountSelectedState{atm: atm}
}

func (ass *ATMAccountSelectedState) InsertCard(cardNumber string) error {
	return fmt.Errorf("card already inserted")
}

func (ass *ATMAccountSelectedState) EnterPIN(pin string) error {
	return fmt.Errorf("PIN already entered")
}

func (ass *ATMAccountSelectedState) SelectAccount(accountNumber string) error {
	return fmt.Errorf("account already selected")
}

func (ass *ATMAccountSelectedState) SelectTransaction(transactionType TransactionType) error {
	ass.atm.CurrentTransactionType = transactionType
	ass.atm.SetState(NewATMTransactionSelectedState(ass.atm))
	fmt.Printf("Transaction selected: %s\n", transactionType)
	return nil
}

func (ass *ATMAccountSelectedState) EnterAmount(amount float64) error {
	return fmt.Errorf("please select transaction first")
}

func (ass *ATMAccountSelectedState) ProcessTransaction() error {
	return fmt.Errorf("please select transaction first")
}

func (ass *ATMAccountSelectedState) EjectCard() error {
	ass.atm.CurrentCard = nil
	ass.atm.CurrentAccount = nil
	ass.atm.PinAttempts = 0
	ass.atm.SetState(NewATMIdleState(ass.atm))
	fmt.Println("Card ejected")
	return nil
}

func (ass *ATMAccountSelectedState) GetStateName() string {
	return "Account Selected"
}

// Transaction Selected State
type ATMTransactionSelectedState struct {
	atm *ATM
}

func NewATMTransactionSelectedState(atm *ATM) *ATMTransactionSelectedState {
	return &ATMTransactionSelectedState{atm: atm}
}

func (tss *ATMTransactionSelectedState) InsertCard(cardNumber string) error {
	return fmt.Errorf("card already inserted")
}

func (tss *ATMTransactionSelectedState) EnterPIN(pin string) error {
	return fmt.Errorf("PIN already entered")
}

func (tss *ATMTransactionSelectedState) SelectAccount(accountNumber string) error {
	return fmt.Errorf("account already selected")
}

func (tss *ATMTransactionSelectedState) SelectTransaction(transactionType TransactionType) error {
	return fmt.Errorf("transaction already selected")
}

func (tss *ATMTransactionSelectedState) EnterAmount(amount float64) error {
	if tss.atm.CurrentTransactionType == BalanceInquiry || tss.atm.CurrentTransactionType == MiniStatement {
		// No amount needed for these transactions
		tss.atm.CurrentAmount = 0
		tss.atm.SetState(NewATMTransactionInProgressState(tss.atm))
		return nil
	}
	
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	
	tss.atm.CurrentAmount = amount
	tss.atm.SetState(NewATMTransactionInProgressState(tss.atm))
	fmt.Printf("Amount entered: $%.2f\n", amount)
	return nil
}

func (tss *ATMTransactionSelectedState) ProcessTransaction() error {
	return fmt.Errorf("please enter amount first")
}

func (tss *ATMTransactionSelectedState) EjectCard() error {
	tss.atm.CurrentCard = nil
	tss.atm.CurrentAccount = nil
	tss.atm.PinAttempts = 0
	tss.atm.SetState(NewATMIdleState(tss.atm))
	fmt.Println("Card ejected")
	return nil
}

func (tss *ATMTransactionSelectedState) GetStateName() string {
	return "Transaction Selected"
}

// Transaction In Progress State
type ATMTransactionInProgressState struct {
	atm *ATM
}

func NewATMTransactionInProgressState(atm *ATM) *ATMTransactionInProgressState {
	return &ATMTransactionInProgressState{atm: atm}
}

func (tips *ATMTransactionInProgressState) InsertCard(cardNumber string) error {
	return fmt.Errorf("transaction in progress")
}

func (tips *ATMTransactionInProgressState) EnterPIN(pin string) error {
	return fmt.Errorf("transaction in progress")
}

func (tips *ATMTransactionInProgressState) SelectAccount(accountNumber string) error {
	return fmt.Errorf("transaction in progress")
}

func (tips *ATMTransactionInProgressState) SelectTransaction(transactionType TransactionType) error {
	return fmt.Errorf("transaction in progress")
}

func (tips *ATMTransactionInProgressState) EnterAmount(amount float64) error {
	return fmt.Errorf("transaction in progress")
}

func (tips *ATMTransactionInProgressState) ProcessTransaction() error {
	// Process the transaction based on type
	var err error
	switch tips.atm.CurrentTransactionType {
	case BalanceInquiry:
		err = tips.processBalanceInquiry()
	case Withdrawal:
		err = tips.processWithdrawal()
	case Deposit:
		err = tips.processDeposit()
	case Transfer:
		err = tips.processTransfer()
	case MiniStatement:
		err = tips.processMiniStatement()
	default:
		err = fmt.Errorf("unknown transaction type")
	}
	
	// If transaction failed, go to error state
	if err != nil {
		tips.atm.SetState(NewATMTransactionErrorState(tips.atm, err))
		return err
	}
	
	return nil
}

func (tips *ATMTransactionInProgressState) processBalanceInquiry() error {
	balance := tips.atm.CurrentAccount.GetBalance()
	fmt.Printf("Account Balance: $%.2f\n", balance)
	
	// Create transaction record
	transaction := NewTransaction(tips.atm.CurrentAccount.AccountNumber, BalanceInquiry, 0, "Balance Inquiry")
	transaction.MarkCompleted()
	tips.atm.CurrentAccount.AddTransaction(transaction)
	
	tips.atm.SetState(NewATMTransactionCompletedState(tips.atm))
	return nil
}

func (tips *ATMTransactionInProgressState) processWithdrawal() error {
	// Check if ATM has enough cash
	if tips.atm.CashDrawer.GetTotalAmount() < tips.atm.CurrentAmount {
		return fmt.Errorf("insufficient cash in ATM")
	}
	
	// Check account balance
	if tips.atm.CurrentAccount.GetBalance() < tips.atm.CurrentAmount {
		return fmt.Errorf("insufficient funds in account")
	}
	
	// Dispense cash first
	err := tips.atm.CashDrawer.DispenseCash(tips.atm.CurrentAmount)
	if err != nil {
		return fmt.Errorf("failed to dispense cash: %v", err)
	}
	
	// Process withdrawal after successful cash dispensing
	err = tips.atm.CurrentAccount.Withdraw(tips.atm.CurrentAmount)
	if err != nil {
		// Rollback cash dispensing (simplified - in real system would need to return cash)
		tips.atm.CashDrawer.AddCash(int(tips.atm.CurrentAmount*100), 1) // Add back as cents
		return fmt.Errorf("failed to process withdrawal: %v", err)
	}
	
	// Create transaction record
	transaction := NewTransaction(tips.atm.CurrentAccount.AccountNumber, Withdrawal, tips.atm.CurrentAmount, "Cash Withdrawal")
	transaction.MarkCompleted()
	tips.atm.CurrentAccount.AddTransaction(transaction)
	
	fmt.Printf("Withdrawal successful: $%.2f\n", tips.atm.CurrentAmount)
	tips.atm.SetState(NewATMTransactionCompletedState(tips.atm))
	return nil
}

func (tips *ATMTransactionInProgressState) processDeposit() error {
	// Process deposit
	err := tips.atm.CurrentAccount.Deposit(tips.atm.CurrentAmount)
	if err != nil {
		return err
	}
	
	// Create transaction record
	transaction := NewTransaction(tips.atm.CurrentAccount.AccountNumber, Deposit, tips.atm.CurrentAmount, "Cash Deposit")
	transaction.MarkCompleted()
	tips.atm.CurrentAccount.AddTransaction(transaction)
	
	fmt.Printf("Deposit successful: $%.2f\n", tips.atm.CurrentAmount)
	tips.atm.SetState(NewATMTransactionCompletedState(tips.atm))
	return nil
}

func (tips *ATMTransactionInProgressState) processTransfer() error {
	// For transfer, we need a destination account
	// This is simplified - in real implementation, user would select destination account
	return fmt.Errorf("transfer functionality not implemented in this demo")
}

func (tips *ATMTransactionInProgressState) processMiniStatement() error {
	transactions := tips.atm.CurrentAccount.GetRecentTransactions(5)
	fmt.Println("Mini Statement:")
	fmt.Println("Date\t\tType\t\tAmount\t\tDescription")
	fmt.Println("------------------------------------------------")
	
	for _, txn := range transactions {
		fmt.Printf("%s\t%s\t$%.2f\t\t%s\n", 
			txn.Timestamp.Format("2006-01-02 15:04"), 
			txn.Type, 
			txn.Amount, 
			txn.Description)
	}
	
	// Create transaction record
	transaction := NewTransaction(tips.atm.CurrentAccount.AccountNumber, MiniStatement, 0, "Mini Statement")
	transaction.MarkCompleted()
	tips.atm.CurrentAccount.AddTransaction(transaction)
	
	tips.atm.SetState(NewATMTransactionCompletedState(tips.atm))
	return nil
}

func (tips *ATMTransactionInProgressState) EjectCard() error {
	return fmt.Errorf("transaction in progress")
}

func (tips *ATMTransactionInProgressState) GetStateName() string {
	return "Transaction In Progress"
}

// Transaction Completed State
type ATMTransactionCompletedState struct {
	atm *ATM
}

func NewATMTransactionCompletedState(atm *ATM) *ATMTransactionCompletedState {
	return &ATMTransactionCompletedState{atm: atm}
}

func (tcs *ATMTransactionCompletedState) InsertCard(cardNumber string) error {
	return fmt.Errorf("transaction completed, please eject card")
}

func (tcs *ATMTransactionCompletedState) EnterPIN(pin string) error {
	return fmt.Errorf("transaction completed, please eject card")
}

func (tcs *ATMTransactionCompletedState) SelectAccount(accountNumber string) error {
	return fmt.Errorf("transaction completed, please eject card")
}

func (tcs *ATMTransactionCompletedState) SelectTransaction(transactionType TransactionType) error {
	return fmt.Errorf("transaction completed, please eject card")
}

func (tcs *ATMTransactionCompletedState) EnterAmount(amount float64) error {
	return fmt.Errorf("transaction completed, please eject card")
}

func (tcs *ATMTransactionCompletedState) ProcessTransaction() error {
	return fmt.Errorf("transaction completed, please eject card")
}

func (tcs *ATMTransactionCompletedState) EjectCard() error {
	tcs.atm.CurrentCard = nil
	tcs.atm.CurrentAccount = nil
	tcs.atm.PinAttempts = 0
	tcs.atm.SetState(NewATMIdleState(tcs.atm))
	fmt.Println("Transaction completed. Card ejected.")
	return nil
}

func (tcs *ATMTransactionCompletedState) GetStateName() string {
	return "Transaction Completed"
}

// Transaction Error State
type ATMTransactionErrorState struct {
	atm *ATM
	err error
}

func NewATMTransactionErrorState(atm *ATM, err error) *ATMTransactionErrorState {
	return &ATMTransactionErrorState{atm: atm, err: err}
}

func (tes *ATMTransactionErrorState) InsertCard(cardNumber string) error {
	return fmt.Errorf("transaction error, please eject card")
}

func (tes *ATMTransactionErrorState) EnterPIN(pin string) error {
	return fmt.Errorf("transaction error, please eject card")
}

func (tes *ATMTransactionErrorState) SelectAccount(accountNumber string) error {
	return fmt.Errorf("transaction error, please eject card")
}

func (tes *ATMTransactionErrorState) SelectTransaction(transactionType TransactionType) error {
	return fmt.Errorf("transaction error, please eject card")
}

func (tes *ATMTransactionErrorState) EnterAmount(amount float64) error {
	return fmt.Errorf("transaction error, please eject card")
}

func (tes *ATMTransactionErrorState) ProcessTransaction() error {
	return fmt.Errorf("transaction error, please eject card")
}

func (tes *ATMTransactionErrorState) EjectCard() error {
	tes.atm.CurrentCard = nil
	tes.atm.CurrentAccount = nil
	tes.atm.PinAttempts = 0
	tes.atm.SetState(NewATMIdleState(tes.atm))
	fmt.Printf("Transaction failed: %v. Card ejected.\n", tes.err)
	return nil
}

func (tes *ATMTransactionErrorState) GetStateName() string {
	return "Transaction Error"
}

// =============================================================================
// CASH DRAWER SYSTEM
// =============================================================================

// Cash Drawer
type ATMCashDrawer struct {
	Denominations map[int]int // denomination -> count
	TotalAmount   float64
	mu            sync.RWMutex
}

func NewATMCashDrawer() *ATMCashDrawer {
	return &ATMCashDrawer{
		Denominations: make(map[int]int),
		TotalAmount:   0.0,
	}
}

func (cd *ATMCashDrawer) AddCash(denomination int, count int) {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	
	cd.Denominations[denomination] += count
	cd.TotalAmount += float64(denomination * count)
}

func (cd *ATMCashDrawer) GetTotalAmount() float64 {
	cd.mu.RLock()
	defer cd.mu.RUnlock()
	return cd.TotalAmount
}

func (cd *ATMCashDrawer) DispenseCash(amount float64) error {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	
	if amount > cd.TotalAmount {
		return fmt.Errorf("insufficient cash in ATM")
	}
	
	remaining := int(amount * 100) // Convert to cents
	
	// Try to dispense using available denominations (in cents)
	denominations := []int{2000, 1000, 500, 200, 100, 50, 20, 10, 5, 1} // In cents
	
	for _, denom := range denominations {
		if remaining <= 0 {
			break
		}
		
		count := remaining / denom
		available := cd.Denominations[denom]
		
		if count > available {
			count = available
		}
		
		if count > 0 {
			cd.Denominations[denom] -= count
			remaining -= count * denom
		}
	}
	
	if remaining > 0 {
		return fmt.Errorf("cannot dispense exact amount with available denominations")
	}
	
	cd.TotalAmount -= amount
	return nil
}

func (cd *ATMCashDrawer) GetDenominationCount(denomination int) int {
	cd.mu.RLock()
	defer cd.mu.RUnlock()
	return cd.Denominations[denomination]
}

// =============================================================================
// BANK SERVICE
// =============================================================================

// Bank Service
type BankService struct {
	Accounts map[string]*BankAccount
	Cards    map[string]*Card
	Customers map[string]*Customer
	mu       sync.RWMutex
}

func NewBankService() *BankService {
	return &BankService{
		Accounts:  make(map[string]*BankAccount),
		Cards:     make(map[string]*Card),
		Customers: make(map[string]*Customer),
	}
}

func (bs *BankService) AddCustomer(customer *Customer) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.Customers[customer.CustomerID] = customer
}

func (bs *BankService) AddAccount(account *BankAccount) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.Accounts[account.AccountNumber] = account
}

func (bs *BankService) AddCard(card *Card) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.Cards[card.CardNumber] = card
}

func (bs *BankService) GetAccount(accountNumber string) *BankAccount {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	return bs.Accounts[accountNumber]
}

func (bs *BankService) GetCard(cardNumber string) *Card {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	return bs.Cards[cardNumber]
}

func (bs *BankService) GetCustomer(customerID string) *Customer {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	return bs.Customers[customerID]
}

// =============================================================================
// ATM SYSTEM
// =============================================================================

// ATM
type ATM struct {
	ID                    string
	Location              string
	BankService           *BankService
	CashDrawer            *ATMCashDrawer
	CurrentState          ATMState
	CurrentCard           *Card
	CurrentAccount        *BankAccount
	CurrentTransactionType TransactionType
	CurrentAmount         float64
	PinAttempts           int
	mu                    sync.RWMutex
}

func NewATM(id, location string, bankService *BankService) *ATM {
	atm := &ATM{
		ID:          id,
		Location:    location,
		BankService: bankService,
		CashDrawer:  NewATMCashDrawer(),
		PinAttempts: 0,
	}
	atm.SetState(NewATMIdleState(atm))
	return atm
}

func (atm *ATM) SetState(state ATMState) {
	atm.mu.Lock()
	defer atm.mu.Unlock()
	atm.CurrentState = state
}

func (atm *ATM) GetState() ATMState {
	atm.mu.RLock()
	defer atm.mu.RUnlock()
	return atm.CurrentState
}

func (atm *ATM) InsertCard(cardNumber string) error {
	return atm.GetState().InsertCard(cardNumber)
}

func (atm *ATM) EnterPIN(pin string) error {
	return atm.GetState().EnterPIN(pin)
}

func (atm *ATM) SelectAccount(accountNumber string) error {
	return atm.GetState().SelectAccount(accountNumber)
}

func (atm *ATM) SelectTransaction(transactionType TransactionType) error {
	return atm.GetState().SelectTransaction(transactionType)
}

func (atm *ATM) EnterAmount(amount float64) error {
	return atm.GetState().EnterAmount(amount)
}

func (atm *ATM) ProcessTransaction() error {
	return atm.GetState().ProcessTransaction()
}

func (atm *ATM) EjectCard() error {
	return atm.GetState().EjectCard()
}

func (atm *ATM) GetATMStatus() ATMStatus {
	atm.mu.RLock()
	defer atm.mu.RUnlock()
	
	return ATMStatus{
		ID:           atm.ID,
		Location:     atm.Location,
		State:        atm.GetState().GetStateName(),
		CashAmount:   atm.CashDrawer.GetTotalAmount(),
		IsOperational: true,
	}
}

type ATMStatus struct {
	ID            string
	Location      string
	State         string
	CashAmount    float64
	IsOperational bool
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== ATM SYSTEM DEMONSTRATION ===\n")

	// Create bank service
	bankService := NewBankService()
	
	// Create customer
	customer := NewCustomer("C001", "John Doe", "john@example.com", "123-456-7890", "123 Main St")
	bankService.AddCustomer(customer)
	
	// Create accounts
	savingsAccount := NewBankAccount("ACC001", Savings, customer)
	checkingAccount := NewBankAccount("ACC002", Checking, customer)
	
	// Add initial balance
	savingsAccount.Deposit(1000.0)
	checkingAccount.Deposit(500.0)
	
	bankService.AddAccount(savingsAccount)
	bankService.AddAccount(checkingAccount)
	customer.AddAccount(savingsAccount)
	customer.AddAccount(checkingAccount)
	
	// Create cards
	card1 := NewCard("1234567890123456", "1234", "ACC001", time.Now().Add(2*365*24*time.Hour))
	card2 := NewCard("9876543210987654", "5678", "ACC002", time.Now().Add(2*365*24*time.Hour))
	
	bankService.AddCard(card1)
	bankService.AddCard(card2)
	
	// Create ATM
	atm := NewATM("ATM001", "Downtown Branch", bankService)
	
	// Initialize cash drawer (amounts in cents)
	atm.CashDrawer.AddCash(2000, 10)  // $20 bills
	atm.CashDrawer.AddCash(1000, 20)  // $10 bills
	atm.CashDrawer.AddCash(500, 50)   // $5 bills
	atm.CashDrawer.AddCash(100, 100)  // $1 bills
	atm.CashDrawer.AddCash(50, 200)   // $0.50 coins
	atm.CashDrawer.AddCash(25, 400)   // $0.25 coins
	
	// Display initial status
	fmt.Println("1. INITIAL ATM STATUS:")
	status := atm.GetATMStatus()
	fmt.Printf("ATM ID: %s\n", status.ID)
	fmt.Printf("Location: %s\n", status.Location)
	fmt.Printf("State: %s\n", status.State)
	fmt.Printf("Cash Available: $%.2f\n", status.CashAmount)
	fmt.Printf("Is Operational: %t\n", status.IsOperational)
	
	// Display account balances
	fmt.Printf("\nAccount Balances:\n")
	fmt.Printf("Savings Account (ACC001): $%.2f\n", savingsAccount.GetBalance())
	fmt.Printf("Checking Account (ACC002): $%.2f\n", checkingAccount.GetBalance())
	
	fmt.Println()
	
	// Test successful transaction flow
	fmt.Println("2. SUCCESSFUL TRANSACTION FLOW:")
	
	// Insert card
	err := atm.InsertCard("1234567890123456")
	if err != nil {
		fmt.Printf("Error inserting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter PIN
	err = atm.EnterPIN("1234")
	if err != nil {
		fmt.Printf("Error entering PIN: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select account
	err = atm.SelectAccount("ACC001")
	if err != nil {
		fmt.Printf("Error selecting account: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select transaction
	err = atm.SelectTransaction(BalanceInquiry)
	if err != nil {
		fmt.Printf("Error selecting transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter amount (0 for balance inquiry)
	err = atm.EnterAmount(0)
	if err != nil {
		fmt.Printf("Error entering amount: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Process transaction
	err = atm.ProcessTransaction()
	if err != nil {
		fmt.Printf("Error processing transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Eject card
	err = atm.EjectCard()
	if err != nil {
		fmt.Printf("Error ejecting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	fmt.Println()
	
	// Test withdrawal transaction
	fmt.Println("3. WITHDRAWAL TRANSACTION:")
	
	// Insert card
	err = atm.InsertCard("1234567890123456")
	if err != nil {
		fmt.Printf("Error inserting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter PIN
	err = atm.EnterPIN("1234")
	if err != nil {
		fmt.Printf("Error entering PIN: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select account
	err = atm.SelectAccount("ACC001")
	if err != nil {
		fmt.Printf("Error selecting account: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select withdrawal transaction
	err = atm.SelectTransaction(Withdrawal)
	if err != nil {
		fmt.Printf("Error selecting transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter amount
	err = atm.EnterAmount(100.0)
	if err != nil {
		fmt.Printf("Error entering amount: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Process transaction
	err = atm.ProcessTransaction()
	if err != nil {
		fmt.Printf("Error processing transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Eject card
	err = atm.EjectCard()
	if err != nil {
		fmt.Printf("Error ejecting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	fmt.Println()
	
	// Test deposit transaction
	fmt.Println("4. DEPOSIT TRANSACTION:")
	
	// Insert card
	err = atm.InsertCard("9876543210987654")
	if err != nil {
		fmt.Printf("Error inserting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter PIN
	err = atm.EnterPIN("5678")
	if err != nil {
		fmt.Printf("Error entering PIN: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select account
	err = atm.SelectAccount("ACC002")
	if err != nil {
		fmt.Printf("Error selecting account: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select deposit transaction
	err = atm.SelectTransaction(Deposit)
	if err != nil {
		fmt.Printf("Error selecting transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter amount
	err = atm.EnterAmount(200.0)
	if err != nil {
		fmt.Printf("Error entering amount: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Process transaction
	err = atm.ProcessTransaction()
	if err != nil {
		fmt.Printf("Error processing transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Eject card
	err = atm.EjectCard()
	if err != nil {
		fmt.Printf("Error ejecting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	fmt.Println()
	
	// Test mini statement
	fmt.Println("5. MINI STATEMENT TRANSACTION:")
	
	// Insert card
	err = atm.InsertCard("1234567890123456")
	if err != nil {
		fmt.Printf("Error inserting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter PIN
	err = atm.EnterPIN("1234")
	if err != nil {
		fmt.Printf("Error entering PIN: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select account
	err = atm.SelectAccount("ACC001")
	if err != nil {
		fmt.Printf("Error selecting account: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Select mini statement transaction
	err = atm.SelectTransaction(MiniStatement)
	if err != nil {
		fmt.Printf("Error selecting transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Enter amount (0 for mini statement)
	err = atm.EnterAmount(0)
	if err != nil {
		fmt.Printf("Error entering amount: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Process transaction
	err = atm.ProcessTransaction()
	if err != nil {
		fmt.Printf("Error processing transaction: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	// Eject card
	err = atm.EjectCard()
	if err != nil {
		fmt.Printf("Error ejecting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	fmt.Println()
	
	// Test error scenarios
	fmt.Println("6. ERROR SCENARIOS:")
	
	// Test incorrect PIN
	err = atm.InsertCard("1234567890123456")
	if err != nil {
		fmt.Printf("Error inserting card: %v\n", err)
	} else {
		fmt.Printf("Current state: %s\n", atm.GetState().GetStateName())
	}
	
	err = atm.EnterPIN("9999") // Wrong PIN
	if err != nil {
		fmt.Printf("Expected error for wrong PIN: %v\n", err)
	}
	
	// Eject card
	atm.EjectCard()
	
	// Test invalid card
	err = atm.InsertCard("0000000000000000")
	if err != nil {
		fmt.Printf("Expected error for invalid card: %v\n", err)
	}
	
	// Test operations without card
	err = atm.EnterPIN("1234")
	if err != nil {
		fmt.Printf("Expected error for PIN without card: %v\n", err)
	}
	
	fmt.Println()
	
	// Display final status
	fmt.Println("7. FINAL STATUS:")
	status = atm.GetATMStatus()
	fmt.Printf("ATM State: %s\n", status.State)
	fmt.Printf("Cash Available: $%.2f\n", status.CashAmount)
	
	fmt.Printf("\nFinal Account Balances:\n")
	fmt.Printf("Savings Account (ACC001): $%.2f\n", savingsAccount.GetBalance())
	fmt.Printf("Checking Account (ACC002): $%.2f\n", checkingAccount.GetBalance())
	
	fmt.Println()
	fmt.Println("=== END OF DEMONSTRATION ===")
}
