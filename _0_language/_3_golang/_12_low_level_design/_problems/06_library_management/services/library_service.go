package services

import (
	"fmt"
	"sort"
	"sync"

	"library_management/models"
)

// Library Management Service
type LibraryManagementService struct {
	libraries map[string]*models.Library
	mu        sync.RWMutex
}

func NewLibraryManagementService() *LibraryManagementService {
	return &LibraryManagementService{
		libraries: make(map[string]*models.Library),
	}
}

// Library Management
func (lms *LibraryManagementService) CreateLibrary(name, address, phone, email string) *models.Library {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := models.NewLibrary(name, address, phone, email)
	lms.libraries[library.ID] = library
	return library
}

func (lms *LibraryManagementService) GetLibrary(libraryID string) *models.Library {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	return lms.libraries[libraryID]
}

func (lms *LibraryManagementService) GetAllLibraries() []*models.Library {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	var libraries []*models.Library
	for _, library := range lms.libraries {
		libraries = append(libraries, library)
	}
	return libraries
}

// Book Management
func (lms *LibraryManagementService) AddBook(libraryID, isbn, title, author string, category models.BookCategory, publisher string, publicationYear, pages int) (*models.Book, error) {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil, fmt.Errorf("library not found")
	}
	
	book := models.NewBook(isbn, title, author, category, publisher, publicationYear, pages)
	library.AddBook(book)
	return book, nil
}

func (lms *LibraryManagementService) GetBook(libraryID, bookID string) *models.Book {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	return library.GetBook(bookID)
}

func (lms *LibraryManagementService) SearchBooks(libraryID, query string) []*models.Book {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	
	var results []*models.Book
	for _, book := range library.GetAllBooks() {
		if containsIgnoreCase(book.Title, query) || 
		   containsIgnoreCase(book.Author, query) ||
		   containsIgnoreCase(book.Publisher, query) {
			results = append(results, book)
		}
	}
	
	// Sort by title
	sort.Slice(results, func(i, j int) bool {
		return results[i].Title < results[j].Title
	})
	
	return results
}

func (lms *LibraryManagementService) GetBooksByCategory(libraryID string, category models.BookCategory) []*models.Book {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	
	var results []*models.Book
	for _, book := range library.GetAllBooks() {
		if book.Category == category {
			results = append(results, book)
		}
	}
	
	// Sort by title
	sort.Slice(results, func(i, j int) bool {
		return results[i].Title < results[j].Title
	})
	
	return results
}

func (lms *LibraryManagementService) GetAvailableBooks(libraryID string) []*models.Book {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	return library.GetAvailableBooks()
}

func (lms *LibraryManagementService) GetPopularBooks(libraryID string, limit int) []*models.Book {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	
	books := library.GetAllBooks()
	
	// Sort by borrow count
	sort.Slice(books, func(i, j int) bool {
		return books[i].GetBorrowCount() > books[j].GetBorrowCount()
	})
	
	if limit > 0 && limit < len(books) {
		return books[:limit]
	}
	return books
}

// Member Management
func (lms *LibraryManagementService) RegisterMember(libraryID, name, email, phone string, memberType models.MemberType) (*models.Member, error) {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil, fmt.Errorf("library not found")
	}
	
	// Check if member already exists
	for _, member := range library.GetAllMembers() {
		if member.Email == email {
			return nil, fmt.Errorf("member with this email already exists")
		}
	}
	
	member := models.NewMember(name, email, phone, memberType)
	library.AddMember(member)
	return member, nil
}

func (lms *LibraryManagementService) GetMember(libraryID, memberID string) *models.Member {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	return library.GetMember(memberID)
}

func (lms *LibraryManagementService) GetMembersByType(libraryID string, memberType models.MemberType) []*models.Member {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	
	var results []*models.Member
	for _, member := range library.GetAllMembers() {
		if member.MemberType == memberType {
			results = append(results, member)
		}
	}
	
	// Sort by name
	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})
	
	return results
}

// Lending Operations
func (lms *LibraryManagementService) CheckoutBook(libraryID, memberID, bookID string) (*models.Transaction, error) {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil, fmt.Errorf("library not found")
	}
	
	member := library.GetMember(memberID)
	if member == nil {
		return nil, fmt.Errorf("member not found")
	}
	
	book := library.GetBook(bookID)
	if book == nil {
		return nil, fmt.Errorf("book not found")
	}
	
	// Check if member can borrow
	if !member.CanBorrowBook() {
		return nil, fmt.Errorf("member cannot borrow more books")
	}
	
	// Check if book is available
	if !book.IsAvailable() {
		return nil, fmt.Errorf("book is not available")
	}
	
	// Create transaction
	transaction := models.NewTransaction(member, book, models.Checkout)
	
	// Update book status
	err := book.Borrow()
	if err != nil {
		return nil, err
	}
	
	// Add book to member's borrowed books
	err = member.AddBorrowedBook(book)
	if err != nil {
		// Rollback book status
		book.Return()
		return nil, err
	}
	
	// Add transaction to library
	library.AddTransaction(transaction)
	
	return transaction, nil
}

func (lms *LibraryManagementService) ReturnBook(libraryID, transactionID string) error {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return fmt.Errorf("library not found")
	}
	
	transaction := library.GetTransaction(transactionID)
	if transaction == nil {
		return fmt.Errorf("transaction not found")
	}
	
	// Complete transaction
	err := transaction.Complete()
	if err != nil {
		return err
	}
	
	// Update book status
	err = transaction.Book.Return()
	if err != nil {
		return err
	}
	
	// Remove book from member's borrowed books
	err = transaction.Member.RemoveBorrowedBook(transaction.Book.ID)
	if err != nil {
		return err
	}
	
	return nil
}

func (lms *LibraryManagementService) RenewBook(libraryID, transactionID string) (*models.Transaction, error) {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil, fmt.Errorf("library not found")
	}
	
	transaction := library.GetTransaction(transactionID)
	if transaction == nil {
		return nil, fmt.Errorf("transaction not found")
	}
	
	// Renew transaction
	err := transaction.Renew()
	if err != nil {
		return nil, err
	}
	
	// Create new renewal transaction
	renewalTransaction := models.NewTransaction(transaction.Member, transaction.Book, models.Renewal)
	renewalTransaction.DueDate = transaction.DueDate
	library.AddTransaction(renewalTransaction)
	
	return renewalTransaction, nil
}

// Reservation Operations
func (lms *LibraryManagementService) ReserveBook(libraryID, memberID, bookID string) (*models.BookReservation, error) {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil, fmt.Errorf("library not found")
	}
	
	member := library.GetMember(memberID)
	if member == nil {
		return nil, fmt.Errorf("member not found")
	}
	
	book := library.GetBook(bookID)
	if book == nil {
		return nil, fmt.Errorf("book not found")
	}
	
	// Check if book is available for reservation
	if book.Status != models.Borrowed {
		return nil, fmt.Errorf("book is not currently borrowed, cannot reserve")
	}
	
	// Create reservation
	reservation := models.NewBookReservation(member, book)
	
	// Add reservation to library and member
	library.AddReservation(reservation)
	member.AddReservation(reservation)
	
	return reservation, nil
}

func (lms *LibraryManagementService) CancelReservation(libraryID, reservationID string) error {
	lms.mu.Lock()
	defer lms.mu.Unlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return fmt.Errorf("library not found")
	}
	
	reservation := library.GetReservation(reservationID)
	if reservation == nil {
		return fmt.Errorf("reservation not found")
	}
	
	// Cancel reservation
	err := reservation.Cancel()
	if err != nil {
		return err
	}
	
	// Remove from library and member
	library.RemoveReservation(reservationID)
	reservation.Member.RemoveReservation(reservationID)
	
	return nil
}

// Reporting
func (lms *LibraryManagementService) GetLibraryStatistics(libraryID string) map[string]interface{} {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	
	return library.GetLibraryStats()
}

func (lms *LibraryManagementService) GetOverdueBooks(libraryID string) []*models.Transaction {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	
	return library.GetOverdueTransactions()
}

func (lms *LibraryManagementService) GetMemberBorrowingHistory(libraryID, memberID string) []*models.Transaction {
	lms.mu.RLock()
	defer lms.mu.RUnlock()
	
	library := lms.libraries[libraryID]
	if library == nil {
		return nil
	}
	
	var history []*models.Transaction
	for _, transaction := range library.GetAllTransactions() {
		if transaction.Member.ID == memberID {
			history = append(history, transaction)
		}
	}
	
	// Sort by date (newest first)
	sort.Slice(history, func(i, j int) bool {
		return history[i].Date.After(history[j].Date)
	})
	
	return history
}

// Helper Functions
func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) && 
		   (s == substr || 
		    len(s) > len(substr) && 
		    (s[:len(substr)] == substr || 
		     s[len(s)-len(substr):] == substr ||
		     containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
