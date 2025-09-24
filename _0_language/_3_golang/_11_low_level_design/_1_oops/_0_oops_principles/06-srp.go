package main

import "fmt"

// 1. SINGLE RESPONSIBILITY PRINCIPLE (SRP)
// Each class should have only one reason to change

// User - handles only user data
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) GetDisplayName() string {
	return u.Name
}

// EmailService - handles only email operations
type EmailService struct{}

func (es *EmailService) SendEmail(to, subject, body string) error {
	fmt.Printf("Sending email to %s: %s\n", to, subject)
	return nil
}
