package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

// SRP violation: one type doing persistence and email responsibilities
// together, making it harder to maintain and test.
type UserService struct{}

func (s *UserService) SaveUser(user User) error {
	fmt.Printf("[DB] Saving user: %#v\n", user)
	return nil
}

func (s *UserService) SendWelcomeEmail(user User) error {
	fmt.Printf("[Email] Sending welcome email to %s\n", user.Email)
	return nil
}

func main() {
	service := &UserService{}
	u := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	_ = service.SaveUser(u)
	_ = service.SendWelcomeEmail(u)
}
