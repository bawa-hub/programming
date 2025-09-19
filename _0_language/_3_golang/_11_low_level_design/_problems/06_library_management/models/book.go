package models

import (
	"fmt"
	"sync"
	"time"
)

// Book Status
type BookStatus int

const (
	Available BookStatus = iota
	Borrowed
	Reserved
	Lost
	Maintenance
)

func (bs BookStatus) String() string {
	switch bs {
	case Available:
		return "Available"
	case Borrowed:
		return "Borrowed"
	case Reserved:
		return "Reserved"
	case Lost:
		return "Lost"
	case Maintenance:
		return "Maintenance"
	default:
		return "Unknown"
	}
}

// Book Category
type BookCategory int

const (
	Fiction BookCategory = iota
	NonFiction
	Science
	History
	Biography
	Reference
	Textbook
	Children
	Other
)

func (bc BookCategory) String() string {
	switch bc {
	case Fiction:
		return "Fiction"
	case NonFiction:
		return "Non-Fiction"
	case Science:
		return "Science"
	case History:
		return "History"
	case Biography:
		return "Biography"
	case Reference:
		return "Reference"
	case Textbook:
		return "Textbook"
	case Children:
		return "Children"
	case Other:
		return "Other"
	default:
		return "Unknown"
	}
}

// Book
type Book struct {
	ID              string
	ISBN            string
	Title           string
	Author          string
	Category        BookCategory
	Publisher       string
	PublicationYear int
	Pages           int
	Status          BookStatus
	Location        string
	AddedDate       time.Time
	LastBorrowed    *time.Time
	BorrowCount     int
	mu              sync.RWMutex
}

func NewBook(isbn, title, author string, category BookCategory, publisher string, publicationYear, pages int) *Book {
	return &Book{
		ID:              fmt.Sprintf("B%d", time.Now().UnixNano()),
		ISBN:            isbn,
		Title:           title,
		Author:          author,
		Category:        category,
		Publisher:       publisher,
		PublicationYear: publicationYear,
		Pages:           pages,
		Status:          Available,
		Location:        "Shelf A-1",
		AddedDate:       time.Now(),
		BorrowCount:     0,
	}
}

func (b *Book) UpdateStatus(status BookStatus) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Status = status
}

func (b *Book) GetStatus() BookStatus {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Status
}

func (b *Book) IsAvailable() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Status == Available
}

func (b *Book) IsBorrowed() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Status == Borrowed
}

func (b *Book) IsReserved() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Status == Reserved
}

func (b *Book) Borrow() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.Status != Available {
		return fmt.Errorf("book is not available for borrowing")
	}
	
	b.Status = Borrowed
	b.BorrowCount++
	now := time.Now()
	b.LastBorrowed = &now
	return nil
}

func (b *Book) Return() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.Status != Borrowed {
		return fmt.Errorf("book is not currently borrowed")
	}
	
	b.Status = Available
	return nil
}

func (b *Book) Reserve() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.Status != Available {
		return fmt.Errorf("book is not available for reservation")
	}
	
	b.Status = Reserved
	return nil
}

func (b *Book) CancelReservation() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if b.Status != Reserved {
		return fmt.Errorf("book is not currently reserved")
	}
	
	b.Status = Available
	return nil
}

func (b *Book) MarkAsLost() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Status = Lost
}

func (b *Book) MarkForMaintenance() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Status = Maintenance
}

func (b *Book) UpdateLocation(location string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Location = location
}

func (b *Book) GetLocation() string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Location
}

func (b *Book) GetBorrowCount() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.BorrowCount
}

func (b *Book) GetLastBorrowed() *time.Time {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.LastBorrowed
}

func (b *Book) UpdateDetails(title, author, publisher string, publicationYear, pages int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Title = title
	b.Author = author
	b.Publisher = publisher
	b.PublicationYear = publicationYear
	b.Pages = pages
}

func (b *Book) GetDetails() map[string]interface{} {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	return map[string]interface{}{
		"id":               b.ID,
		"isbn":             b.ISBN,
		"title":            b.Title,
		"author":           b.Author,
		"category":         b.Category.String(),
		"publisher":        b.Publisher,
		"publication_year": b.PublicationYear,
		"pages":            b.Pages,
		"status":           b.Status.String(),
		"location":         b.Location,
		"added_date":       b.AddedDate,
		"borrow_count":     b.BorrowCount,
		"last_borrowed":    b.LastBorrowed,
	}
}
