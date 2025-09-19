package main

import (
	"fmt"
	"time"

	"library_management/models"
	"library_management/services"
)

func main() {
	fmt.Println("=== LIBRARY MANAGEMENT SYSTEM DEMONSTRATION ===\n")

	// Create library management service
	lms := services.NewLibraryManagementService()

	// Create libraries
	library1 := lms.CreateLibrary("Central Library", "123 Main St, City", "555-0101", "central@library.com")
	library2 := lms.CreateLibrary("University Library", "456 University Ave, Campus", "555-0102", "university@library.com")

	fmt.Println("1. LIBRARY CREATION:")
	fmt.Printf("Library 1: %s (%s)\n", library1.Name, library1.Address)
	fmt.Printf("Library 2: %s (%s)\n", library2.Name, library2.Address)

	// Add books to library 1
	book1, _ := lms.AddBook(library1.ID, "978-0134685991", "Effective Go", "Robert Griesemer", models.Reference, "Addison-Wesley", 2022, 300)
	book2, _ := lms.AddBook(library1.ID, "978-0132350884", "Clean Code", "Robert Martin", models.NonFiction, "Prentice Hall", 2008, 464)
	book3, _ := lms.AddBook(library1.ID, "978-0201633610", "Design Patterns", "Gang of Four", models.Reference, "Addison-Wesley", 1994, 395)
	_, _ = lms.AddBook(library1.ID, "978-0134685991", "The Go Programming Language", "Alan Donovan", models.Textbook, "Addison-Wesley", 2015, 380)
	_, _ = lms.AddBook(library1.ID, "978-0132350884", "Harry Potter", "J.K. Rowling", models.Fiction, "Bloomsbury", 1997, 223)

	// Add books to library 2
	book6, _ := lms.AddBook(library2.ID, "978-0134685991", "Introduction to Algorithms", "Thomas Cormen", models.Textbook, "MIT Press", 2009, 1312)
	_, _ = lms.AddBook(library2.ID, "978-0132350884", "Database System Concepts", "Abraham Silberschatz", models.Textbook, "McGraw-Hill", 2019, 1344)

	fmt.Println()
	fmt.Println("2. BOOK ADDITION:")
	fmt.Printf("Added %d books to %s\n", library1.GetBookCount(), library1.Name)
	fmt.Printf("Added %d books to %s\n", library2.GetBookCount(), library2.Name)

	// Register members
	member1, _ := lms.RegisterMember(library1.ID, "Alice Johnson", "alice@email.com", "555-1001", models.Student)
	member2, _ := lms.RegisterMember(library1.ID, "Bob Smith", "bob@email.com", "555-1002", models.Faculty)
	member3, _ := lms.RegisterMember(library1.ID, "Charlie Brown", "charlie@email.com", "555-1003", models.General)
	member4, _ := lms.RegisterMember(library2.ID, "Diana Prince", "diana@email.com", "555-1004", models.VIP)

	fmt.Println()
	fmt.Println("3. MEMBER REGISTRATION:")
	fmt.Printf("Member 1: %s (%s) - %s\n", member1.Name, member1.Email, member1.MemberType.String())
	fmt.Printf("Member 2: %s (%s) - %s\n", member2.Name, member2.Email, member2.MemberType.String())
	fmt.Printf("Member 3: %s (%s) - %s\n", member3.Name, member3.Email, member3.MemberType.String())
	fmt.Printf("Member 4: %s (%s) - %s\n", member4.Name, member4.Email, member4.MemberType.String())

	// Checkout books
	transaction1, _ := lms.CheckoutBook(library1.ID, member1.ID, book1.ID)
	transaction2, _ := lms.CheckoutBook(library1.ID, member2.ID, book2.ID)
	_, _ = lms.CheckoutBook(library1.ID, member3.ID, book3.ID)
	_, _ = lms.CheckoutBook(library2.ID, member4.ID, book6.ID)

	fmt.Println()
	fmt.Println("4. BOOK CHECKOUT:")
	fmt.Printf("Transaction 1: %s borrowed %s\n", member1.Name, book1.Title)
	fmt.Printf("Transaction 2: %s borrowed %s\n", member2.Name, book2.Title)
	fmt.Printf("Transaction 3: %s borrowed %s\n", member3.Name, book3.Title)
	fmt.Printf("Transaction 4: %s borrowed %s\n", member4.Name, book6.Title)

	// Reserve a book
	reservation1, _ := lms.ReserveBook(library1.ID, member1.ID, book2.ID) // Alice reserves Clean Code

	fmt.Println()
	fmt.Println("5. BOOK RESERVATION:")
	fmt.Printf("Reservation 1: %s reserved %s\n", member1.Name, book2.Title)

	// Search books
	searchResults := lms.SearchBooks(library1.ID, "Go")
	fmt.Println()
	fmt.Println("6. BOOK SEARCH:")
	fmt.Printf("Search results for 'Go' in %s: %d books found\n", library1.Name, len(searchResults))
	for i, book := range searchResults {
		if i < 3 { // Show first 3 results
			fmt.Printf("  %d. %s by %s (%s)\n", i+1, book.Title, book.Author, book.Category.String())
		}
	}

	// Get books by category
	fictionBooks := lms.GetBooksByCategory(library1.ID, models.Fiction)
	fmt.Println()
	fmt.Println("7. BOOKS BY CATEGORY:")
	fmt.Printf("Fiction books in %s: %d books\n", library1.Name, len(fictionBooks))
	for _, book := range fictionBooks {
		fmt.Printf("  - %s by %s\n", book.Title, book.Author)
	}

	// Get popular books
	popularBooks := lms.GetPopularBooks(library1.ID, 3)
	fmt.Println()
	fmt.Println("8. POPULAR BOOKS:")
	fmt.Printf("Top 3 popular books in %s:\n", library1.Name)
	for i, book := range popularBooks {
		fmt.Printf("  %d. %s (borrowed %d times)\n", i+1, book.Title, book.GetBorrowCount())
	}

	// Get available books
	availableBooks := lms.GetAvailableBooks(library1.ID)
	fmt.Println()
	fmt.Println("9. AVAILABLE BOOKS:")
	fmt.Printf("Available books in %s: %d books\n", library1.Name, len(availableBooks))
	for _, book := range availableBooks {
		fmt.Printf("  - %s by %s\n", book.Title, book.Author)
	}

	// Get members by type
	facultyMembers := lms.GetMembersByType(library1.ID, models.Faculty)
	fmt.Println()
	fmt.Println("10. MEMBERS BY TYPE:")
	fmt.Printf("Faculty members in %s: %d members\n", library1.Name, len(facultyMembers))
	for _, member := range facultyMembers {
		fmt.Printf("  - %s (%s)\n", member.Name, member.Email)
	}

	// Return a book
	err := lms.ReturnBook(library1.ID, transaction1.ID)
	if err != nil {
		fmt.Printf("Error returning book: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("11. BOOK RETURN:")
		fmt.Printf("Book '%s' returned by %s\n", book1.Title, member1.Name)
	}

	// Renew a book
	renewalTransaction, err := lms.RenewBook(library1.ID, transaction2.ID)
	if err != nil {
		fmt.Printf("Error renewing book: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("12. BOOK RENEWAL:")
		fmt.Printf("Book '%s' renewed by %s\n", book2.Title, member2.Name)
		fmt.Printf("New due date: %s\n", renewalTransaction.DueDate.Format("2006-01-02"))
	}

	// Get member borrowing history
	history := lms.GetMemberBorrowingHistory(library1.ID, member1.ID)
	fmt.Println()
	fmt.Println("13. MEMBER BORROWING HISTORY:")
	fmt.Printf("Borrowing history for %s: %d transactions\n", member1.Name, len(history))
	for i, transaction := range history {
		if i < 3 { // Show first 3 transactions
			fmt.Printf("  %d. %s - %s (%s)\n", i+1, transaction.Book.Title, transaction.Type.String(), transaction.Date.Format("2006-01-02"))
		}
	}

	// Get library statistics
	stats1 := lms.GetLibraryStatistics(library1.ID)
	stats2 := lms.GetLibraryStatistics(library2.ID)

	fmt.Println()
	fmt.Println("14. LIBRARY STATISTICS:")
	fmt.Printf("Library 1 (%s) Statistics:\n", library1.Name)
	fmt.Printf("  Total Books: %v\n", stats1["total_books"])
	fmt.Printf("  Available Books: %v\n", stats1["available_books"])
	fmt.Printf("  Borrowed Books: %v\n", stats1["borrowed_books"])
	fmt.Printf("  Total Members: %v\n", stats1["total_members"])
	fmt.Printf("  Active Members: %v\n", stats1["active_members"])
	fmt.Printf("  Total Transactions: %v\n", stats1["total_transactions"])

	fmt.Printf("\nLibrary 2 (%s) Statistics:\n", library2.Name)
	fmt.Printf("  Total Books: %v\n", stats2["total_books"])
	fmt.Printf("  Available Books: %v\n", stats2["available_books"])
	fmt.Printf("  Borrowed Books: %v\n", stats2["borrowed_books"])
	fmt.Printf("  Total Members: %v\n", stats2["total_members"])
	fmt.Printf("  Active Members: %v\n", stats2["active_members"])
	fmt.Printf("  Total Transactions: %v\n", stats2["total_transactions"])

	// Get overdue books
	overdueBooks := lms.GetOverdueBooks(library1.ID)
	fmt.Println()
	fmt.Println("15. OVERDUE BOOKS:")
	fmt.Printf("Overdue books in %s: %d books\n", library1.Name, len(overdueBooks))
	for _, transaction := range overdueBooks {
		fmt.Printf("  - %s borrowed by %s (overdue by %d days)\n", 
			transaction.Book.Title, transaction.Member.Name, transaction.GetDaysOverdue())
	}

	// Test member limits
	fmt.Println()
	fmt.Println("16. MEMBER BORROWING LIMITS:")
	fmt.Printf("%s can borrow up to %d books (currently has %d)\n", 
		member1.Name, member1.GetMaxBorrowLimit(), member1.GetBorrowedBookCount())
	fmt.Printf("%s can borrow up to %d books (currently has %d)\n", 
		member2.Name, member2.GetMaxBorrowLimit(), member2.GetBorrowedBookCount())
	fmt.Printf("%s can borrow up to %d books (currently has %d)\n", 
		member3.Name, member3.GetMaxBorrowLimit(), member3.GetBorrowedBookCount())
	fmt.Printf("%s can borrow up to %d books (currently has %d)\n", 
		member4.Name, member4.GetMaxBorrowLimit(), member4.GetBorrowedBookCount())

	// Test fine calculation
	fmt.Println()
	fmt.Println("17. FINE CALCULATION:")
	// Simulate an overdue transaction
	overdueTransaction := models.NewTransaction(member1, book1, models.Checkout)
	overdueTransaction.DueDate = &time.Time{} // Set to past date
	overdueTransaction.UpdateFine()
	fmt.Printf("Fine for overdue book: $%.2f\n", overdueTransaction.GetFineAmount())

	// Test book status changes
	fmt.Println()
	fmt.Println("18. BOOK STATUS MANAGEMENT:")
	book1.UpdateStatus(models.Maintenance)
	fmt.Printf("Book '%s' status changed to: %s\n", book1.Title, book1.GetStatus().String())
	
	book1.UpdateStatus(models.Available)
	fmt.Printf("Book '%s' status changed to: %s\n", book1.Title, book1.GetStatus().String())

	// Test member status
	fmt.Println()
	fmt.Println("19. MEMBER STATUS MANAGEMENT:")
	member1.UpdateStatus(models.Suspended)
	fmt.Printf("Member '%s' status changed to: %s\n", member1.Name, member1.Status.String())
	
	member1.UpdateStatus(models.Active)
	fmt.Printf("Member '%s' status changed to: %s\n", member1.Name, member1.Status.String())

	// Test reservation expiration
	fmt.Println()
	fmt.Println("20. RESERVATION MANAGEMENT:")
	fmt.Printf("Reservation for '%s' by %s expires at: %s\n", 
		book2.Title, member1.Name, reservation1.ExpiryDate.Format("2006-01-02 15:04:05"))
	fmt.Printf("Is reservation expired? %t\n", reservation1.IsExpired())

	fmt.Println()
	fmt.Println("=== END OF DEMONSTRATION ===")
}
