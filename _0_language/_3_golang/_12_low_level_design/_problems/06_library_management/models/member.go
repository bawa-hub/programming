package models

import (
	"fmt"
	"sync"
	"time"
)

// Member Type
type MemberType int

const (
	Student MemberType = iota
	Faculty
	General
	VIP
)

func (mt MemberType) String() string {
	switch mt {
	case Student:
		return "Student"
	case Faculty:
		return "Faculty"
	case General:
		return "General"
	case VIP:
		return "VIP"
	default:
		return "Unknown"
	}
}

// Member Status
type MemberStatus int

const (
	Active MemberStatus = iota
	Inactive
	Suspended
	Expired
)

func (ms MemberStatus) String() string {
	switch ms {
	case Active:
		return "Active"
	case Inactive:
		return "Inactive"
	case Suspended:
		return "Suspended"
	case Expired:
		return "Expired"
	default:
		return "Unknown"
	}
}

// Member
type Member struct {
	ID           string
	Name         string
	Email        string
	Phone        string
	MemberType   MemberType
	Status       MemberStatus
	JoinDate     time.Time
	ExpiryDate   time.Time
	BorrowedBooks []*Book
	Reservations []*BookReservation
	TotalFines   float64
	mu           sync.RWMutex
}

func NewMember(name, email, phone string, memberType MemberType) *Member {
	now := time.Now()
	expiryDate := now.AddDate(1, 0, 0) // 1 year from now
	
	return &Member{
		ID:            fmt.Sprintf("M%d", time.Now().UnixNano()),
		Name:          name,
		Email:         email,
		Phone:         phone,
		MemberType:    memberType,
		Status:        Active,
		JoinDate:      now,
		ExpiryDate:    expiryDate,
		BorrowedBooks: make([]*Book, 0),
		Reservations:  make([]*BookReservation, 0),
		TotalFines:    0.0,
	}
}

func (m *Member) GetMaxBorrowLimit() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.getMaxBorrowLimit()
}

func (m *Member) getMaxBorrowLimit() int {
	switch m.MemberType {
	case Student:
		return 5
	case Faculty:
		return 10
	case General:
		return 3
	case VIP:
		return 15
	default:
		return 3
	}
}

func (m *Member) GetBorrowDuration() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.getBorrowDuration()
}

func (m *Member) getBorrowDuration() int {
	switch m.MemberType {
	case Student:
		return 14 // 14 days
	case Faculty:
		return 30 // 30 days
	case General:
		return 14 // 14 days
	case VIP:
		return 30 // 30 days
	default:
		return 14
	}
}

func (m *Member) CanBorrowBook() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	maxLimit := m.getMaxBorrowLimit()
	return m.Status == Active && 
		   len(m.BorrowedBooks) < maxLimit &&
		   time.Now().Before(m.ExpiryDate)
}

func (m *Member) AddBorrowedBook(book *Book) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// Check if member can borrow (without calling CanBorrowBook to avoid deadlock)
	maxLimit := m.getMaxBorrowLimit()
	if m.Status != Active || len(m.BorrowedBooks) >= maxLimit || time.Now().After(m.ExpiryDate) {
		return fmt.Errorf("member cannot borrow more books")
	}
	
	// Check if book is already borrowed by this member
	for _, borrowedBook := range m.BorrowedBooks {
		if borrowedBook.ID == book.ID {
			return fmt.Errorf("book is already borrowed by this member")
		}
	}
	
	m.BorrowedBooks = append(m.BorrowedBooks, book)
	return nil
}

func (m *Member) RemoveBorrowedBook(bookID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	for i, book := range m.BorrowedBooks {
		if book.ID == bookID {
			m.BorrowedBooks = append(m.BorrowedBooks[:i], m.BorrowedBooks[i+1:]...)
			return nil
		}
	}
	
	return fmt.Errorf("book not found in borrowed books")
}

func (m *Member) GetBorrowedBooks() []*Book {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	books := make([]*Book, len(m.BorrowedBooks))
	copy(books, m.BorrowedBooks)
	return books
}

func (m *Member) GetBorrowedBookCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.BorrowedBooks)
}

func (m *Member) AddReservation(reservation *BookReservation) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Reservations = append(m.Reservations, reservation)
}

func (m *Member) RemoveReservation(reservationID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	for i, reservation := range m.Reservations {
		if reservation.ID == reservationID {
			m.Reservations = append(m.Reservations[:i], m.Reservations[i+1:]...)
			return nil
		}
	}
	
	return fmt.Errorf("reservation not found")
}

func (m *Member) GetReservations() []*BookReservation {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	reservations := make([]*BookReservation, len(m.Reservations))
	copy(reservations, m.Reservations)
	return reservations
}

func (m *Member) AddFine(amount float64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.TotalFines += amount
}

func (m *Member) PayFine(amount float64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if amount > m.TotalFines {
		return fmt.Errorf("payment amount exceeds total fines")
	}
	
	m.TotalFines -= amount
	return nil
}

func (m *Member) GetTotalFines() float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.TotalFines
}

func (m *Member) UpdateStatus(status MemberStatus) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Status = status
}

func (m *Member) RenewMembership() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ExpiryDate = time.Now().AddDate(1, 0, 0)
	m.Status = Active
}

func (m *Member) IsMembershipExpired() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return time.Now().After(m.ExpiryDate)
}

func (m *Member) UpdateProfile(name, email, phone string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Name = name
	m.Email = email
	m.Phone = phone
}

func (m *Member) GetProfile() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	return map[string]interface{}{
		"id":                m.ID,
		"name":              m.Name,
		"email":             m.Email,
		"phone":             m.Phone,
		"member_type":       m.MemberType.String(),
		"status":            m.Status.String(),
		"join_date":         m.JoinDate,
		"expiry_date":       m.ExpiryDate,
		"borrowed_books":    len(m.BorrowedBooks),
		"max_borrow_limit":  m.GetMaxBorrowLimit(),
		"borrow_duration":   m.GetBorrowDuration(),
		"total_fines":       m.TotalFines,
		"is_expired":        m.IsMembershipExpired(),
	}
}
