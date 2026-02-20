package exercises

import (
	"context"
	"fmt"
	"time"
)

// Exercise 9: Context with Database Operations
type User struct {
	ID    string
	Email string
	Role  string
}

func Exercise9() {
	fmt.Println("\nExercise 9: Context with Database Operations")
	fmt.Println("===========================================")
	
	// TODO: Simulate database operations with context
	// 1. Create context with timeout
	// 2. Simulate database queries
	// 3. Handle timeouts and cancellations
	// 4. Implement proper cleanup
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Simulate database operations
	user, err := getUserFromDB(ctx, "user-123")
	if err != nil {
		fmt.Printf("  Exercise 9: Error getting user: %v\n", err)
		return
	}
	
	fmt.Printf("  Exercise 9: User retrieved: %+v\n", user)
	
	// Update user
	err = updateUserInDB(ctx, user)
	if err != nil {
		fmt.Printf("  Exercise 9: Error updating user: %v\n", err)
		return
	}
	
	fmt.Println("  Exercise 9: User updated successfully")
	fmt.Println("Exercise 9 completed")
}

func getUserFromDB(ctx context.Context, userID string) (*User, error) {
	// Simulate database query
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("database query cancelled: %v", ctx.Err())
	case <-time.After(2 * time.Second):
		return &User{ID: userID, Email: "john@example.com", Role: "admin"}, nil
	}
}

func updateUserInDB(ctx context.Context, user *User) error {
	// Simulate database update
	select {
	case <-ctx.Done():
		return fmt.Errorf("database update cancelled: %v", ctx.Err())
	case <-time.After(1 * time.Second):
		user.Email = "john.updated@example.com"
		return nil
	}
}