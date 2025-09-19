package models

import (
	"fmt"
	"sync"
	"time"
)

// Library
type Library struct {
	ID            string
	Name          string
	Address       string
	Phone         string
	Email         string
	Books         map[string]*Book
	Members       map[string]*Member
	Transactions  map[string]*Transaction
	Reservations  map[string]*BookReservation
	OperatingHours OperatingHours
	CreatedAt     time.Time
	mu            sync.RWMutex
}

// Operating Hours
type OperatingHours struct {
	Monday    DayHours
	Tuesday   DayHours
	Wednesday DayHours
	Thursday  DayHours
	Friday    DayHours
	Saturday  DayHours
	Sunday    DayHours
}

type DayHours struct {
	Open  time.Time
	Close time.Time
	IsOpen bool
}

func NewLibrary(name, address, phone, email string) *Library {
	return &Library{
		ID:           fmt.Sprintf("L%d", time.Now().UnixNano()),
		Name:         name,
		Address:      address,
		Phone:        phone,
		Email:        email,
		Books:        make(map[string]*Book),
		Members:      make(map[string]*Member),
		Transactions: make(map[string]*Transaction),
		Reservations: make(map[string]*BookReservation),
		OperatingHours: OperatingHours{
			Monday:    DayHours{Open: time.Date(0, 1, 1, 9, 0, 0, 0, time.UTC), Close: time.Date(0, 1, 1, 21, 0, 0, 0, time.UTC), IsOpen: true},
			Tuesday:   DayHours{Open: time.Date(0, 1, 1, 9, 0, 0, 0, time.UTC), Close: time.Date(0, 1, 1, 21, 0, 0, 0, time.UTC), IsOpen: true},
			Wednesday: DayHours{Open: time.Date(0, 1, 1, 9, 0, 0, 0, time.UTC), Close: time.Date(0, 1, 1, 21, 0, 0, 0, time.UTC), IsOpen: true},
			Thursday:  DayHours{Open: time.Date(0, 1, 1, 9, 0, 0, 0, time.UTC), Close: time.Date(0, 1, 1, 21, 0, 0, 0, time.UTC), IsOpen: true},
			Friday:    DayHours{Open: time.Date(0, 1, 1, 9, 0, 0, 0, time.UTC), Close: time.Date(0, 1, 1, 21, 0, 0, 0, time.UTC), IsOpen: true},
			Saturday:  DayHours{Open: time.Date(0, 1, 1, 10, 0, 0, 0, time.UTC), Close: time.Date(0, 1, 1, 18, 0, 0, 0, time.UTC), IsOpen: true},
			Sunday:    DayHours{Open: time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC), Close: time.Date(0, 1, 1, 17, 0, 0, 0, time.UTC), IsOpen: true},
		},
		CreatedAt: time.Now(),
	}
}

func (l *Library) AddBook(book *Book) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	book, exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("book not found")
	}
	
	if book.Status == Borrowed {
		return fmt.Errorf("cannot remove borrowed book")
	}
	
	delete(l.Books, bookID)
	return nil
}

func (l *Library) GetBook(bookID string) *Book {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.Books[bookID]
}

func (l *Library) GetAllBooks() []*Book {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	books := make([]*Book, 0, len(l.Books))
	for _, book := range l.Books {
		books = append(books, book)
	}
	return books
}

func (l *Library) AddMember(member *Member) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Members[member.ID] = member
}

func (l *Library) RemoveMember(memberID string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	member, exists := l.Members[memberID]
	if !exists {
		return fmt.Errorf("member not found")
	}
	
	if len(member.BorrowedBooks) > 0 {
		return fmt.Errorf("member has borrowed books, cannot remove")
	}
	
	delete(l.Members, memberID)
	return nil
}

func (l *Library) GetMember(memberID string) *Member {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.Members[memberID]
}

func (l *Library) GetAllMembers() []*Member {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	members := make([]*Member, 0, len(l.Members))
	for _, member := range l.Members {
		members = append(members, member)
	}
	return members
}

func (l *Library) AddTransaction(transaction *Transaction) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Transactions[transaction.ID] = transaction
}

func (l *Library) GetTransaction(transactionID string) *Transaction {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.Transactions[transactionID]
}

func (l *Library) GetAllTransactions() []*Transaction {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	transactions := make([]*Transaction, 0, len(l.Transactions))
	for _, transaction := range l.Transactions {
		transactions = append(transactions, transaction)
	}
	return transactions
}

func (l *Library) AddReservation(reservation *BookReservation) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Reservations[reservation.ID] = reservation
}

func (l *Library) RemoveReservation(reservationID string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	_, exists := l.Reservations[reservationID]
	if !exists {
		return fmt.Errorf("reservation not found")
	}
	
	delete(l.Reservations, reservationID)
	return nil
}

func (l *Library) GetReservation(reservationID string) *BookReservation {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.Reservations[reservationID]
}

func (l *Library) GetAllReservations() []*BookReservation {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	reservations := make([]*BookReservation, 0, len(l.Reservations))
	for _, reservation := range l.Reservations {
		reservations = append(reservations, reservation)
	}
	return reservations
}

func (l *Library) GetBookCount() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.Books)
}

func (l *Library) GetMemberCount() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return len(l.Members)
}

func (l *Library) GetAvailableBooks() []*Book {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	var availableBooks []*Book
	for _, book := range l.Books {
		if book.Status == Available {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) GetBorrowedBooks() []*Book {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	var borrowedBooks []*Book
	for _, book := range l.Books {
		if book.Status == Borrowed {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	return borrowedBooks
}

func (l *Library) GetOverdueTransactions() []*Transaction {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	var overdueTransactions []*Transaction
	for _, transaction := range l.Transactions {
		if transaction.IsOverdue() {
			overdueTransactions = append(overdueTransactions, transaction)
		}
	}
	return overdueTransactions
}

func (l *Library) GetLibraryStats() map[string]interface{} {
	l.mu.RLock()
	defer l.mu.RUnlock()
	
	totalBooks := len(l.Books)
	availableBooks := 0
	borrowedBooks := 0
	reservedBooks := 0
	
	for _, book := range l.Books {
		switch book.Status {
		case Available:
			availableBooks++
		case Borrowed:
			borrowedBooks++
		case Reserved:
			reservedBooks++
		}
	}
	
	totalMembers := len(l.Members)
	activeMembers := 0
	for _, member := range l.Members {
		if member.Status == Active {
			activeMembers++
		}
	}
	
	totalTransactions := len(l.Transactions)
	overdueTransactions := len(l.GetOverdueTransactions())
	
	return map[string]interface{}{
		"library_id":         l.ID,
		"library_name":       l.Name,
		"total_books":        totalBooks,
		"available_books":    availableBooks,
		"borrowed_books":     borrowedBooks,
		"reserved_books":     reservedBooks,
		"total_members":      totalMembers,
		"active_members":     activeMembers,
		"total_transactions": totalTransactions,
		"overdue_transactions": overdueTransactions,
		"created_at":         l.CreatedAt,
	}
}
