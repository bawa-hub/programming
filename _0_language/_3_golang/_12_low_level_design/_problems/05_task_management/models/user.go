package models

import (
	"fmt"
	"sync"
	"time"
)

// User Roles
type UserRole int

const (
	Admin UserRole = iota
	ProjectManager
	Developer
	Tester
	Viewer
)

func (ur UserRole) String() string {
	switch ur {
	case Admin:
		return "Admin"
	case ProjectManager:
		return "Project Manager"
	case Developer:
		return "Developer"
	case Tester:
		return "Tester"
	case Viewer:
		return "Viewer"
	default:
		return "Unknown"
	}
}

// User Preferences
type UserPreferences struct {
	Theme           string
	Language        string
	EmailNotifications bool
	PushNotifications  bool
	TimeZone        string
	DateFormat      string
}

// User
type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	Role         UserRole
	Team         *Team
	Preferences  UserPreferences
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	mu           sync.RWMutex
}

func NewUser(username, email, passwordHash, firstName, lastName string) *User {
	return &User{
		ID:        fmt.Sprintf("U%d", time.Now().UnixNano()),
		Username:  username,
		Email:     email,
		PasswordHash: passwordHash,
		FirstName: firstName,
		LastName:  lastName,
		Role:      Developer, // Default role
		Preferences: UserPreferences{
			Theme:              "light",
			Language:           "en",
			EmailNotifications: true,
			PushNotifications:  true,
			TimeZone:           "UTC",
			DateFormat:         "2006-01-02",
		},
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) GetFullName() string {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) UpdateProfile(firstName, lastName, email string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
	u.UpdatedAt = time.Now()
}

func (u *User) UpdatePreferences(preferences UserPreferences) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Preferences = preferences
	u.UpdatedAt = time.Now()
}

func (u *User) SetRole(role UserRole) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Role = role
	u.UpdatedAt = time.Now()
}

func (u *User) SetTeam(team *Team) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Team = team
	u.UpdatedAt = time.Now()
}

func (u *User) Deactivate() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

func (u *User) Activate() {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

func (u *User) HasPermission(permission string) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	
	switch permission {
	case "create_project":
		return u.Role == Admin || u.Role == ProjectManager
	case "manage_users":
		return u.Role == Admin
	case "assign_tasks":
		return u.Role == Admin || u.Role == ProjectManager
	case "view_all_tasks":
		return u.Role == Admin || u.Role == ProjectManager
	case "delete_tasks":
		return u.Role == Admin || u.Role == ProjectManager
	default:
		return false
	}
}

func (u *User) GetRole() UserRole {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.Role
}

func (u *User) GetTeam() *Team {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.Team
}

func (u *User) IsActiveUser() bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.IsActive
}
