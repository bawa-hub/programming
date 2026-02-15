package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	Save(user User) error
}

type EmailSender interface {
	SendWelcome(user User) error
}

type InMemoryUserRepository struct{}

func (r *InMemoryUserRepository) Save(user User) error {
	fmt.Printf("[DB] Saving user: %#v\n", user)
	return nil
}

type ConsoleEmailSender struct{}

func (s *ConsoleEmailSender) SendWelcome(user User) error {
	fmt.Printf("[Email] Sending welcome email to %s\n", user.Email)
	return nil
}

type UserOnboardingService struct {
	repo   UserRepository
	emails EmailSender
}

func NewUserOnboardingService(repo UserRepository, emails EmailSender) *UserOnboardingService {
	return &UserOnboardingService{repo: repo, emails: emails}
}

func (u *UserOnboardingService) Onboard(user User) error {
	if err := u.repo.Save(user); err != nil {
		return err
	}
	return u.emails.SendWelcome(user)
}

func main() {
	repo := &InMemoryUserRepository{}
	emails := &ConsoleEmailSender{}
	service := NewUserOnboardingService(repo, emails)

	u := User{ID: 2, Name: "Bob", Email: "bob@example.com"}
	_ = service.Onboard(u)
}
