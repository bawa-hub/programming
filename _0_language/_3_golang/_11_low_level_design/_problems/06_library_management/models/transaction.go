package models

import (
	"fmt"
	"sync"
	"time"
)

// Transaction Type
type TransactionType int

const (
	Checkout TransactionType = iota
	Return
	Renewal
	Reservation
	Fine
)

func (tt TransactionType) String() string {
	switch tt {
	case Checkout:
		return "Checkout"
	case Return:
		return "Return"
	case Renewal:
		return "Renewal"
	case Reservation:
		return "Reservation"
	case Fine:
		return "Fine"
	default:
		return "Unknown"
	}
}

// Transaction Status
type TransactionStatus int

const (
	TransactionActive TransactionStatus = iota
	TransactionCompleted
	TransactionCancelled
	TransactionOverdue
)

func (ts TransactionStatus) String() string {
	switch ts {
	case TransactionActive:
		return "Active"
	case TransactionCompleted:
		return "Completed"
	case TransactionCancelled:
		return "Cancelled"
	case TransactionOverdue:
		return "Overdue"
	default:
		return "Unknown"
	}
}

// Transaction
type Transaction struct {
	ID          string
	Member      *Member
	Book        *Book
	Type        TransactionType
	Date        time.Time
	DueDate     *time.Time
	ReturnDate  *time.Time
	Status      TransactionStatus
	FineAmount  float64
	RenewalCount int
	mu          sync.RWMutex
}

func NewTransaction(member *Member, book *Book, transactionType TransactionType) *Transaction {
	now := time.Now()
	var dueDate *time.Time
	
	if transactionType == Checkout {
		borrowDuration := member.GetBorrowDuration()
		due := now.AddDate(0, 0, borrowDuration)
		dueDate = &due
	}
	
	return &Transaction{
		ID:           fmt.Sprintf("T%d", time.Now().UnixNano()),
		Member:       member,
		Book:         book,
		Type:         transactionType,
		Date:         now,
		DueDate:      dueDate,
		Status:       TransactionActive,
		FineAmount:   0.0,
		RenewalCount: 0,
	}
}

func (t *Transaction) IsOverdue() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	if t.DueDate == nil || t.Status != TransactionActive {
		return false
	}
	
	return time.Now().After(*t.DueDate)
}

func (t *Transaction) GetDaysOverdue() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	if !t.IsOverdue() {
		return 0
	}
	
	return int(time.Now().Sub(*t.DueDate).Hours() / 24)
}

func (t *Transaction) CalculateFine() float64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	// Check if overdue without calling IsOverdue() to avoid deadlock
	if t.DueDate == nil || t.Status != TransactionActive {
		return 0.0
	}
	
	if !time.Now().After(*t.DueDate) {
		return 0.0
	}
	
	daysOverdue := int(time.Now().Sub(*t.DueDate).Hours() / 24)
	
	// Fine calculation: $0.50 per day for first week, $1.00 per day after
	if daysOverdue <= 7 {
		return float64(daysOverdue) * 0.50
	} else {
		return 7*0.50 + float64(daysOverdue-7)*1.00
	}
}

func (t *Transaction) UpdateFine() {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	// Calculate fine without calling CalculateFine() to avoid deadlock
	if t.DueDate == nil || t.Status != TransactionActive {
		t.FineAmount = 0.0
		return
	}
	
	if !time.Now().After(*t.DueDate) {
		t.FineAmount = 0.0
		return
	}
	
	daysOverdue := int(time.Now().Sub(*t.DueDate).Hours() / 24)
	
	// Fine calculation: $0.50 per day for first week, $1.00 per day after
	if daysOverdue <= 7 {
		t.FineAmount = float64(daysOverdue) * 0.50
	} else {
		t.FineAmount = 7*0.50 + float64(daysOverdue-7)*1.00
	}
}

func (t *Transaction) GetFineAmount() float64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.FineAmount
}

func (t *Transaction) Complete() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	if t.Status != TransactionActive {
		return fmt.Errorf("transaction is not active")
	}
	
	now := time.Now()
	t.ReturnDate = &now
	t.Status = TransactionCompleted
	return nil
}

func (t *Transaction) Cancel() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	if t.Status != TransactionActive {
		return fmt.Errorf("transaction is not active")
	}
	
	t.Status = TransactionCancelled
	return nil
}

func (t *Transaction) Renew() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	
	if t.Status != TransactionActive {
		return fmt.Errorf("transaction is not active")
	}
	
	if t.Type != Checkout {
		return fmt.Errorf("only checkout transactions can be renewed")
	}
	
	// Check renewal limit (max 2 renewals)
	if t.RenewalCount >= 2 {
		return fmt.Errorf("maximum renewal limit reached")
	}
	
	// Extend due date by borrow duration
	borrowDuration := t.Member.GetBorrowDuration()
	newDueDate := t.DueDate.AddDate(0, 0, borrowDuration)
	t.DueDate = &newDueDate
	t.RenewalCount++
	
	return nil
}

func (t *Transaction) MarkAsOverdue() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Status = TransactionOverdue
}

func (t *Transaction) GetDetails() map[string]interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	return map[string]interface{}{
		"id":             t.ID,
		"member_id":      t.Member.ID,
		"member_name":    t.Member.Name,
		"book_id":        t.Book.ID,
		"book_title":     t.Book.Title,
		"type":           t.Type.String(),
		"date":           t.Date,
		"due_date":       t.DueDate,
		"return_date":    t.ReturnDate,
		"status":         t.Status.String(),
		"fine_amount":    t.FineAmount,
		"renewal_count":  t.RenewalCount,
		"is_overdue":     t.IsOverdue(),
		"days_overdue":   t.GetDaysOverdue(),
	}
}

// BookReservation
type BookReservation struct {
	ID        string
	Member    *Member
	Book      *Book
	Date      time.Time
	ExpiryDate time.Time
	Status    TransactionStatus
	mu        sync.RWMutex
}

func NewBookReservation(member *Member, book *Book) *BookReservation {
	now := time.Now()
	expiryDate := now.Add(48 * time.Hour) // 48 hours to pick up
	
	return &BookReservation{
		ID:         fmt.Sprintf("R%d", time.Now().UnixNano()),
		Member:     member,
		Book:       book,
		Date:       now,
		ExpiryDate: expiryDate,
		Status:     TransactionActive,
	}
}

func (r *BookReservation) IsExpired() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return time.Now().After(r.ExpiryDate)
}

func (r *BookReservation) Complete() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if r.Status != TransactionActive {
		return fmt.Errorf("reservation is not active")
	}
	
	r.Status = TransactionCompleted
	return nil
}

func (r *BookReservation) Cancel() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if r.Status != TransactionActive {
		return fmt.Errorf("reservation is not active")
	}
	
	r.Status = TransactionCancelled
	return nil
}

func (r *BookReservation) GetDetails() map[string]interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	return map[string]interface{}{
		"id":          r.ID,
		"member_id":   r.Member.ID,
		"member_name": r.Member.Name,
		"book_id":     r.Book.ID,
		"book_title":  r.Book.Title,
		"date":        r.Date,
		"expiry_date": r.ExpiryDate,
		"status":      r.Status.String(),
		"is_expired":  r.IsExpired(),
	}
}
